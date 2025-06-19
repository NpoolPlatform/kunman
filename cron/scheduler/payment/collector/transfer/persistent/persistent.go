package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/collector/transfer/types"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	payaccmw "github.com/NpoolPlatform/kunman/middleware/account/payment"
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

	if !_account.Locked {
		locked := true
		lockedBy := basetypes.AccountLockedBy_Collecting

		handler, err := payaccmw.NewHandler(
			ctx,
			payaccmw.WithID(&_account.ID, true),
			payaccmw.WithCoinTypeID(&_account.CoinTypeID, true),
			payaccmw.WithAccountID(&_account.PaymentAccountID, true),
			payaccmw.WithLocked(&locked, true),
			payaccmw.WithLockedBy(&lockedBy, true),
			payaccmw.WithCollectingTID(_account.CollectingTIDCandidate, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.UpdateAccount(ctx); err != nil {
			return err
		}
		_account.Locked = true
	}

	extra := fmt.Sprintf(
		`{"AccountID":"%v","Address":"%v","FromAddress":"%v","ToAddress":"%v"}`,
		_account.AccountID,
		_account.Address,
		_account.PaymentAddress,
		_account.CollectAddress,
	)
	txType := basetypes.TxType_TxPaymentCollect

	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithEntID(_account.CollectingTIDCandidate, true),
		txmw.WithCoinTypeID(&_account.CoinTypeID, true),
		txmw.WithFromAccountID(&_account.PaymentAccountID, true),
		txmw.WithToAccountID(&_account.CollectAccountID, true),
		txmw.WithAmount(&_account.CollectAmount, true),
		txmw.WithFeeAmount(&_account.FeeAmount, true),
		txmw.WithExtra(&extra, true),
		txmw.WithType(&txType, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.CreateTx(ctx); err != nil {
		return err
	}

	return nil
}
