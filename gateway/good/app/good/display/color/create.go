package displaycolor

import (
	"context"

	appgooddisplaycolormwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/display/color"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
	appgooddisplaycolormwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/display/color"

	"github.com/google/uuid"
)

func (h *Handler) CreateDisplayColor(ctx context.Context) (*npool.DisplayColor, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := appgooddisplaycolormwcli.CreateDisplayColor(ctx, &appgooddisplaycolormwpb.DisplayColorReq{
		EntID:     h.EntID,
		AppGoodID: h.AppGoodID,
		Color:     h.Color,
		Index:     h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetDisplayColor(ctx)
}
