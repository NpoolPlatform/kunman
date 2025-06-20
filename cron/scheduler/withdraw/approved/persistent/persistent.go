package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/approved/types"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	withdrawmw "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"

	"github.com/google/uuid"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, withdraw interface{}, reward, notif, done chan interface{}) error {
	_withdraw, ok := withdraw.(*types.PersistentWithdraw)
	if !ok {
		return fmt.Errorf("invalid withdraw")
	}

	defer asyncfeed.AsyncFeed(ctx, _withdraw, done)

	req := &withdrawmwpb.WithdrawReq{
		ID:    &_withdraw.ID,
		State: &_withdraw.NewWithdrawState,
	}
	if _withdraw.NewWithdrawState == ledgertypes.WithdrawState_Transferring {
		id := uuid.NewString()
		req.PlatformTransactionID = &id
	}

	handler, err := withdrawmw.NewHandler(
		ctx,
		withdrawmw.WithID(req.ID, true),
		withdrawmw.WithPlatformTransactionID(req.PlatformTransactionID, false),
		withdrawmw.WithChainTransactionID(req.ChainTransactionID, false),
		withdrawmw.WithState(req.State, false),
		withdrawmw.WithRollback(req.Rollback, false),
		withdrawmw.WithFeeAmount(req.FeeAmount, false),
		withdrawmw.WithReviewID(req.ReviewID, false),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateWithdraw(ctx); err != nil {
		return err
	}
	if _withdraw.NewWithdrawState == ledgertypes.WithdrawState_Transferring {
		txType := basetypes.TxType_TxWithdraw

		handler, err := txmw.NewHandler(
			ctx,
			txmw.WithEntID(req.PlatformTransactionID, true),
			txmw.WithCoinTypeID(&_withdraw.CoinTypeID, true),
			txmw.WithFromAccountID(&_withdraw.UserBenefitHotAccountID, true),
			txmw.WithToAccountID(&_withdraw.AccountID, true),
			txmw.WithAmount(&_withdraw.WithdrawAmount, true),
			txmw.WithFeeAmount(&_withdraw.WithdrawFeeAmount, true),
			txmw.WithType(&txType, true),
			txmw.WithExtra(&_withdraw.WithdrawExtra, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.CreateTx(ctx); err != nil {
			return err
		}
	}

	return nil
}
