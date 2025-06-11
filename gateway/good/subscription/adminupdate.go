package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	subscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription"
	subscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/subscription"
)

func (h *Handler) AdminUpdateSubscrption(ctx context.Context) (*subscriptionmwpb.Subscription, error) {
	handler, err := subscriptionmw.NewHandler(
		ctx,
		subscriptionmw.WithID(h.ID, true),
		subscriptionmw.WithEntID(h.EntID, true),
		subscriptionmw.WithGoodID(h.GoodID, true),
		subscriptionmw.WithName(h.Name, false),
		subscriptionmw.WithDurationDisplayType(h.DurationDisplayType, false),
		subscriptionmw.WithDurationUnits(h.DurationUnits, false),
		subscriptionmw.WithDurationQuota(h.DurationQuota, false),
		subscriptionmw.WithDailyBonusQuota(h.DailyBonusQuota, false),
		subscriptionmw.WithUSDPrice(h.USDPrice, false),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.UpdateSubscription(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetSubscription(ctx)
}
