package paypal

import (
	"context"
	"encoding/json"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (cli *PaymentClient) onSubscriptionActivated(ctx context.Context, event *WebhookEvent) error {
	var resource SubscriptionActivatedResource
	if err := json.Unmarshal(event.Resource, &resource); err != nil {
		return wlog.WrapError(err)
	}

	appID, userID, orderID, err := CustomID2AppUserOrderID(resource.CustomID)
	if err != nil {
		return wlog.WrapError(err)
	}

	conds := &subscriptionordermwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: appID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: userID},
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: orderID},
	}
	handler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	subscriptionOrder, err := handler.GetSubscriptionOrderOnly(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	handler, err = subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithID(&subscriptionOrder.ID, true),
		subscriptionordermw.WithEntID(&subscriptionOrder.EntID, true),
		subscriptionordermw.WithOrderID(&subscriptionOrder.OrderID, true),
		subscriptionordermw.WithDealEventID(&event.ID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return handler.UpdateSubscriptionOrder(ctx)
}

func (cli *PaymentClient) onSubscriptionUpdated(ctx context.Context, event *WebhookEvent) error {
	return nil
}

func (cli *PaymentClient) onSubscriptionSuspended(ctx context.Context, event *WebhookEvent) error {
	return nil
}

func (cli *PaymentClient) onSubscriptionCancelled(ctx context.Context, event *WebhookEvent) error {
	return nil
}

func (cli *PaymentClient) OnWebhook(ctx context.Context, event *WebhookEvent) error {
	switch event.EventType {
	case "BILLING.SUBSCRIPTION.CREATED":
		// Ignore, the order is already created
	case "BILLING.SUBSCRIPTION.ACTIVATED":
		// Activated, means the subscription is paid, update the order deal event id then the state machine can continue
		return cli.onSubscriptionActivated(ctx, event)
	case "BILLING.SUBSCRIPTION.UPDATED":
		// Extension, should create new paid order
		return cli.onSubscriptionUpdated(ctx, event)
	case "BILLING.SUBSCRIPTION.SUSPENDED":
		// Do nothing, quota won't be added
	case "BILLING.SUBSCRIPTION.CANCELED":
		// Cancel subscription, then frontend can let user to subscribe again
		return cli.onSubscriptionCancelled(ctx, event)
	default:
		return wlog.Errorf("Unknown event type: %s", event.EventType)
	}
	return nil
}
