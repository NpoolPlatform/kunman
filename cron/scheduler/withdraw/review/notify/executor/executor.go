package executor

import (
	"context"
	"fmt"

	ledgerwithdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, withdraws interface{}, persistent, notif, done chan interface{}) error {
	_withdraws, ok := withdraws.([]*ledgerwithdrawmwpb.Withdraw)
	if !ok {
		return fmt.Errorf("invalid withdraws")
	}
	h := &withdrawReviewNotifyHandler{
		withdraws:  _withdraws,
		persistent: persistent,
		notif:      notif,
		done:       done,
	}
	return h.exec(ctx)
}
