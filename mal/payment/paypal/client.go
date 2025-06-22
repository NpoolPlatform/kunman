package paypal

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/google/uuid"
)

type PaymentClient struct {
	config *Config

	OrderID   *string
	ReturnURL string
	CancelURL string

	orderHandler   *orderHandler
	appGoodHandler *appGoodHandler

	// For single payment
	PaypalPaymentID string

	// For subscription plan
	AppGoodID *string

	// For created subscription
	PaypalPlanID string

	PaypalSubscriptionID string
}

func NewPaymentClient(ctx context.Context, options ...func(context.Context, *PaymentClient) error) (cli *PaymentClient, err error) {
	cli = &PaymentClient{}
	for _, opt := range options {
		if err := opt(ctx, cli); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	cli.config, err = LoadConfig()
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if cli.OrderID != nil {
		if err := cli.GetOrder(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if cli.AppGoodID != nil {
		if err := cli.GetAppGood(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	return cli, nil
}

func WithOrderID(orderID string) func(context.Context, *PaymentClient) error {
	return func(ctx context.Context, h *PaymentClient) error {
		if _, err := uuid.Parse(orderID); err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = &orderID
		return nil
	}
}

func WithAppGoodID(appGoodID string) func(context.Context, *PaymentClient) error {
	return func(ctx context.Context, h *PaymentClient) error {
		if _, err := uuid.Parse(appGoodID); err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodID = &appGoodID
		return nil
	}
}

func WithReturnURL(url string) func(context.Context, *PaymentClient) error {
	return func(ctx context.Context, h *PaymentClient) error {
		if url == "" {
			return wlog.Errorf("invalid return url")
		}
		h.ReturnURL = url
		return nil
	}
}

func WithCancelURL(url string) func(context.Context, *PaymentClient) error {
	return func(ctx context.Context, h *PaymentClient) error {
		if url == "" {
			return wlog.Errorf("invalid cancel url")
		}
		h.CancelURL = url
		return nil
	}
}

func WithPaypalPaymentID(s string) func(context.Context, *PaymentClient) error {
	return func(ctx context.Context, h *PaymentClient) error {
		if s == "" {
			return wlog.Errorf("invalid paymentid")
		}
		h.PaypalPaymentID = s
		return nil
	}
}

func WithPaypalPlanID(s string) func(context.Context, *PaymentClient) error {
	return func(ctx context.Context, h *PaymentClient) error {
		if s == "" {
			return wlog.Errorf("invalid planid")
		}
		h.PaypalPlanID = s
		return nil
	}
}

func WithPaypalSubscriptionID(s string) func(context.Context, *PaymentClient) error {
	return func(ctx context.Context, h *PaymentClient) error {
		if s == "" {
			return wlog.Errorf("invalid subscriptionid")
		}
		h.PaypalSubscriptionID = s
		return nil
	}
}
