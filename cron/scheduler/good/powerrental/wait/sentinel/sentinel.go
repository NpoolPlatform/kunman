package sentinel

import (
	"context"

	"github.com/NpoolPlatform/build-chain/pkg/constant"
	goodpowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	goodbasepb "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	goodpowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/wait/types"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) scanPowerRentals(ctx context.Context, state goodbasepb.GoodState, goodType goodbasepb.GoodType, stockMode goodbasepb.GoodStockMode, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		goods, _, err := goodpowerrentalmwcli.GetPowerRentals(ctx, &goodpowerrentalmwpb.Conds{
			State: &v1.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(state),
			},
			GoodType: &v1.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(goodType),
			},
			StockMode: &v1.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(stockMode),
			},
		}, offset, limit)
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
		goodbasepb.GoodState_GoodStateWait,
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
	return ent.(*goodpowerrentalmwpb.PowerRental).EntID
}
