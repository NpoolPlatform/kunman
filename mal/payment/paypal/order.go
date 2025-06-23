package paypal

import (
	"context"
	"fmt"
	"strings"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	fiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	fiatmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"

	"github.com/shopspring/decimal"
)

type orderHandler struct {
	subscriptionOrder *subscriptionordermwpb.SubscriptionOrder
	// We may have other order types

	fiat *fiatmwpb.Fiat
	user *usermwpb.User
}

func (cli *PaymentClient) GetOrder(ctx context.Context) error {
	handler := &orderHandler{}

	_orderHandler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithOrderID(cli.OrderID, true),
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

	userHandler, err := usermw.NewHandler(
		ctx,
		usermw.WithAppID(&order.AppID, true),
		usermw.WithEntID(&order.UserID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	user, err := userHandler.GetUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	handler.subscriptionOrder = order
	handler.fiat = fiat
	handler.user = user

	cli.orderHandler = handler

	return nil
}

func (h *orderHandler) FiatPaymentCurrency() string {
	return h.fiat.Unit
}

func (h *orderHandler) FiatPaymentAmount() (string, error) {
	amount, err := decimal.NewFromString(h.subscriptionOrder.PaymentFiats[0].Amount)
	if err != nil {
		return "", wlog.WrapError(err)
	}

	return amount.Round(2).String(), nil
}
func (h *orderHandler) Paid() bool {
	return h.subscriptionOrder.PaymentFiats[0].ChannelPaymentID != ""
}

func (h *orderHandler) CustomID() string {
	return fmt.Sprintf("%v@%v@%v", h.user.EntID, h.user.AppID, h.subscriptionOrder.OrderID)
}

func (h *orderHandler) GivenName() string {
	if h.user.FirstName == "" {
		return "Cute"
	}
	return h.user.FirstName
}

func (h *orderHandler) Surname() string {
	if h.user.LastName == "" {
		return "User"
	}
	return h.user.LastName
}

func (h *orderHandler) EmailAddress() string {
	return h.user.EmailAddress
}

func (h *orderHandler) CountryCode() string {
	return h.user.CountryCode
}

func (h *orderHandler) NationalNumber() string {
	if strings.HasPrefix(h.user.PhoneNO, h.user.CountryCode) {
		return h.user.PhoneNO[len(h.user.CountryCode):]
	}
	return h.user.PhoneNO
}
