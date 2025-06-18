package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/done/types"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	goodcoinrewardmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) updateOrders(ctx context.Context, good *types.PersistentGood) error {
	reqs := []*powerrentalordermwpb.PowerRentalOrderReq{}
	state := ordertypes.BenefitState_BenefitWait
	for _, id := range good.BenefitOrderIDs {
		_id := id
		reqs = append(reqs, &powerrentalordermwpb.PowerRentalOrderReq{
			ID:           &_id,
			BenefitState: &state,
		})
	}

	multiHandler := &powerrentalordermw.MultiHandler{}
	for _, req := range reqs {
		handler, err := powerrentalordermw.NewHandler(
			ctx,
			powerrentalordermw.WithID(req.ID, true),
			powerrentalordermw.WithBenefitState(req.BenefitState, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		multiHandler.AppendHandler(handler)
	}

	return multiHandler.UpdatePowerRentals(ctx)
}

func (p *handler) updateGood(ctx context.Context, good *types.PersistentGood) error {
	reqs := []*goodcoinrewardmwpb.RewardReq{}
	for _, reward := range good.CoinNextRewards {
		reqs = append(reqs, &goodcoinrewardmwpb.RewardReq{
			CoinTypeID:            &reward.CoinTypeID,
			NextRewardStartAmount: &reward.NextRewardStartAmount,
		})
	}
	state := goodtypes.BenefitState_BenefitWait

	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithID(&good.ID, true),
		powerrentalmw.WithRewardState(&state, true),
		powerrentalmw.WithRewards(reqs, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdatePowerRental(ctx)
}

func (p *handler) Update(ctx context.Context, good interface{}, reward, notif, done chan interface{}) error {
	_good, ok := good.(*types.PersistentGood)
	if !ok {
		return fmt.Errorf("invalid good")
	}

	defer asyncfeed.AsyncFeed(ctx, _good, done)
	asyncfeed.AsyncFeed(ctx, _good, notif)

	if err := p.updateOrders(ctx, _good); err != nil {
		return err
	}
	return p.updateGood(ctx, _good)
}
