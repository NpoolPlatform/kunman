package displaycolor

import (
	"context"

	appgooddisplaycolormwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/display/color"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
	appgooddisplaycolormwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/display/color"
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
	if err := appgooddisplaycolormwcli.UpdateDisplayColor(ctx, &appgooddisplaycolormwpb.DisplayColorReq{
		ID:    h.ID,
		EntID: h.EntID,
		Color: h.Color,
		Index: h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetDisplayColor(ctx)
}
