package displaycolor

import (
	"context"

	appgooddisplaycolormwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/color"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/display/color"
	appgooddisplaycolormwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/color"
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
