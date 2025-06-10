package displaycolor

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
	appgooddisplaycolormw "github.com/NpoolPlatform/kunman/middleware/good/app/good/display/color"

	"github.com/google/uuid"
)

func (h *Handler) CreateDisplayColor(ctx context.Context) (*npool.DisplayColor, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := appgooddisplaycolormw.NewHandler(
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

	if err := handler.CreateDisplayColor(ctx); err != nil {
		return nil, err
	}
	return h.GetDisplayColor(ctx)
}
