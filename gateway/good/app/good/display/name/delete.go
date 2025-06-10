package displayname

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
	appgooddisplaynamemw "github.com/NpoolPlatform/kunman/middleware/good/app/good/display/name"
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

	displayNameHandler, err := appgooddisplaynamemw.NewHandler(
		ctx,
		appgooddisplaynamemw.WithID(h.ID, true),
		appgooddisplaynamemw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := displayNameHandler.DeleteDisplayName(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
