package subscription

import (
	"context"

	subscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription"
	subscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/subscription"
)

func (h *Handler) CountSubscriptions(ctx context.Context) (uint32, error) {
	handler, err := subscriptionmw.NewHandler(
		ctx,
		subscriptionmw.WithConds(&subscriptionmwpb.Conds{}),
	)
	if err != nil {
		return 0, err
	}

	return handler.CountSubscriptions(ctx)
}
