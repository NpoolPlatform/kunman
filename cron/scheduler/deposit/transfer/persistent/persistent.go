package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/transfer/types"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	depositaccmw "github.com/NpoolPlatform/kunman/middleware/account/deposit"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"

	"github.com/google/uuid"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

// Here we could not use dtm to create transfer
func (p *handler) Update(ctx context.Context, account interface{}, reward, notif, done chan interface{}) error {
	_account, ok := account.(*types.PersistentAccount)
	if !ok {
		return fmt.Errorf("invalid account")
	}

	defer asyncfeed.AsyncFeed(ctx, _account, done)

	if _account.CollectingTIDCandidate == nil {
		collectingTID := uuid.NewString()
		_account.CollectingTIDCandidate = &collectingTID
	}

	locked := true
	lockedBy := basetypes.AccountLockedBy_Collecting

	depositHandler, err := depositaccmw.NewHandler(
		ctx,
		depositaccmw.WithID(&_account.ID, true),
		depositaccmw.WithAppID(&_account.AppID, true),
		depositaccmw.WithUserID(&_account.UserID, true),
		depositaccmw.WithCoinTypeID(&_account.CoinTypeID, true),
		depositaccmw.WithAccountID(&_account.DepositAccountID, true),
		depositaccmw.WithLocked(&locked, true),
		depositaccmw.WithLockedBy(&lockedBy, true),
		depositaccmw.WithCollectingTID(_account.CollectingTIDCandidate, true),
	)
	if err != nil {
		return err
	}

	if _, err := depositHandler.UpdateAccount(ctx); err != nil {
		return err
	}

	extra := fmt.Sprintf(
		`{"AppID":"%v","UserID":"%v","FromAddress":"%v","ToAddress":"%v"}`,
		_account.AppID,
		_account.UserID,
		_account.DepositAddress,
		_account.CollectAddress,
	)
	txType := basetypes.TxType_TxPaymentCollect

	txHandler, err := txmw.NewHandler(
		ctx,
		txmw.WithEntID(_account.CollectingTIDCandidate, true),
		txmw.WithCoinTypeID(&_account.CoinTypeID, true),
		txmw.WithFromAccountID(&_account.DepositAccountID, true),
		txmw.WithToAccountID(&_account.CollectAccountID, true),
		txmw.WithAmount(&_account.CollectAmount, true),
		txmw.WithFeeAmount(&_account.FeeAmount, true),
		txmw.WithExtra(&extra, true),
		txmw.WithType(&txType, true),
	)

	if _, err := txHandler.CreateTx(ctx); err != nil {
		return err
	}

	return nil
}
