package displaycolor

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
	appgooddisplaycolormw "github.com/NpoolPlatform/kunman/middleware/good/app/good/display/color"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateDisplayColor(ctx context.Context) (*npool.DisplayColor, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDisplayColor(ctx); err != nil {
		return nil, err
	}

	displayColorHandler, err := appgooddisplaycolormw.NewHandler(
		ctx,
		appgooddisplaycolormw.WithEntID(h.EntID, true),
		appgooddisplaycolormw.WithAppGoodID(h.AppGoodID, true),
		appgooddisplaycolormw.WithColor(h.Color, true),
		appgooddisplaycolormw.WithIndex(func() *uint8 {
			if h.Index == nil {
				return nil
			}
			index := uint8(*h.Index)
			return &index
		}(), true),
	)
	if err != nil {
		return nil, err
	}

	if err := displayColorHandler.UpdateDisplayColor(ctx); err != nil {
		return nil, err
	}
	return h.GetDisplayColor(ctx)
}
