package capacity

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/capacity"
	capacitymw "github.com/NpoolPlatform/kunman/middleware/agi/capacity"

	"github.com/google/uuid"
)

func (h *Handler) CreateCapacity(ctx context.Context) (*npool.Capacity, error) {
	handler, err := capacitymw.NewHandler(
		ctx,
		capacitymw.WithEntID(func() *string { _uid := uuid.NewString(); return &_uid }(), true),
		capacitymw.WithAppGoodID(h.AppGoodID, true),
		capacitymw.WithCapacityKey(h.CapacityKey, true),
		capacitymw.WithValue(h.Value, true),
		capacitymw.WithDescription(h.Description, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateCapacity(ctx); err != nil {
		return nil, err
	}

	return handler.GetCapacity(ctx)
}
