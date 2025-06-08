package fiat

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	fiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateFiat(ctx context.Context, cli *ent.Client) error {
	if _, err := fiatcrud.UpdateSet(
		cli.Fiat.UpdateOneID(*h.ID),
		&fiatcrud.Req{
			Name: h.Name,
			Logo: h.Logo,
			Unit: h.Unit,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateFiat(ctx context.Context) (*npool.Fiat, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	// TODO: deduplicate

	h.Conds = &fiatcrud.Conds{
		Name: &cruder.Cond{Op: cruder.EQ, Val: *h.Name},
	}
	h.Offset = 0
	h.Limit = 2

	fiat, err := h.GetFiatOnly(ctx)
	if err != nil {
		return nil, err
	}
	if fiat != nil && fiat.ID != *h.ID {
		return nil, fmt.Errorf("fiat exist")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateFiat(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFiat(ctx)
}
