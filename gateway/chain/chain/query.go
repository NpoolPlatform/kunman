package chain

import (
	"context"
	"fmt"

	chainmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/chain"
	chainmw "github.com/NpoolPlatform/kunman/middleware/chain/chain"
)

func (h *Handler) GetChains(ctx context.Context) ([]*chainmwpb.Chain, uint32, error) {
	handler, err := chainmw.NewHandler(
		ctx,
		chainmw.WithConds(&chainmwpb.Conds{}),
		chainmw.WithOffset(h.Offset),
		chainmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetChains(ctx)
}

func (h *Handler) GetChain(ctx context.Context) (*chainmwpb.Chain, error) {
	handler, err := chainmw.NewHandler(
		ctx,
		chainmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	chain, err := handler.GetChain(ctx)
	if err != nil {
		return nil, err
	}
	if chain == nil {
		return nil, fmt.Errorf("invalid chain")
	}

	return chain, nil
}
