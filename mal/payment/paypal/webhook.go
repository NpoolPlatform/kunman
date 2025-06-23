package paypal

import (
	"context"
	"encoding/json"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	agisubscriptionmwpb "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/subscription"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	fiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	fiatcurrencymwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	agisubscriptionmw "github.com/NpoolPlatform/kunman/middleware/agi/subscription"
	fiatmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
	fiatcurrencymw "github.com/NpoolPlatform/kunman/middleware/chain/fiat/currency"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
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

	orderConds := &subscriptionordermwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: appID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: userID},
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: orderID},
	}
	orderHandler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithConds(orderConds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	subscriptionOrder, err := orderHandler.GetSubscriptionOrderOnly(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	orderHandler, err = subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithID(&subscriptionOrder.ID, true),
		subscriptionordermw.WithEntID(&subscriptionOrder.EntID, true),
		subscriptionordermw.WithOrderID(&subscriptionOrder.OrderID, true),
		subscriptionordermw.WithDealEventID(&event.ID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	if err := orderHandler.UpdateSubscriptionOrder(ctx); err != nil {
		return wlog.WrapError(err)
	}

	agiConds := &agisubscriptionmwpb.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: appID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: userID},
	}
	agiHandler, err := agisubscriptionmw.NewHandler(
		ctx,
		agisubscriptionmw.WithConds(agiConds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	agiSubscription, err := agiHandler.GetSubscriptionOnly(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if agiSubscription != nil {
		if agiSubscription.ActivatedEventID != "" && agiSubscription.ActivatedEventID != event.ID {
			return wlog.Errorf("Invalid activated event id")
		}
		if agiSubscription.SubscriptionID != resource.ID {
			return wlog.Errorf("invalid subscriptionid")
		}
		return nil
	}

	activatedAt, err := time.Parse(resource.BillingInfo.LastPayment.Time, "2024-05-01T00:00:00Z")
	if err != nil {
		return wlog.WrapError(err)
	}
	activatedAtUnix := uint32(activatedAt.Unix())

	agiHandler, err = agisubscriptionmw.NewHandler(
		ctx,
		agisubscriptionmw.WithID(&agiSubscription.ID, true),
		agisubscriptionmw.WithEntID(&agiSubscription.EntID, true),
		agisubscriptionmw.WithActivatedAt(&activatedAtUnix, true),
		agisubscriptionmw.WithLastPaymentAt(&activatedAtUnix, true),
		agisubscriptionmw.WithActivatedEventID(&event.ID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return agiHandler.UpdateSubscription(ctx)
}

func (cli *PaymentClient) onSubscriptionUpdated(ctx context.Context, event *WebhookEvent) error {
	var resource SubscriptionUpdatedResource
	if err := json.Unmarshal(event.Resource, &resource); err != nil {
		return wlog.WrapError(err)
	}

	appID, userID, orderID, err := CustomID2AppUserOrderID(resource.CustomID)
	if err != nil {
		return wlog.WrapError(err)
	}

	orderConds := &subscriptionordermwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: appID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: userID},
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: orderID},
	}
	orderHandler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithConds(orderConds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	subscriptionOrder, err := orderHandler.GetSubscriptionOrderOnly(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if subscriptionOrder == nil {
		return wlog.Errorf("invalid subscriptionorder")
	}

	agiConds := &agisubscriptionmwpb.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: appID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: userID},
	}
	agiHandler, err := agisubscriptionmw.NewHandler(
		ctx,
		agisubscriptionmw.WithConds(agiConds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	agiSubscription, err := agiHandler.GetSubscriptionOnly(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if agiSubscription == nil {
		return wlog.Errorf("Invalid agisubscription")
	}

	lastPaymentAt, err := time.Parse(resource.BillingInfo.LastPayment.Time, "2024-05-01T00:00:00Z")
	if err != nil {
		return wlog.WrapError(err)
	}
	if uint32(lastPaymentAt.Unix()) < agiSubscription.LastPaymentAt {
		return nil
	}
	if event.ID == agiSubscription.LastUpdatedEventID {
		return nil
	}

	if agiSubscription.SubscriptionID != resource.ID {
		return wlog.Errorf("invalid subscriptionid")
	}

	fiatConds := &fiatmwpb.Conds{
		Unit: &basetypes.StringVal{Op: cruder.EQ, Value: resource.BillingInfo.LastPayment.Amount.CurrencyCode},
	}
	fiatHandler, err := fiatmw.NewHandler(
		ctx,
		fiatmw.WithConds(fiatConds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	fiat, err := fiatHandler.GetFiatOnly(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if fiat == nil {
		return wlog.Errorf("invalid fiat")
	}

	currencyConds := &fiatcurrencymwpb.Conds{
		FiatID: &basetypes.StringVal{Op: cruder.EQ, Value: fiat.EntID},
	}
	currencyHandler, err := fiatcurrencymw.NewHandler(
		ctx,
		fiatcurrencymw.WithConds(currencyConds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	currency, err := currencyHandler.GetCurrencyOnly(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if currency == nil {
		return wlog.Errorf("invalid currency")
	}

	currencyAmount, err := decimal.NewFromString(currency.MarketValueLow)
	if err != nil {
		return wlog.WrapError(err)
	}
	paymentAmount, err := decimal.NewFromString(resource.BillingInfo.LastPayment.Amount.Value)
	if err != nil {
		return wlog.WrapError(err)
	}

	paymentAmountUSDStr := paymentAmount.Div(currencyAmount).String()

	// Create new order
	orderHandler, err = subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithAppID(&appID, true),
		subscriptionordermw.WithUserID(&userID, true),
		subscriptionordermw.WithGoodID(&subscriptionOrder.GoodID, true),
		subscriptionordermw.WithGoodType(&subscriptionOrder.GoodType, true),
		subscriptionordermw.WithAppGoodID(&subscriptionOrder.AppGoodID, true),
		subscriptionordermw.WithPaymentFiats(func() []*paymentmwpb.PaymentFiatReq {
			return []*paymentmwpb.PaymentFiatReq{
				&paymentmwpb.PaymentFiatReq{
					FiatID:         &fiat.EntID,
					Amount:         &resource.BillingInfo.LastPayment.Amount.Value,
					PaymentChannel: ordertypes.FiatPaymentChannel_PaymentChannelPaypal.Enum(),
					USDCurrency:    &currency.MarketValueLow,
				},
			}
		}(), true),
		subscriptionordermw.WithCreateMethod(ordertypes.OrderCreateMethod_OrderCreatedBySubscriptionExtension.Enum(), true),
		subscriptionordermw.WithOrderType(ordertypes.OrderType_Normal.Enum(), true),
		subscriptionordermw.WithPaymentAmountUSD(&paymentAmountUSDStr, true),
		subscriptionordermw.WithGoodValueUSD(&paymentAmountUSDStr, true),
		subscriptionordermw.WithDealEventID(&event.ID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return orderHandler.CreateSubscriptionOrder(ctx)
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
