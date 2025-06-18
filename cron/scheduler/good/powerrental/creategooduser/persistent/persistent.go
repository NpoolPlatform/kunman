package persistent

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/creategooduser/types"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	goodusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdatePowerrentalState(ctx context.Context, good *types.PersistentGoodPowerRental) error {
	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithID(&good.ID, true),
		powerrentalmw.WithEntID(&good.EntID, true),
		powerrentalmw.WithGoodID(&good.GoodID, true),
		powerrentalmw.WithState(v1.GoodState_GoodStateCheckHashRate.Enum(), true),
		powerrentalmw.WithStocks(good.MiningGoodStockReqs, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdatePowerRental(ctx)
}

func (p *handler) withCreatePoolGoodUser(ctx context.Context, good *types.PersistentGoodPowerRental) error {
	for _, req := range good.GoodUserReqs {
		handler, err := goodusermw.NewHandler(
			ctx,
			goodusermw.WithEntID(req.EntID, false),
			goodusermw.WithRootUserID(req.RootUserID, true),
			goodusermw.WithCoinTypeIDs(req.CoinTypeIDs, true),
		)
		if err != nil {
			return err
		}

		if err := handler.CreateGoodUser(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (p *handler) Update(ctx context.Context, good interface{}, reward, notif, done chan interface{}) error {
	_good, ok := good.(*types.PersistentGoodPowerRental)
	if !ok {
		return fmt.Errorf("invalid feeorder")
	}

	defer asyncfeed.AsyncFeed(ctx, _good, done)

	if err := p.withCreatePoolGoodUser(ctx, _good); err != nil {
		return err
	}
	if err := p.withUpdatePowerrentalState(ctx, _good); err != nil {
		return err
	}

	return nil
}
