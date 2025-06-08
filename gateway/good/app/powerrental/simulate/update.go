package simulate

import (
	"context"

	simulatemwcli "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental/simulate"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"
	simulatemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental/simulate"
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

	if err := simulatemwcli.UpdateSimulate(ctx, &simulatemwpb.SimulateReq{
		ID:                   h.ID,
		EntID:                h.EntID,
		AppGoodID:            h.AppGoodID,
		OrderUnits:           h.OrderUnits,
		OrderDurationSeconds: h.OrderDurationSeconds,
	}); err != nil {
		return nil, err
	}
	return h.GetSimulate(ctx)
}
