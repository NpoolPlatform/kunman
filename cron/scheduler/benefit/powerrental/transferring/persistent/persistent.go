package persistent

import (
	"context"
	"fmt"

	txmwcli "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	powerrentalmwcli "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/transferring/types"

	"github.com/shopspring/decimal"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, good interface{}, reward, notif, done chan interface{}) error {
	_good, ok := good.(*types.PersistentPowerRental)
	if !ok {
		return fmt.Errorf("invalid good")
	}

	defer asyncfeed.AsyncFeed(ctx, _good, done)

	if err := powerrentalmwcli.UpdatePowerRental(ctx, &powerrentalmwpb.PowerRentalReq{
		ID:          &_good.ID,
		RewardState: &_good.NewBenefitState,
	}); err != nil {
		return err
	}

	if len(_good.CoinRewards) == 0 {
		return nil
	}

	txReqs := []*txmwpb.TxReq{}
	for _, reward := range _good.CoinRewards {
		if !reward.Transferrable {
			continue
		}
		txReqs = append(txReqs, &txmwpb.TxReq{
			CoinTypeID:    &reward.CoinTypeID,
			FromAccountID: &reward.UserBenefitHotAccountID,
			ToAccountID:   &reward.PlatformColdAccountID,
			Amount:        &reward.ToPlatformAmount,
			FeeAmount:     func() *string { s := decimal.NewFromInt(0).String(); return &s }(),
			Extra:         &reward.Extra,
			Type:          func() *basetypes.TxType { e := basetypes.TxType_TxPlatformBenefit; return &e }(),
		})
	}
	if len(txReqs) == 0 {
		return nil
	}

	if _, err := txmwcli.CreateTxs(ctx, txReqs); err != nil {
		return err
	}

	return nil
}
