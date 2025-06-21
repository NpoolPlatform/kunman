package appcoin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/coin"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
)

func (h *Handler) DeleteCoin(ctx context.Context) (*npool.Coin, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	// TODO: check appid / cointypeid / id

	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithID(h.ID, true),
		appcoinmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	deleteInfo, err := handler.DeleteCoin(ctx)
	if err != nil {
		return nil, err
	}

	if deleteInfo == nil {
		return nil, nil
	}

	info, err := h.GetCoinExt(ctx, deleteInfo)
	if err != nil {
		return nil, err
	}

	return info, nil
}
