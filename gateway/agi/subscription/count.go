package subscription

import (
	"context"

	subscriptionmw "github.com/NpoolPlatform/kunman/middleware/agi/subscription"
)

func (h *Handler) CountSubscriptions(ctx context.Context) (uint32, error) {
	handler, err := subscriptionmw.NewHandler(
		ctx,
		subscriptionmw.WithAppGoodID(h.AppGoodID, false),
	)
	if err != nil {
		return 0, err
	}

	return handler.CountSubscriptions(ctx)
}
