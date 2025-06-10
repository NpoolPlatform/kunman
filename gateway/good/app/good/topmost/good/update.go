package topmostgood

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
	topmostgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkTopMostGood(ctx); err != nil {
		return nil, err
	}

	goodHandler, err := topmostgoodmw.NewHandler(
		ctx,
		topmostgoodmw.WithID(h.ID, true),
		topmostgoodmw.WithEntID(h.EntID, true),
		topmostgoodmw.WithUnitPrice(h.UnitPrice, false),
		topmostgoodmw.WithDisplayIndex(h.DisplayIndex, false),
	)
	if err != nil {
		return nil, err
	}

	if err := goodHandler.UpdateTopMostGood(ctx); err != nil {
		return nil, err
	}
	return h.GetTopMostGood(ctx)
}
