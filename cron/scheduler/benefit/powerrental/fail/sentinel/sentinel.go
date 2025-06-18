package sentinel

import (
	"context"

	powerrentalmwcli "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/fail/types"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	h := &handler{}
	return h
}

func (h *handler) scanGoods(ctx context.Context, state goodtypes.BenefitState, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		goods, _, err := powerrentalmwcli.GetPowerRentals(ctx, &powerrentalmwpb.Conds{
			RewardState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(state)},
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
	return h.scanGoods(ctx, goodtypes.BenefitState_BenefitFail, exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return h.scanGoods(ctx, goodtypes.BenefitState_BenefitFail, exec)
}

func (h *handler) ObjectID(ent interface{}) string {
	if good, ok := ent.(*types.PersistentGood); ok {
		return good.GoodID
	}
	return ent.(*powerrentalmwpb.PowerRental).GoodID
}
