package capacity

import (
	"context"

	capacitymw "github.com/NpoolPlatform/kunman/middleware/agi/capacity"
)

func (h *Handler) CountCapacities(ctx context.Context) (uint32, error) {
	handler, err := capacitymw.NewHandler(
		ctx,
		capacitymw.WithAppGoodID(h.AppGoodID, false),
	)
	if err != nil {
		return 0, err
	}

	return handler.CountCapacities(ctx)
}
