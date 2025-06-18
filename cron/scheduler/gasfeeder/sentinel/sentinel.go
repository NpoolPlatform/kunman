package sentinel

import (
	"context"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/gasfeeder/types"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/google/uuid"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		handler, err := coinmw.NewHandler(
			ctx,
			coinmw.WithConds(&coinmwpb.Conds{}),
			coinmw.WithOffset(offset),
			coinmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		coins, _, err := handler.GetCoins(ctx)
		if err != nil {
			return err
		}
		if len(coins) == 0 {
			return nil
		}

		for _, coin := range coins {
			if _, err := uuid.Parse(coin.FeeCoinTypeID); err != nil {
				continue
			}
			if coin.FeeCoinTypeID == uuid.Nil.String() {
				continue
			}
			if coin.FeeCoinTypeID == coin.EntID {
				continue
			}
			cancelablefeed.CancelableFeed(ctx, coin, exec)
		}

		offset += limit
	}
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if coin, ok := ent.(*types.PersistentCoin); ok {
		return coin.EntID
	}
	return ent.(*coinmwpb.Coin).EntID
}
