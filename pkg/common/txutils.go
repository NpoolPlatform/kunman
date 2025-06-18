//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetTxs(ctx context.Context, txIDs []string) (map[string]*txmwpb.Tx, error) {
	for _, txID := range txIDs {
		if _, err := uuid.Parse(txID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &txmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: txIDs},
	}
	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithConds(conds),
		txmw.WithOffset(0),
		txmw.WithLimit(int32(len(txIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	txs, _, err := handler.GetTxs(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	txMap := map[string]*txmwpb.Tx{}
	for _, tx := range txs {
		txMap[tx.EntID] = tx
	}
	return txMap, nil
}
