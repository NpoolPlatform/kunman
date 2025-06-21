package chain

import (
	"context"

	chainmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/chain"
	chainmw "github.com/NpoolPlatform/kunman/middleware/chain/chain"
)

func (h *Handler) UpdateChain(ctx context.Context) (*chainmwpb.Chain, error) {
	handler, err := chainmw.NewHandler(
		ctx,
		chainmw.WithID(h.ID, true),
		chainmw.WithChainType(h.ChainType, true),
		chainmw.WithNativeUnit(h.NativeUnit, true),
		chainmw.WithAtomicUnit(h.AtomicUnit, true),
		chainmw.WithUnitExp(h.UnitExp, true),
		chainmw.WithENV(h.ENV, true),
		chainmw.WithChainID(h.ChainID, true),
		chainmw.WithNickname(h.Nickname, true),
		chainmw.WithGasType(h.GasType, true),
		chainmw.WithLogo(h.Logo, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.UpdateChain(ctx); err != nil {
		return nil, err
	}

	return h.GetChain(ctx)
}
