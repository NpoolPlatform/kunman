package simulate

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"
	simulatemw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental/simulate"

	"github.com/google/uuid"
)

func (h *Handler) CreateSimulate(ctx context.Context) (*npool.Simulate, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := simulatemw.NewHandler(
		ctx,
		simulatemw.WithEntID(h.EntID, true),
		simulatemw.WithAppGoodID(h.AppGoodID, true),
		simulatemw.WithOrderUnits(h.OrderUnits, true),
		simulatemw.WithOrderDurationSeconds(h.OrderDurationSeconds, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateSimulate(ctx); err != nil {
		return nil, err
	}

	return h.GetSimulate(ctx)
}
