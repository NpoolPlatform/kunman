package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/transferring/types"
	logger "github.com/NpoolPlatform/kunman/framework/logger"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type txHandler struct {
	*txmwpb.Tx
	persistent chan interface{}
	notif      chan interface{}
	done       chan interface{}
	newState   basetypes.TxState
	txExtra    string
	txCID      *string
}

func (h *txHandler) checkTransfer(ctx context.Context) error {
	tx, err := sphinxproxycli.GetTransaction(ctx, h.EntID)
	if err != nil {
		switch status.Code(err) {
		case codes.InvalidArgument:
			fallthrough //nolint
		case codes.NotFound:
			fallthrough //nolint
		case codes.Aborted:
			h.newState = basetypes.TxState_TxStateFail
			return nil
		default:
			return err
		}
	} else if tx == nil {
		return fmt.Errorf("invalid transactionid")
	}

	switch tx.TransactionState {
	case sphinxproxypb.TransactionState_TransactionStateFail:
		h.newState = basetypes.TxState_TxStateFail
		h.txCID = &tx.CID
		if tx.CID == "" {
			txCID := "(fail without CID)"
			h.txCID = &txCID
		}
	case sphinxproxypb.TransactionState_TransactionStateDone:
		h.newState = basetypes.TxState_TxStateSuccessful
		h.txCID = &tx.CID
		if tx.CID == "" {
			txCID := "(successful without CID)"
			h.txCID = &txCID
			h.newState = basetypes.TxState_TxStateFail
		}
	}
	return nil
}

//nolint:gocritic
func (h *txHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Tx", h,
			"NewTxState", h.newState,
			"Error", *err,
		)
	}

	persistentTx := &types.PersistentTx{
		Tx:         h.Tx,
		NewTxState: h.newState,
		TxExtra:    h.txExtra,
		TxCID:      h.txCID,
	}
	if h.newState == h.State && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentTx, h.done)
		return
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentTx, h.notif)
	}
	if h.newState != h.State {
		asyncfeed.AsyncFeed(ctx, persistentTx, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentTx, h.done)
}

//nolint:gocritic
func (h *txHandler) exec(ctx context.Context) error {
	h.newState = h.State
	var err error

	defer h.final(ctx, &err)

	if err = h.checkTransfer(ctx); err != nil {
		return err
	}

	return nil
}
