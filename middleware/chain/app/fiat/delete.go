package appfiat

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/fiat"
	appfiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/fiat"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteAppFiat(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	info, err := appfiatcrud.UpdateSet(
		tx.AppFiat.UpdateOneID(*h.ID),
		&appfiatcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.AppID = &info.AppID
	h.FiatID = &info.FiatID

	return nil
}

func (h *Handler) DeleteFiat(ctx context.Context) (*npool.Fiat, error) {
	info, err := h.GetFiat(ctx)
	if err != nil {
		return nil, err
	}

	handler := &deleteHandler{
		Handler: h,
	}

	h.ID = &info.ID

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppFiat(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
