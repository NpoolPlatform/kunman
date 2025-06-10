package default1

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/default"
	defaultmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/default"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateDefault(ctx context.Context) (*npool.Default, error) {
	if h.AppGoodID != nil {
		if err := h.CheckAppGood(ctx); err != nil {
			return nil, err
		}
	}

	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDefault(ctx); err != nil {
		return nil, err
	}

	defaultHandler, err := defaultmw.NewHandler(
		ctx,
		defaultmw.WithID(h.ID, true),
		defaultmw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := defaultHandler.UpdateDefault(ctx); err != nil {
		return nil, err
	}
	return h.GetDefault(ctx)
}
