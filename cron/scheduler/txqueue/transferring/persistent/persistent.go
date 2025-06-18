package persistent

import (
	"context"
	"fmt"

	txmwcli "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/transferring/types"
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

	if _, err := txmwcli.UpdateTx(ctx, &txmwpb.TxReq{
		ID:        &_tx.ID,
		State:     &_tx.NewTxState,
		ChainTxID: _tx.TxCID,
		Extra:     &_tx.TxExtra,
	}); err != nil {
		return err
	}

	return nil
}
