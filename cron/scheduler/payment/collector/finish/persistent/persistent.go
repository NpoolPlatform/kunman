package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/collector/finish/types"
	paymentaccountmw "github.com/NpoolPlatform/kunman/middleware/account/payment"

	"github.com/google/uuid"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, account interface{}, reward, notif, done chan interface{}) error {
	_account, ok := account.(*types.PersistentAccount)
	if !ok {
		return fmt.Errorf("invalid account")
	}

	defer asyncfeed.AsyncFeed(ctx, _account, done)

	locked := false
	collectingID := uuid.Nil.String()

	handler, err := paymentaccountmw.NewHandler(
		ctx,
		paymentaccountmw.WithID(&_account.ID, true),
		paymentaccountmw.WithCoinTypeID(&_account.CoinTypeID, true),
		paymentaccountmw.WithAccountID(&_account.AccountID, true),
		paymentaccountmw.WithLocked(&locked, true),
		paymentaccountmw.WithCollectingTID(&collectingID, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateAccount(ctx); err != nil {
		return err
	}

	return nil
}
