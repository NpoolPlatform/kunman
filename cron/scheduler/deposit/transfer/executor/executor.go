package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, account interface{}, persistent, notif, done chan interface{}) error {
	_account, ok := account.(*depositaccmwpb.Account)
	if !ok {
		return fmt.Errorf("invalid account")
	}

	h := &accountHandler{
		Account:    _account,
		persistent: persistent,
		notif:      notif,
		done:       done,
	}
	return h.exec(ctx)
}
