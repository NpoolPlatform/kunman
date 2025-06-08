package displayname

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgooddisplaynamemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/name"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDisplayName(ctx context.Context) (*npool.DisplayName, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDisplayName(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetDisplayName(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid displayname")
	}

	if err := appgooddisplaynamemwcli.DeleteDisplayName(ctx, h.ID, h.EntID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
