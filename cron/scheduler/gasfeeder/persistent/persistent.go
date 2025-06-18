package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/gasfeeder/types"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, coin interface{}, reward, notif, done chan interface{}) error {
	_coin, ok := coin.(*types.PersistentCoin)
	if !ok {
		return fmt.Errorf("invalid coin")
	}

	defer asyncfeed.AsyncFeed(ctx, _coin, done)

	txType := basetypes.TxType_TxFeedGas

	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithCoinTypeID(&_coin.FeeCoinTypeID, true),
		txmw.WithFromAccountID(&_coin.FromAccountID, true),
		txmw.WithToAccountID(&_coin.ToAccountID, true),
		txmw.WithAmount(&_coin.Amount, true),
		txmw.WithFeeAmount(&_coin.FeeAmount, true),
		txmw.WithExtra(&_coin.Extra, true),
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
