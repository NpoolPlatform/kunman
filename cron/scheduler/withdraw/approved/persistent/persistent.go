package persistent

import (
	"context"
	"fmt"

	txmwcli "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	withdrawmwcli "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/approved/types"

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
	if _, err := withdrawmwcli.UpdateWithdraw(ctx, req); err != nil {
		return err
	}
	if _withdraw.NewWithdrawState == ledgertypes.WithdrawState_Transferring {
		txType := basetypes.TxType_TxWithdraw
		if _, err := txmwcli.CreateTx(ctx, &txmwpb.TxReq{
			EntID:         req.PlatformTransactionID,
			CoinTypeID:    &_withdraw.CoinTypeID,
			FromAccountID: &_withdraw.UserBenefitHotAccountID,
			ToAccountID:   &_withdraw.AccountID,
			Amount:        &_withdraw.WithdrawAmount,
			FeeAmount:     &_withdraw.WithdrawFeeAmount,
			Type:          &txType,
			Extra:         &_withdraw.WithdrawExtra,
		}); err != nil {
			return err
		}
	}

	return nil
}
