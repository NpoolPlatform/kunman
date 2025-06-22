package paypal

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/google/uuid"
)

type PaymentClient struct {
	config *Config

	OrderID   string
	ReturnURL string
	CancelURL string

	orderHandler *orderHandler
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

	if err := cli.GetOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return cli, nil
}

func WithOrderID(orderID string) func(context.Context, *PaymentClient) error {
	return func(ctx context.Context, h *PaymentClient) error {
		if _, err := uuid.Parse(orderID); err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = orderID
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
