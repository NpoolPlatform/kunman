package simulate

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"
	simulatemw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental/simulate"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteSimulate(ctx context.Context) (*npool.Simulate, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSimulate(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetSimulate(ctx)
	if err != nil {
		return nil, err
	}

	simulateHandler, err := simulatemw.NewHandler(
		ctx,
		simulatemw.WithID(h.ID, true),
		simulatemw.WithEntID(h.EntID, true),
		simulatemw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := simulateHandler.DeleteSimulate(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
