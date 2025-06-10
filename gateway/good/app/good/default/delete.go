package default1

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/default"
	defaultmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/default"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDefault(ctx context.Context) (*npool.Default, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDefault(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetDefault(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid default")
	}

	defaultHandler, err := defaultmw.NewHandler(
		ctx,
		defaultmw.WithID(h.ID, true),
		defaultmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := defaultHandler.DeleteDefault(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
