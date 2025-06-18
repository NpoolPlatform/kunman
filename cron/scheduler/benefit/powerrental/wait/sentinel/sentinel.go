package sentinel

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/logger"
	powerrentalmwcli "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	common "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/wait/common"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/wait/types"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
)

type handler struct {
	*common.Handler
}

func NewSentinel() basesentinel.Scanner {
	h := &handler{
		Handler: common.NewHandler(),
	}
	return h
}

func (h *handler) scanGoods(ctx context.Context, state goodtypes.BenefitState, cond *types.TriggerCond, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		conds := &powerrentalmwpb.Conds{
			RewardState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(state)},
			StockMode:   &basetypes.Uint32Val{Op: cruder.NEQ, Value: uint32(goodtypes.GoodStockMode_GoodStockByMiningPool)},
		}
		if cond != nil {
			conds.GoodIDs = &basetypes.StringSliceVal{Op: cruder.IN, Value: cond.GoodIDs}
		}
		goods, _, err := powerrentalmwcli.GetPowerRentals(ctx, conds, offset, limit)
		if err != nil {
			return err
		}
		if len(goods) == 0 {
			return nil
		}

		for _, good := range goods {
			_good := &types.FeedPowerRental{
				PowerRental: good,
			}
			if cond != nil {
				_good.TriggerBenefitTimestamp = cond.RewardAt
			}
			cancelablefeed.CancelableFeed(ctx, _good, exec)
		}

		offset += limit
	}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	if time.Now().Before(h.NextBenefitAt()) {
		return nil
	}
	h.CalculateNextBenefitAt()
	return h.scanGoods(ctx, goodtypes.BenefitState_BenefitWait, nil, exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	_cond, ok := cond.(*types.TriggerCond)
	if !ok {
		return fmt.Errorf("invalid cond")
	}
	logger.Sugar().Infow(
		"TriggerScan",
		"GoodIDs", _cond.GoodIDs,
		"RewardAt", _cond.RewardAt,
	)
	return h.scanGoods(ctx, goodtypes.BenefitState_BenefitWait, _cond, exec)
}

func (h *handler) ObjectID(ent interface{}) string {
	if good, ok := ent.(*types.PersistentPowerRental); ok {
		return good.GoodID
	}
	if good, ok := ent.(*types.FeedPowerRental); ok {
		return good.GoodID
	}
	return ent.(*powerrentalmwpb.PowerRental).GoodID
}
