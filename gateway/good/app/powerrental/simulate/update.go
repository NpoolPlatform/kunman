package simulate

import (
	"context"

	simulatemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental/simulate"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental/simulate"
	simulatemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
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
