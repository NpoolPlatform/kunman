package coin

import (
	"context"

	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
)

func (h *Handler) CreateCoin(ctx context.Context) (*coinmwpb.Coin, error) {
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithName(h.Name, true),
		coinmw.WithUnit(h.Unit, true),
		coinmw.WithENV(h.ENV, true),
		coinmw.WithChainType(h.ChainType, true),
		coinmw.WithChainNativeUnit(h.ChainNativeUnit, true),
		coinmw.WithChainAtomicUnit(h.ChainAtomicUnit, true),
		coinmw.WithChainUnitExp(h.ChainUnitExp, true),
		coinmw.WithGasType(h.GasType, true),
		coinmw.WithChainID(h.ChainID, true),
		coinmw.WithChainNickname(h.ChainNickname, true),
		coinmw.WithChainNativeCoinName(h.ChainNativeCoinName, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.CreateCoin(ctx)
}
