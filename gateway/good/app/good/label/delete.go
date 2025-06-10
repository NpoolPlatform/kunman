package label

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/label"
	appgoodlabelmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/label"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteLabel(ctx context.Context) (*npool.Label, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkLabel(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetLabel(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid label")
	}

	labelHandler, err := appgoodlabelmw.NewHandler(
		ctx,
		appgoodlabelmw.WithID(h.ID, true),
		appgoodlabelmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := labelHandler.DeleteLabel(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
