package fiat

import (
	"context"

	fiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	fiatmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
)

func (h *Handler) CreateFiat(ctx context.Context) (*fiatmwpb.Fiat, error) {
	handler, err := fiatmw.NewHandler(
		ctx,
		fiatmw.WithName(h.Name, true),
		fiatmw.WithUnit(h.Unit, true),
		fiatmw.WithLogo(h.Logo, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.CreateFiat(ctx)
}
