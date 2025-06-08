package capacity

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/capacity"
	capacitymw "github.com/NpoolPlatform/kunman/middleware/agi/capacity"
)

func (h *Handler) DeleteCapacity(ctx context.Context) (*npool.Capacity, error) {
	handler, err := capacitymw.NewHandler(
		ctx,
		capacitymw.WithID(h.ID, true),
		capacitymw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.DeleteCapacity(ctx); err != nil {
		return nil, err
	}

	return handler.GetCapacity(ctx)
}
