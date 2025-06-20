package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, coin interface{}, persistent, notif, done chan interface{}) error {
	_coin, ok := coin.(*coinmwpb.Coin)
	if !ok {
		return fmt.Errorf("invalid coin")
	}

	h := &coinHandler{
		Coin:       _coin,
		persistent: persistent,
		notif:      notif,
		done:       done,
	}
	return h.exec(ctx)
}
