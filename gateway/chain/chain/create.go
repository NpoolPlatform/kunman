package chain

import (
	"context"

	chainmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/chain"
	chainmw "github.com/NpoolPlatform/kunman/middleware/chain/chain"

	"github.com/google/uuid"
)

func (h *Handler) CreateChain(ctx context.Context) (*chainmwpb.Chain, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := chainmw.NewHandler(
		ctx,
		chainmw.WithEntID(h.EntID, true),
		chainmw.WithChainType(h.ChainType, true),
		chainmw.WithNativeUnit(h.NativeUnit, true),
		chainmw.WithAtomicUnit(h.AtomicUnit, true),
		chainmw.WithUnitExp(h.UnitExp, true),
		chainmw.WithENV(h.ENV, true),
		chainmw.WithChainID(h.ChainID, true),
		chainmw.WithNickname(h.Nickname, true),
		chainmw.WithGasType(h.GasType, true),
	)
	if err != nil {
		return nil, err
	}

	err = handler.CreateChain(ctx)
	if err != nil {
		return nil, err
	}

	return h.GetChain(ctx)
}
