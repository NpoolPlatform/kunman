package persistent

import (
	"context"
	"fmt"

	txmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/tx"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/gasfeeder/types"
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
	if _, err := txmwcli.CreateTx(ctx, &txmwpb.TxReq{
		CoinTypeID:    &_coin.FeeCoinTypeID,
		FromAccountID: &_coin.FromAccountID,
		ToAccountID:   &_coin.ToAccountID,
		Amount:        &_coin.Amount,
		FeeAmount:     &_coin.FeeAmount,
		Extra:         &_coin.Extra,
		Type:          &txType,
	}); err != nil {
		return err
	}

	return nil
}
