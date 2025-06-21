package appfiat

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/fiat"
	appfiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/fiat"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateAppFiat(ctx context.Context, tx *ent.Tx) error {
	info, err := appfiatcrud.UpdateSet(
		tx.AppFiat.UpdateOneID(*h.ID),
		&appfiatcrud.Req{
			Name:         h.Name,
			DisplayNames: h.DisplayNames,
			Logo:         h.Logo,
			Disabled:     h.Disabled,
			Display:      h.Display,
			DisplayIndex: h.DisplayIndex,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.AppID = &info.AppID
	h.FiatID = &info.FiatID

	return nil
}

func (h *Handler) UpdateFiat(ctx context.Context) (*npool.Fiat, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppFiat(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFiat(ctx)
}
