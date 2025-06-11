package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	subscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription"
	subscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/subscription"

	"github.com/google/uuid"
)

func (h *Handler) AdminCreateSubscrption(ctx context.Context) (*subscriptionmwpb.Subscription, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.GoodID == nil {
		h.GoodID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := subscriptionmw.NewHandler(
		ctx,
		subscriptionmw.WithEntID(h.EntID, true),
		subscriptionmw.WithGoodID(h.GoodID, true),
		subscriptionmw.WithName(h.Name, true),
		subscriptionmw.WithDurationDisplayType(h.DurationDisplayType, true),
		subscriptionmw.WithDurationUnits(h.DurationUnits, false),
		subscriptionmw.WithDurationQuota(h.DurationQuota, true),
		subscriptionmw.WithDailyBonusQuota(h.DailyBonusQuota, false),
		subscriptionmw.WithUSDPrice(h.USDPrice, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.CreateSubscription(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetSubscription(ctx)
}
