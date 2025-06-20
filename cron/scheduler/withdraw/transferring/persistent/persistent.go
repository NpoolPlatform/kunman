package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/transferring/types"
	withdrawmw "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, withdraw interface{}, reward, notif, done chan interface{}) error {
	_withdraw, ok := withdraw.(*types.PersistentWithdraw)
	if !ok {
		return fmt.Errorf("invalid withdraw")
	}

	defer asyncfeed.AsyncFeed(ctx, _withdraw, done)

	handler, err := withdrawmw.NewHandler(
		ctx,
		withdrawmw.WithID(&_withdraw.ID, true),
		withdrawmw.WithState(&_withdraw.NewWithdrawState, true),
		withdrawmw.WithChainTransactionID(&_withdraw.ChainTxID, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateWithdraw(ctx); err != nil {
		return err
	}

	return nil
}
