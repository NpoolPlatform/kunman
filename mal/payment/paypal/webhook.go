package paypal

import (
	"context"
	"encoding/json"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
)

type WebhookEvent struct {
	ID            string          `json:"id"`
	EventType     string          `json:"event_type"`
	Resource      json.RawMessage `json:"resource"`
	TransactionID string          `json:"transaction_id,omitempty"`
	PaymentAmount string          `json:"amount,omitempty"`
	Currency      string          `json:"currency,omitempty"`
}

func (cli *PaymentClient) onSubscriptionCreated(ctx context.Context, event *WebhookEvent) error {
	return nil
}

func (cli *PaymentClient) onSubscriptionActivated(ctx context.Context, event *WebhookEvent) error {
	return nil
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
		return cli.onSubscriptionCreated(ctx, event)
	case "BILLING.SUBSCRIPTION.ACTIVATED":
		return cli.onSubscriptionActivated(ctx, event)
	case "BILLING.SUBSCRIPTION.UPDATED":
		return cli.onSubscriptionUpdated(ctx, event)
	case "BILLING.SUBSCRIPTION.SUSPENDED":
		return cli.onSubscriptionSuspended(ctx, event)
	case "BILLING.SUBSCRIPTION.CANCELLED":
		return cli.onSubscriptionCancelled(ctx, event)
	default:
		return wlog.Errorf("Unknown event type: %s", event.EventType)
	}
}
