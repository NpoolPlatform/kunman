package tx

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	txcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/tx"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateTx(ctx context.Context) (*npool.Tx, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := txcrud.CreateSet(
			cli.Tran.Create(),
			&txcrud.Req{
				EntID:         h.EntID,
				CoinTypeID:    h.CoinTypeID,
				FromAccountID: h.FromAccountID,
				ToAccountID:   h.ToAccountID,
				Amount:        h.Amount,
				FeeAmount:     h.FeeAmount,
				ChainTxID:     h.ChainTxID,
				State:         h.State,
				Extra:         h.Extra,
				Type:          h.Type,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTx(ctx)
}

func (h *Handler) CreateTxs(ctx context.Context) ([]*npool.Tx, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := txcrud.CreateSet(tx.Tran.Create(), req).Save(_ctx)
			if err != nil {
				return err
			}
			ids = append(ids, info.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &txcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetTxs(ctx)
	return infos, err
}
