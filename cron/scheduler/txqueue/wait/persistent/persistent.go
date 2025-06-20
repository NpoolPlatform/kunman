package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/wait/types"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, tx interface{}, reward, notif, done chan interface{}) error {
	_tx, ok := tx.(*types.PersistentTx)
	if !ok {
		return fmt.Errorf("invalid tx")
	}

	defer asyncfeed.AsyncFeed(ctx, _tx, done)

	if _tx.NewTxState != basetypes.TxState_TxStateTransferring {
		handler, err := txmw.NewHandler(
			ctx,
			txmw.WithID(&_tx.ID, true),
			txmw.WithState(&_tx.NewTxState, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.UpdateTx(ctx); err != nil {
			return err
		}
		return nil
	}

	if !_tx.TransactionExist {
		if err := sphinxproxycli.CreateTransaction(ctx, &sphinxproxypb.CreateTransactionRequest{
			TransactionID: _tx.EntID,
			Name:          _tx.CoinName,
			Amount:        _tx.FloatAmount,
			From:          _tx.FromAddress,
			Memo:          _tx.AccountMemo,
			To:            _tx.ToAddress,
		}); err != nil {
			return err
		}
	}

	_tx.TransactionExist = true

	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithID(&_tx.ID, true),
		txmw.WithState(&_tx.NewTxState, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateTx(ctx); err != nil {
		return err
	}

	return nil
}
