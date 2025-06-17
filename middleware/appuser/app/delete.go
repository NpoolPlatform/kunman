package app

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	entappctrl "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appcontrol"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteApp(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		App.
		UpdateOneID(*h.ID).
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return err
	}
	h.EntID = &info.EntID
	return nil
}

func (h *deleteHandler) deleteAppCtrl(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppControl.
		Query().
		Where(
			entappctrl.AppID(*h.EntID),
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

func (h *Handler) DeleteApp(ctx context.Context) (info *npool.App, err error) {
	info, err = h.GetApp(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteApp(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppCtrl(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
