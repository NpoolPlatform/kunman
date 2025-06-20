package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/simulate/types"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	simstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/simulate/ledger/statement"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	simulateledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/simulate/ledger/statement"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"

	"github.com/google/uuid"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderBenefitState(ctx context.Context, good *types.PersistentGood) error {
	reqs := []*powerrentalordermwpb.PowerRentalOrderReq{}
	state := ordertypes.BenefitState_BenefitBookKept
	for _, order := range good.OrderRewards {
		reqs = append(reqs, &powerrentalordermwpb.PowerRentalOrderReq{
			OrderID:      &order.OrderID,
			BenefitState: &state,
		})
	}

	multiHandler := &powerrentalordermw.MultiHandler{}
	for _, req := range reqs {
		handler, err := powerrentalordermw.NewHandler(
			ctx,
			powerrentalordermw.WithOrderID(req.OrderID, true),
			powerrentalordermw.WithBenefitState(req.BenefitState, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		multiHandler.AppendHandler(handler)
	}

	return multiHandler.UpdatePowerRentals(ctx)
}

func (p *handler) withCreateLedgerStatements(ctx context.Context, good *types.PersistentGood) error {
	simReqs := []*simstatementmwpb.StatementReq{}
	realReqs := []*statementmwpb.StatementReq{}

	ioType := ledgertypes.IOType_Incoming

	for _, reward := range good.OrderRewards {
		for _, coinReward := range reward.CoinRewards {
			simReqs = append(simReqs, &simstatementmwpb.StatementReq{
				EntID:      func() *string { s := uuid.NewString(); return &s }(),
				AppID:      &reward.AppID,
				UserID:     &reward.UserID,
				CoinTypeID: &coinReward.CoinTypeID,
				IOType:     &ioType,
				IOSubType:  func() *ledgertypes.IOSubType { e := ledgertypes.IOSubType_MiningBenefit; return &e }(),
				Amount:     &coinReward.Amount,
				IOExtra:    &reward.Extra,
				CreatedAt:  &good.LastRewardAt,
				SendCoupon: &coinReward.SendCoupon,
				Cashable:   &coinReward.Cashable,
			})
			if !coinReward.Cashable {
				continue
			}
			realReqs = append(realReqs, &statementmwpb.StatementReq{
				EntID:      func() *string { s := uuid.NewString(); return &s }(),
				AppID:      &reward.AppID,
				UserID:     &reward.UserID,
				CurrencyID: &coinReward.CoinTypeID,
				IOType:     &ioType,
				IOSubType:  func() *ledgertypes.IOSubType { e := ledgertypes.IOSubType_SimulateMiningBenefit; return &e }(),
				Amount:     &coinReward.Amount,
				IOExtra:    &reward.Extra,
				CreatedAt:  &good.LastRewardAt,
			})
		}
	}

	simulateHandler, err := simulateledgerstatementmw.NewHandler(
		ctx,
		simulateledgerstatementmw.WithReqs(simReqs, true),
	)
	if err != nil {
		return err
	}

	if _, err := simulateHandler.CreateStatements(ctx); err != nil {
		return err
	}

	ledgerHandler, err := ledgerstatementmw.NewHandler(
		ctx,
		ledgerstatementmw.WithReqs(realReqs, true),
	)
	if err != nil {
		return err
	}

	if _, err := ledgerHandler.CreateStatements(ctx); err != nil {
		return err
	}

	return nil
}

func (p *handler) updateGood(ctx context.Context, good *types.PersistentGood) error {
	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithID(&good.ID, true),
		powerrentalmw.WithRewardState(goodtypes.BenefitState_BenefitDone.Enum(), true),
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

	defer asyncfeed.AsyncFeed(ctx, _good, reward)

	if len(_good.OrderRewards) == 0 {
		if err := p.updateGood(ctx, _good); err != nil {
			return err
		}
		return nil
	}

	if err := p.withCreateLedgerStatements(ctx, _good); err != nil {
		return err
	}
	if err := p.withUpdateOrderBenefitState(ctx, _good); err != nil {
		return err
	}

	return nil
}
