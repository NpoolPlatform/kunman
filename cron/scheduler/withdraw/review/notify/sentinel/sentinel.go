package sentinel

import (
	"context"
	"fmt"
	"math"
	"os"
	"time"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ledgerwithdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	ledgerwithdrawmw "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type handler struct {
	ID             string
	nextNotifyAt   uint32
	notifyInterval uint32
}

func NewSentinel() basesentinel.Scanner {
	_interval := timedef.SecondsPerDay
	if interval, err := time.ParseDuration(
		fmt.Sprintf("%vh", os.Getenv("ENV_WITHDRAW_REVIEW_NOTIFY_INTERVAL_HOURS"))); err == nil && math.Round(interval.Seconds()) > 0 {
		_interval = int(math.Round(interval.Seconds()))
	}
	return &handler{
		ID:             uuid.NewString(),
		nextNotifyAt:   uint32((int(time.Now().Unix()) + _interval) / _interval * _interval),
		notifyInterval: uint32(_interval),
	}
}

func (h *handler) scanWithdraws(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	withdraws := []*ledgerwithdrawmwpb.Withdraw{}

	conds := &ledgerwithdrawmwpb.Conds{
		State: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ledgertypes.WithdrawState_Reviewing)},
	}

	for {
		handler, err := ledgerwithdrawmw.NewHandler(
			ctx,
			ledgerwithdrawmw.WithConds(conds),
			ledgerwithdrawmw.WithOffset(offset),
			ledgerwithdrawmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		_withdraws, _, err := handler.GetWithdraws(ctx)
		if err != nil {
			return err
		}
		if len(_withdraws) == 0 {
			break
		}
		withdraws = append(withdraws, _withdraws...)
		offset += limit
	}
	if len(withdraws) > 0 {
		cancelablefeed.CancelableFeed(ctx, withdraws, exec)
	}
	return nil
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	if uint32(time.Now().Unix()) < h.nextNotifyAt {
		return nil
	}
	if err := h.scanWithdraws(ctx, exec); err != nil {
		return err
	}
	h.nextNotifyAt = (uint32(time.Now().Unix()) + h.notifyInterval) / h.notifyInterval * h.notifyInterval
	return nil
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return h.scanWithdraws(ctx, exec)
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	return uuid.Nil.String()
}
