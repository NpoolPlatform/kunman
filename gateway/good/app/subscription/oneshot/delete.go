package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"
	apponeshotmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription/oneshot"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteOneShot(ctx context.Context) (*npool.AppOneShot, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkOneShot(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetOneShot(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid oneshot")
	}

	prHandler, err := apponeshotmw.NewHandler(
		ctx,
		apponeshotmw.WithID(h.ID, true),
		apponeshotmw.WithEntID(h.EntID, true),
		apponeshotmw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := prHandler.DeleteOneShot(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
