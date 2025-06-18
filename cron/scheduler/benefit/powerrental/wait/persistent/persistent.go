package persistent

import (
	"context"

	txmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/tx"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	powerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	goodcoinrewardmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/wait/types"
	powerrentalordermwcli "github.com/NpoolPlatform/order-middleware/pkg/client/powerrental"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) updateOrders(ctx context.Context, good *types.PersistentPowerRental) error {
	reqs := []*powerrentalordermwpb.PowerRentalOrderReq{}
	state := ordertypes.BenefitState_BenefitCalculated
	for _, id := range good.BenefitOrderIDs {
		_id := id
		reqs = append(reqs, &powerrentalordermwpb.PowerRentalOrderReq{
			ID:            &_id,
			LastBenefitAt: &good.BenefitTimestamp,
			BenefitState:  &state,
		})
	}
	return powerrentalordermwcli.UpdatePowerRentalOrders(ctx, reqs)
}

func (p *handler) Update(ctx context.Context, good interface{}, reward, notif, done chan interface{}) error {
	_good, ok := good.(*types.PersistentPowerRental)
	if !ok {
		return wlog.Errorf("invalid good")
	}

	defer asyncfeed.AsyncFeed(ctx, _good, done)

	if len(_good.CoinRewards) > 0 {
		if err := p.updateOrders(ctx, _good); err != nil {
			return wlog.WrapError(err)
		}
	}

	rewardReqs := []*goodcoinrewardmwpb.RewardReq{}
	for _, reward := range _good.CoinRewards {
		rewardReqs = append(rewardReqs, &goodcoinrewardmwpb.RewardReq{
			GoodID:     &_good.GoodID,
			CoinTypeID: &reward.CoinTypeID,
			RewardTID: func() *string {
				s := uuid.NewString()
				if !reward.Transferrable {
					s = uuid.Nil.String()
				}
				return &s
			}(),
			RewardAmount:          &reward.Amount,
			NextRewardStartAmount: &reward.NextRewardStartAmount,
		})
	}

	if err := powerrentalmwcli.UpdatePowerRental(ctx, &powerrentalmwpb.PowerRentalReq{
		ID:          &_good.ID,
		RewardState: func() *goodtypes.BenefitState { e := goodtypes.BenefitState_BenefitTransferring; return &e }(),
		RewardAt:    &_good.BenefitTimestamp,
		Rewards:     rewardReqs,
	}); err != nil {
		return wlog.WrapError(err)
	}

	if len(_good.CoinRewards) == 0 {
		return nil
	}

	txReqs := []*txmwpb.TxReq{}
	for i, reward := range _good.CoinRewards {
		if !reward.Transferrable {
			continue
		}
		txReqs = append(txReqs, &txmwpb.TxReq{
			EntID:         rewardReqs[i].RewardTID,
			CoinTypeID:    &reward.CoinTypeID,
			FromAccountID: &reward.GoodBenefitAccountID,
			ToAccountID:   &reward.UserBenefitHotAccountID,
			Amount:        &reward.Amount,
			FeeAmount:     func() *string { s := decimal.NewFromInt(0).String(); return &s }(),
			Extra:         &reward.Extra,
			Type:          func() *basetypes.TxType { e := basetypes.TxType_TxUserBenefit; return &e }(),
		})
	}

	if _, err := txmwcli.CreateTxs(ctx, txReqs); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}
