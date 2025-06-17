package oneshot

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"
	apponeshotmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription/oneshot"
)

// TODO: check start mode with power rental start mode

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateOneShot(ctx context.Context) (*npool.AppOneShot, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkOneShot(ctx); err != nil {
		return nil, err
	}

	prHandler, err := apponeshotmw.NewHandler(
		ctx,
		apponeshotmw.WithID(h.ID, true),
		apponeshotmw.WithEntID(h.EntID, true),
		apponeshotmw.WithAppGoodID(h.AppGoodID, true),
		apponeshotmw.WithName(h.Name, false),
		apponeshotmw.WithBanner(h.Banner, false),
		// apponeshotmw.WithEnableSetCommission(h.EnableSetCommission, false),
		apponeshotmw.WithUSDPrice(h.USDPrice, false),
	)
	if err != nil {
		return nil, err
	}

	if err := prHandler.UpdateOneShot(ctx); err != nil {
		return nil, err
	}
	return h.GetOneShot(ctx)
}
