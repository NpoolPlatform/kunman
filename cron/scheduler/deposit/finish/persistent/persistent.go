package persistent

import (
	"context"
	"fmt"
	"time"

	depositaccmwcli "github.com/NpoolPlatform/account-middleware/pkg/client/deposit"
	accountlock "github.com/NpoolPlatform/account-middleware/pkg/lock"
	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/finish/types"

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

	if err := accountlock.Lock(_account.AccountID); err != nil {
		return err
	}
	defer func() {
		_ = accountlock.Unlock(_account.AccountID) //nolint
	}()

	scannableAt := uint32(time.Now().Unix() + timedef.SecondsPerHour)
	locked := false
	collectingID := uuid.Nil.String()
	if _, err := depositaccmwcli.UpdateAccount(ctx, &depositaccmwpb.AccountReq{
		ID:            &_account.ID,
		AppID:         &_account.AppID,
		UserID:        &_account.UserID,
		CoinTypeID:    &_account.CoinTypeID,
		AccountID:     &_account.AccountID,
		Locked:        &locked,
		CollectingTID: &collectingID,
		ScannableAt:   &scannableAt,
		Outcoming:     _account.CollectOutcoming,
	}); err != nil {
		return err
	}

	return nil
}
