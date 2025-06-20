package sentinel

import (
	"context"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/checkhashrate/types"
	goodbasepb "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) scanPowerRentals(ctx context.Context, state goodbasepb.GoodState, goodType goodbasepb.GoodType, stockMode goodbasepb.GoodStockMode, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &powerrentalmwpb.Conds{
		State:     &v1.Uint32Val{Op: cruder.EQ, Value: uint32(state)},
		GoodType:  &v1.Uint32Val{Op: cruder.EQ, Value: uint32(goodType)},
		StockMode: &v1.Uint32Val{Op: cruder.EQ, Value: uint32(stockMode)},
	}

	for {
		handler, err := powerrentalmw.NewHandler(
			ctx,
			powerrentalmw.WithConds(conds),
			powerrentalmw.WithOffset(offset),
			powerrentalmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		goods, _, err := handler.GetPowerRentals(ctx)
		if err != nil {
			return err
		}
		if len(goods) == 0 {
			return nil
		}
		for _, good := range goods {
			cancelablefeed.CancelableFeed(ctx, good, exec)
		}
		offset += limit
	}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	return h.scanPowerRentals(ctx,
		goodbasepb.GoodState_GoodStateCheckHashRate,
		goodbasepb.GoodType_PowerRental,
		goodbasepb.GoodStockMode_GoodStockByMiningPool,
		exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if tx, ok := ent.(*types.PersistentGoodPowerRental); ok {
		return tx.EntID
	}
	return ent.(*powerrentalmwpb.PowerRental).EntID
}
