package fiat

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	fiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createFiat(ctx context.Context, cli *ent.Client) error {
	if _, err := fiatcrud.CreateSet(
		cli.Fiat.Create(),
		&fiatcrud.Req{
			EntID: h.EntID,
			Name:  h.Name,
			Logo:  h.Logo,
			Unit:  h.Unit,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateFiat(ctx context.Context) (*npool.Fiat, error) {
	if h.Name == nil {
		return nil, fmt.Errorf("invalid fiatname")
	}
	if h.Unit == nil {
		return nil, fmt.Errorf("invalid fiatunit")
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
	if fiat != nil {
		return fiat, nil
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createFiat(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFiat(ctx)
}
