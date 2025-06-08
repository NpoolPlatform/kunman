package kyc

import (
	"context"
	"fmt"

	kyccrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/kyc"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent"
	user1 "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	redis2 "github.com/NpoolPlatform/kunman/framework/redis"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/kyc"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
)

func (h *Handler) CreateKyc(ctx context.Context) (*npool.Kyc, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateUser, *h.AppID, *h.UserID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	userID := h.UserID.String()
	appID := h.AppID.String()

	h1, err := user1.NewHandler(
		ctx,
		user1.WithAppID(&appID, true),
		user1.WithEntID(&userID, true),
	)
	if err != nil {
		return nil, err
	}
	exist, err := h1.ExistUser(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := kyccrud.SetQueryConds(
			cli.Kyc.Query(),
			&kyccrud.Conds{
				AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
			},
		)
		if err != nil {
			return err
		}

		info, err := stm.Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info != nil {
			return fmt.Errorf("kyc exist")
		}

		if _, err := kyccrud.CreateSet(
			cli.Kyc.Create(),
			&kyccrud.Req{
				EntID:        h.EntID,
				AppID:        h.AppID,
				UserID:       h.UserID,
				DocumentType: h.DocumentType,
				IDNumber:     h.IDNumber,
				FrontImg:     h.FrontImg,
				BackImg:      h.BackImg,
				SelfieImg:    h.SelfieImg,
				EntityType:   h.EntityType,
				ReviewID:     h.ReviewID,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetKyc(ctx)
}
