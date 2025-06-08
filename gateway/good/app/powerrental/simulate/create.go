package simulate

import (
	"context"

	simulatemwcli "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental/simulate"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"
	simulatemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental/simulate"

	"github.com/google/uuid"
)

func (h *Handler) CreateSimulate(ctx context.Context) (*npool.Simulate, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := simulatemwcli.CreateSimulate(ctx, &simulatemwpb.SimulateReq{
		EntID:                h.EntID,
		AppGoodID:            h.AppGoodID,
		OrderUnits:           h.OrderUnits,
		OrderDurationSeconds: h.OrderDurationSeconds,
	}); err != nil {
		return nil, err
	}

	return h.GetSimulate(ctx)
}
