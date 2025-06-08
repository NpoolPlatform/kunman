//nolint:dupl
package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	usercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user"
	entapproleuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/approleuser"
	entappusercontrol "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appusercontrol"
	entappuserextra "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuserextra"
	entappusersecret "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appusersecret"
	entappuserthirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuserthirdparty"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteAppUser(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppUser.
		UpdateOneID(*h.ID).
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		fmt.Println("invalid id: ", err)
		return err
	}
	h.AppID = &info.AppID
	h.EntID = &info.EntID
	return nil
}

func (h *deleteHandler) deleteAppUserExtra(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppUserExtra.
		Query().
		Where(
			entappuserextra.AppID(*h.AppID),
			entappuserextra.UserID(*h.EntID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return nil
	}

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppUserControl(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppUserControl.
		Query().
		Where(
			entappusercontrol.AppID(*h.AppID),
			entappusercontrol.UserID(*h.EntID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return nil
	}

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppUserSecret(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppUserSecret.
		Query().
		Where(
			entappusersecret.AppID(*h.AppID),
			entappusersecret.UserID(*h.EntID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return nil
	}

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppUserThirdParties(ctx context.Context, tx *ent.Tx) error {
	if _, err := tx.
		AppUserThirdParty.
		Update().
		Where(
			entappuserthirdparty.AppID(*h.AppID),
			entappuserthirdparty.UserID(*h.EntID),
		).
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	return nil
}

func (h *deleteHandler) deleteAppRoleUser(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppRoleUser.
		Query().
		Where(
			entapproleuser.AppID(*h.AppID),
			entapproleuser.UserID(*h.EntID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return nil
	}

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteUser(ctx context.Context) (info *npool.User, err error) {
	info, err = h.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid user")
	}
	h.ID = &info.ID

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppUser(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserExtra(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserControl(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserSecret(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserThirdParties(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppRoleUser(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (h *Handler) DeleteThirdUser(ctx context.Context) (info *npool.User, err error) {
	h.Conds = &usercrud.Conds{
		ThirdPartyID:     &cruder.Cond{Op: cruder.EQ, Val: *h.ThirdPartyID},
		ThirdPartyUserID: &cruder.Cond{Op: cruder.EQ, Val: *h.ThirdPartyUserID},
	}
	info, err = h.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	if info.ThirdPartyID == nil || info.ThirdPartyUserID == nil {
		return nil, fmt.Errorf("invalid thirdparty")
	}

	id1, err := uuid.Parse(info.EntID)
	if err != nil {
		return nil, err
	}
	h.EntID = &id1

	id2, err := uuid.Parse(info.AppID)
	if err != nil {
		return nil, err
	}
	h.AppID = &id2

	h.ThirdPartyUserID = info.ThirdPartyUserID

	id3, err := uuid.Parse(*info.ThirdPartyID)
	if err != nil {
		return nil, err
	}
	h.ThirdPartyID = &id3

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			AppUserThirdParty.
			Query().
			Where(
				entappuserthirdparty.AppID(*h.AppID),
				entappuserthirdparty.UserID(*h.EntID),
				entappuserthirdparty.ThirdPartyUserID(*h.ThirdPartyUserID),
				entappuserthirdparty.ThirdPartyID(*h.ThirdPartyID),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
			return nil
		}

		if _, err := info.
			Update().
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
