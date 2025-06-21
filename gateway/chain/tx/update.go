package tx

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/tx"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"
)

func (h *Handler) UpdateTx(ctx context.Context) (*npool.Tx, error) {
	conds := &txmwpb.Conds{
		// ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	}
	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithConds(conds),
	)
	if err != nil {
		return nil, err
	}

	tx, err := handler.GetTxOnly(ctx)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, wlog.Errorf("invalid tx")
	}
	if tx.State != basetypes.TxState_TxStateFail {
		return nil, wlog.Errorf("permission denied")
	}

	tx1, err := sphinxproxycli.GetTransaction(ctx, tx.EntID)
	if err != nil {
		return nil, err
	}
	if tx1 == nil {
		return nil, wlog.Errorf("invalid tx")
	}
	if tx1.TransactionState != sphinxproxypb.TransactionState_TransactionStateFail {
		return nil, wlog.Errorf("permission denied")
	}

	handler, err = txmw.NewHandler(
		ctx,
		txmw.WithID(h.ID, true),
		txmw.WithEntID(h.EntID, true),
		txmw.WithState(h.State, true),
	)
	if err != nil {
		return nil, err
	}

	if _, err := handler.UpdateTx(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := sphinxproxycli.UpdateTransaction(ctx, &sphinxproxypb.UpdateTransactionRequest{
		TransactionID:        tx.EntID,
		TransactionState:     tx1.TransactionState,
		NextTransactionState: sphinxproxypb.TransactionState_TransactionStateWait,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetTx(ctx)
}
