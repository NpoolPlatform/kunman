package recoverycode

import (
	"context"
	"fmt"
	"time"

	"github.com/AmirSoleimani/VoucherCodeGenerator/vcgen"
	recoverycodecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user/recoverycode"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuser"
	entrecoverycode "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/recoverycode"
	redis2 "github.com/NpoolPlatform/kunman/framework/redis"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user/recoverycode"
)

const RecoveryCodeLen = 16

type generateHandler struct {
	*Handler
	ids     []uuid.UUID
	updated bool
}

func (h *generateHandler) expireExistCodes(ctx context.Context, tx *ent.Tx) error {
	if h.updated {
		return nil
	}

	now := time.Now().Unix()

	if _, err := tx.
		RecoveryCode.
		Update().
		Where(
			entrecoverycode.AppID(*h.AppID),
			entrecoverycode.UserID(*h.UserID),
			entrecoverycode.DeletedAt(0),
		).
		SetDeletedAt(uint32(now)).
		Save(ctx); err != nil {
		return err
	}
	h.updated = true
	return nil
}

func (h *Handler) getUser(ctx context.Context) error {
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if _, err := cli.
			AppUser.
			Query().
			Where(
				entappuser.AppID(*h.AppID),
				entappuser.EntID(*h.UserID),
				entappuser.DeletedAt(0),
			).
			Only(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *generateHandler) createCode(ctx context.Context, tx *ent.Tx, code string) error {
	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateRecoveryCode, *h.AppID, code)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	exist, err := tx.
		RecoveryCode.
		Query().
		Where(
			entrecoverycode.AppID(*h.AppID),
			entrecoverycode.Code(code),
		).
		Exist(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("code already exist")
	}

	id := uuid.New()
	if _, err := tx.
		RecoveryCode.
		Create().
		SetEntID(id).
		SetAppID(*h.AppID).
		SetUserID(*h.UserID).
		SetCode(code).
		SetUsed(false).
		Save(ctx); err != nil {
		return err
	}
	h.ids = append(h.ids, id)
	return nil
}

func (h *Handler) Generate(ctx context.Context) (string, error) {
	vc, err := vcgen.NewWithOptions(
		vcgen.SetCount(1),
		vcgen.SetPattern("#####-#####"),
		vcgen.SetCharset("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"),
	)
	if err != nil {
		return "", err
	}
	codes, err := vc.Run()
	if err != nil {
		return "", err
	}
	return codes[0], nil
}

func (h *Handler) GenerateRecoveryCodes(ctx context.Context) ([]*npool.RecoveryCode, error) {
	if err := h.getUser(ctx); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateRecoveryCode, *h.AppID, *h.UserID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	codes := []string{}
	for {
		code, err := h.Generate(ctx)
		if err != nil {
			return nil, err
		}
		h.Conds = &recoverycodecrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			Code:  &cruder.Cond{Op: cruder.EQ, Val: code},
		}

		key1 := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateRecoveryCode, *h.AppID, code)
		if err := redis2.TryLock(key1, 0); err != nil {
			return nil, err
		}

		exist, err := h.ExistRecoveryCodeConds(ctx)
		if err != nil {
			_ = redis2.Unlock(key1)
			return nil, err
		}
		_ = redis2.Unlock(key1)

		if exist {
			continue
		}

		codes = append(codes, code)
		if len(codes) >= RecoveryCodeLen {
			break
		}
	}

	handler := &generateHandler{
		Handler: h,
		updated: false,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, code := range codes {
			_fn := func() error {
				if err := handler.expireExistCodes(ctx, tx); err != nil {
					return err
				}
				if err := handler.createCode(ctx, tx, code); err != nil {
					return err
				}
				return nil
			}
			if err := _fn(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &recoverycodecrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: handler.ids},
	}
	h.Limit = 16
	infos, _, err := h.GetRecoveryCodes(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}
