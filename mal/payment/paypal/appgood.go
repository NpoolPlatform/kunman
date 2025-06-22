package paypal

import (
	"context"

	appsubscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription"
)

type appGoodHandler struct {
	appSubscription *appsubscriptionmwpb.Subscription
}

func (cli *PaymentClient) GetAppGood(ctx context.Context) error {
	return nil
}
