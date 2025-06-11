package simulate

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"
	simulatemw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental/simulate"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateSimulate(ctx context.Context) (*npool.Simulate, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSimulate(ctx); err != nil {
		return nil, err
	}

	simulateHandler, err := simulatemw.NewHandler(
		ctx,
		simulatemw.WithEntID(h.EntID, true),
		simulatemw.WithAppGoodID(h.AppGoodID, true),
		simulatemw.WithOrderUnits(h.OrderUnits, true),
		simulatemw.WithOrderDurationSeconds(h.OrderDurationSeconds, true),
	)
	if err != nil {
		return nil, err
	}

	if err := simulateHandler.UpdateSimulate(ctx); err != nil {
		return nil, err
	}
	return h.GetSimulate(ctx)
}
