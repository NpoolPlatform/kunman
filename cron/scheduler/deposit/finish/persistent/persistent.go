package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/finish/types"
	depositaccmw "github.com/NpoolPlatform/kunman/middleware/account/deposit"

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

	handler, err := depositaccmw.NewHandler(
		ctx,
		depositaccmw.WithID(&_account.ID, true),
		depositaccmw.WithAppID(&_account.AppID, true),
		depositaccmw.WithUserID(&_account.UserID, true),
		depositaccmw.WithCoinTypeID(&_account.CoinTypeID, true),
		depositaccmw.WithAccountID(&_account.AccountID, true),
		depositaccmw.WithLocked(&locked, true),
		depositaccmw.WithCollectingTID(&collectingID, true),
		depositaccmw.WithOutcoming(_account.CollectOutcoming, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateAccount(ctx); err != nil {
		return err
	}

	return nil
}
