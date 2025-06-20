package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/created/types"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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

	state := basetypes.TxState_TxStateWait

	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithID(&_tx.ID, true),
		txmw.WithState(&state, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateTx(ctx); err != nil {
		return err
	}

	return nil
}
