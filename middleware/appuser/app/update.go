package app

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	appcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/app"
	banappcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/app/ban"
	ctrlcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/app/control"
	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entappctrl "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appcontrol"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	oldAppID *uuid.UUID
}

func (h *updateHandler) updateApp(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		App.
		Query().
		Where(
			entapp.ID(*h.ID),
			entapp.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info == nil {
		return nil
	}
	h.oldAppID = &info.EntID
	req := &appcrud.Req{
		Name:        h.Name,
		Logo:        h.Logo,
		Description: h.Description,
	}
	if h.EntID != nil && h.oldAppID != h.EntID {
		req.EntID = h.EntID
	}
	info, err = appcrud.UpdateSet(
		info.Update(),
		req,
	).Save(ctx)
	if err != nil {
		return err
	}
	if h.EntID == nil {
		h.EntID = &info.EntID
	}
	return nil
}

func (h *updateHandler) updateAppCtrl(ctx context.Context, tx *ent.Tx) error {
	if h.oldAppID == nil {
		return nil
	}
	info, err := tx.
		AppControl.
		Query().
		Where(
			entappctrl.AppID(*h.oldAppID),
			entappctrl.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	req := &ctrlcrud.Req{
		AppID:                    h.EntID,
		SignupMethods:            h.SignupMethods,
		ExtSigninMethods:         h.ExtSigninMethods,
		RecaptchaMethod:          h.RecaptchaMethod,
		KycEnable:                h.KycEnable,
		SigninVerifyEnable:       h.SigninVerifyEnable,
		InvitationCodeMust:       h.InvitationCodeMust,
		CreateInvitationCodeWhen: h.CreateInvitationCodeWhen,
		MaxTypedCouponsPerOrder:  h.MaxTypedCouponsPerOrder,
		Maintaining:              h.Maintaining,
		CouponWithdrawEnable:     h.CouponWithdrawEnable,
		CommitButtonTargets:      h.CommitButtonTargets,
		ResetUserMethod:          h.ResetUserMethod,
	}

	if info == nil {
		if _, err = ctrlcrud.CreateSet(
			tx.AppControl.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err = ctrlcrud.UpdateSet(info.Update(), req).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateBanApp(ctx context.Context, tx *ent.Tx) error {
	if h.Banned == nil {
		return nil
	}
	if h.oldAppID == nil {
		return nil
	}

	stm, err := banappcrud.SetQueryConds(
		tx.BanApp.Query(),
		&banappcrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.oldAppID},
		})
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if *h.Banned && info == nil {
		if _, err := banappcrud.CreateSet(
			tx.BanApp.Create(),
			&banappcrud.Req{
				AppID:   h.EntID,
				Message: h.BanMessage,
			},
		).Save(ctx); err != nil {
			return err
		}
	} else if !*h.Banned && info != nil {
		now := uint32(time.Now().Unix())
		if _, err := banappcrud.UpdateSet(
			tx.BanApp.UpdateOneID(info.ID),
			&banappcrud.Req{
				AppID:     h.EntID,
				EntID:     &info.EntID,
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) UpdateApp(ctx context.Context) (*npool.App, error) {
	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateApp(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateAppCtrl(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateBanApp(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetApp(ctx)
}
