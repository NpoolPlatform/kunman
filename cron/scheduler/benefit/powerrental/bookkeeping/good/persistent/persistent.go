package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/good/types"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	goodstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/statement"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	goodstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/good/ledger/statement"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) updateGood(ctx context.Context, good *types.PersistentGood) error {
	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithID(&good.ID, true),
		powerrentalmw.WithRewardState(goodtypes.BenefitState_BenefitUserBookKeeping.Enum(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return handler.UpdatePowerRental(ctx)
}

func (p *handler) createGoodStatements(ctx context.Context, good *types.PersistentGood) error {
	stReqs := []*goodstatementmwpb.GoodStatementReq{}
	for _, reward := range good.CoinRewards {
		stReqs = append(stReqs, &goodstatementmwpb.GoodStatementReq{
			GoodID:                    &good.GoodID,
			CoinTypeID:                &reward.CoinTypeID,
			TotalAmount:               &reward.TotalRewardAmount,
			UnsoldAmount:              &reward.UnsoldRewardAmount,
			TechniqueServiceFeeAmount: &reward.TechniqueFeeAmount,
			BenefitDate:               &good.LastRewardAt,
		})
	}

	handler, err := goodstatementmw.NewHandler(
		ctx,
		goodstatementmw.WithReqs(stReqs, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	if _, err := handler.CreateGoodStatements(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (p *handler) Update(ctx context.Context, good interface{}, reward, notif, done chan interface{}) error {
	_good, ok := good.(*types.PersistentGood)
	if !ok {
		return fmt.Errorf("invalid good")
	}

	defer asyncfeed.AsyncFeed(ctx, _good, done)

	if len(_good.CoinRewards) > 0 {
		if err := p.createGoodStatements(ctx, _good); err != nil {
			return err
		}
	}
	return p.updateGood(ctx, _good)
}
