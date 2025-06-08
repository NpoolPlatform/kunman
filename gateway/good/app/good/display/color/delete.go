package displaycolor

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgooddisplaycolormwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/color"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDisplayColor(ctx context.Context) (*npool.DisplayColor, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDisplayColor(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetDisplayColor(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid displaycolor")
	}

	if err := appgooddisplaycolormwcli.DeleteDisplayColor(ctx, h.ID, h.EntID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
