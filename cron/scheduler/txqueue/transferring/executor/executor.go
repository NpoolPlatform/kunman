package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, tx interface{}, persistent, notif, done chan interface{}) error {
	_tx, ok := tx.(*txmwpb.Tx)
	if !ok {
		return fmt.Errorf("invalid tx")
	}
	h := &txHandler{
		Tx:         _tx,
		persistent: persistent,
		notif:      notif,
		done:       done,
	}
	return h.exec(ctx)
}
