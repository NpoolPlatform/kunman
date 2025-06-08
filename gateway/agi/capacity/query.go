package capacity

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/capacity"
	capacitymw "github.com/NpoolPlatform/kunman/middleware/agi/capacity"
)

func (h *Handler) GetCapacity(ctx context.Context) (*npool.Capacity, error) {
	if h.ID == nil && h.EntID == nil {
		return nil, wlog.Errorf("invalid id")
	}

	handler, err := capacitymw.NewHandler(
		ctx,
		capacitymw.WithID(h.ID, false),
		capacitymw.WithEntID(h.EntID, false),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetCapacity(ctx)
}

func (h *Handler) GetCapacities(ctx context.Context) ([]*npool.Capacity, error) {
	handler, err := capacitymw.NewHandler(
		ctx,
		capacitymw.WithAppGoodID(h.AppGoodID, false),
		capacitymw.WithOffset(h.Offset),
		capacitymw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetCapacities(ctx)
}
