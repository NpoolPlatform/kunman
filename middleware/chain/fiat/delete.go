package fiat

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	fiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteFiat(ctx context.Context, cli *ent.Client) error {
	if _, err := fiatcrud.UpdateSet(
		cli.Fiat.UpdateOneID(*h.ID),
		&fiatcrud.Req{
			DeletedAt: func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteFiat(ctx context.Context) (*npool.Fiat, error) {
	if h.ID == nil && h.EntID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	fiat, err := h.GetFiat(ctx)
	if err != nil {
		return nil, err
	}

	h.ID = &fiat.ID
	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.deleteFiat(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return fiat, nil
}
