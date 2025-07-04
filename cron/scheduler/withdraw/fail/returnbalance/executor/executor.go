package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, withdraw interface{}, persistent, notif, done chan interface{}) error {
	_withdraw, ok := withdraw.(*withdrawmwpb.Withdraw)
	if !ok {
		return fmt.Errorf("invalid withdraw")
	}

	h := &withdrawHandler{
		Withdraw:   _withdraw,
		persistent: persistent,
		done:       done,
	}
	return h.exec(ctx)
}
