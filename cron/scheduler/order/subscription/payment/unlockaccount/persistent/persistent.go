package persistent

import (
	"context"
	"fmt"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/unlockaccount/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	paymentaccmw "github.com/NpoolPlatform/kunman/middleware/account/payment"
	agisubscriptionmw "github.com/NpoolPlatform/kunman/middleware/agi/subscription"
	subscriptionquotamw "github.com/NpoolPlatform/kunman/middleware/agi/subscription/quota"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	state := ordertypes.OrderState_OrderStatePaid

	handler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithID(&order.ID, true),
		subscriptionordermw.WithOrderState(&state, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdateSubscriptionOrder(ctx)
}

func (p *handler) withUnlockPaymentAccount(ctx context.Context, order *types.PersistentOrder) error {
	// TODO: use UpdateAccounts in future
	for _, id := range order.PaymentAccountIDs {
		locked := false

		handler, err := paymentaccmw.NewHandler(
			ctx,
			paymentaccmw.WithID(&id, true),
			paymentaccmw.WithLocked(&locked, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.UpdateAccount(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (p *handler) withUpdateSubscription(ctx context.Context, order *types.PersistentOrder) error {
	handler, err := agisubscriptionmw.NewHandler(
		ctx,
		agisubscriptionmw.WithID(&order.ID, true),
		agisubscriptionmw.WithPermanentQuota(&order.LifeSeconds, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdateSubscription(ctx)
}

func (p *handler) withCreateQuota(ctx context.Context, order *types.PersistentOrder) error {
	expiredAt := uint32(time.Now().Unix()) + order.LifeSeconds

	handler, err := subscriptionquotamw.NewHandler(
		ctx,
		subscriptionquotamw.WithAppID(&order.AppID, true),
		subscriptionquotamw.WithUserID(&order.UserID, true),
		subscriptionquotamw.WithQuota(&order.OrderQuota, true),
		subscriptionquotamw.WithExpiredAt(&expiredAt, true),
	)
	if err != nil {
		return err
	}

	return handler.CreateQuota(ctx)
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, notif)
	defer asyncfeed.AsyncFeed(ctx, _order, reward)

	if err := p.withUpdateOrderState(ctx, _order); err != nil {
		return err
	}
	if err := p.withUnlockPaymentAccount(ctx, _order); err != nil {
		return err
	}
	if _order.LifeSeconds == 0 {
		if err := p.withUpdateSubscription(ctx, _order); err != nil {
			return err
		}
	} else {
		if err := p.withCreateQuota(ctx, _order); err != nil {
			return err
		}
	}

	return nil
}
