package appfiat

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/fiat"
	appfiatmw "github.com/NpoolPlatform/kunman/middleware/chain/app/fiat"
)

func (h *Handler) DeleteFiat(ctx context.Context) (*npool.Fiat, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	// TODO: check appid / fiatid / id

	handler, err := appfiatmw.NewHandler(
		ctx,
		appfiatmw.WithID(h.ID, true),
		appfiatmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	deleteInfo, err := handler.DeleteFiat(ctx)
	if err != nil {
		return nil, err
	}

	if deleteInfo == nil {
		return nil, nil
	}

	info, err := h.GetFiatExt(ctx, deleteInfo)
	if err != nil {
		return nil, err
	}

	return info, nil
}
