package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/transferring/types"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, tx interface{}, reward, notif, done chan interface{}) error {
	_tx, ok := tx.(*types.PersistentTx)
	if !ok {
		return fmt.Errorf("invalid tx")
	}

	defer asyncfeed.AsyncFeed(ctx, _tx, done)

	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithID(&_tx.ID, true),
		txmw.WithState(&_tx.NewTxState, true),
		txmw.WithChainTxID(_tx.TxCID, true),
		txmw.WithExtra(&_tx.TxExtra, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateTx(ctx); err != nil {
		return err
	}

	return nil
}
