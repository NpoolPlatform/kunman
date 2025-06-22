package paypal

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	fiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	fiatmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
)

type orderHandler struct {
	subscriptionOrder *subscriptionordermwpb.SubscriptionOrder
	// We may have other order types

	fiat *fiatmwpb.Fiat
}

func (cli *PaymentClient) GetOrder(ctx context.Context) error {
	handler := &orderHandler{}

	_orderHandler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithOrderID(&cli.OrderID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	order, err := _orderHandler.GetSubscriptionOrder(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if len(order.PaymentFiats) != 1 {
		return wlog.Errorf("invalid paymentfiats")
	}
	if order.PaymentFiats[0].PaymentChannel != types.FiatPaymentChannel_PaymentChannelPaypal {
		return wlog.Errorf("invalid paymentchannel")
	}

	fiatHandler, err := fiatmw.NewHandler(
		ctx,
		fiatmw.WithEntID(&order.PaymentFiats[0].FiatID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	fiat, err := fiatHandler.GetFiat(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	// TODO: check payment state

	handler.subscriptionOrder = order
	handler.fiat = fiat
	cli.orderHandler = handler

	return nil
}

func (h *orderHandler) FiatPaymentCurrency() string {
	return h.fiat.Unit
}

func (h *orderHandler) FiatPaymentAmount() string {
	return h.subscriptionOrder.PaymentFiats[0].Amount
}
