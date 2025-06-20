package statement

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/statement"
	goodledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/good/ledger"
	goodstatementcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/good/ledger/statement"
	unsoldcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/good/ledger/unsold"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entgoodstatement "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/goodstatement"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/unsoldstatement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) updateGoodLedger(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	statement, err := tx.
		GoodStatement.
		Query().
		Where(
			entgoodstatement.ID(*req.ID),
			entgoodstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	stm, err := goodledgercrud.SetQueryConds(
		tx.GoodLedger.Query(),
		&goodledgercrud.Conds{
			GoodID:     &cruder.Cond{Op: cruder.EQ, Val: statement.GoodID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: statement.CoinTypeID},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	amount := decimal.NewFromInt(0).Sub(statement.Amount)
	toUser := decimal.NewFromInt(0).Sub(statement.ToUser)
	toPlatform := decimal.NewFromInt(0).Sub(statement.ToPlatform)

	stm1, err := goodledgercrud.UpdateSetWithValidate(
		info,
		&goodledgercrud.Req{
			Amount:     &amount,
			ToUser:     &toUser,
			ToPlatform: &toPlatform,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm1.Save(ctx); err != nil {
		return err
	}
	return nil
}

// nolint
func (h *deleteHandler) deleteGoodStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		GoodStatement.
		Query().
		Where(
			entgoodstatement.ID(*req.ID),
			entgoodstatement.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	now := uint32(time.Now().Unix())
	if _, err := goodstatementcrud.UpdateSet(
		info.Update(),
		&goodstatementcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

// nolint
func (h *deleteHandler) deleteUnsoldStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		UnsoldStatement.
		Query().
		Where(
			unsoldstatement.StatementID(*req.EntID),
			unsoldstatement.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	now := uint32(time.Now().Unix())
	if _, err := unsoldcrud.UpdateSet(
		info.Update(),
		&unsoldcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteGoodStatements(ctx context.Context) ([]*npool.GoodStatement, error) { //nolint
	handler := &deleteHandler{
		Handler: h,
	}
	ids := []uint32{}
	entIDs := []uuid.UUID{}
	for _, req := range h.Reqs {
		if req.ID == nil && req.EntID == nil {
			return nil, fmt.Errorf("need id or entid")
		}
		if req.ID != nil {
			ids = append(ids, *req.ID)
			continue
		}
		if req.EntID != nil {
			entIDs = append(entIDs, *req.EntID)
		}
	}
	infos := []*npool.GoodStatement{}
	// if either EntIDs or IDs is empty, you cannot use EntIDs and IDs as conditional queries at the same time,
	// ent will add 'AND FALSE' at 'Where'
	if len(ids) > 0 {
		h.Conds = &goodstatementcrud.Conds{IDs: &cruder.Cond{Op: cruder.IN, Val: ids}}
		h.Limit = int32(len(ids))
		statements, _, err := h.GetGoodStatements(ctx)
		if err != nil {
			return nil, err
		}
		infos = append(infos, statements...)
	}
	if len(entIDs) > 0 {
		h.Conds = &goodstatementcrud.Conds{EntIDs: &cruder.Cond{Op: cruder.IN, Val: entIDs}}
		h.Limit = int32(len(entIDs))
		statements, _, err := h.GetGoodStatements(ctx)
		if err != nil {
			return nil, err
		}
		infos = append(infos, statements...)
	}
	if len(infos) != len(h.Reqs) {
		if h.Rollback != nil && *h.Rollback {
			return nil, nil
		}
		return nil, fmt.Errorf("good statement not found")
	}

	goodStatementMap := map[string]*npool.GoodStatement{}
	idMap := map[uint32]*npool.GoodStatement{}
	for _, val := range infos {
		goodStatementMap[val.EntID] = val
		idMap[val.ID] = val
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.ID == nil {
				goodStatement, ok := goodStatementMap[req.EntID.String()]
				if !ok {
					return fmt.Errorf("good statement not found")
				}
				req.ID = &goodStatement.ID
			}
			if req.EntID == nil {
				goodStatement, ok := idMap[*req.ID]
				if !ok {
					return fmt.Errorf("good statement not found")
				}
				id, err := uuid.Parse(goodStatement.EntID)
				if err != nil {
					return err
				}
				req.EntID = &id
			}
			_fn := func() error {
				if err := handler.deleteUnsoldStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.updateGoodLedger(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.deleteGoodStatement(req, ctx, tx); err != nil {
					return err
				}
				return nil
			}
			if err := _fn(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (h *Handler) DeleteGoodStatement(ctx context.Context) (*npool.GoodStatement, error) {
	info, err := h.GetGoodStatement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		if h.Rollback != nil && *h.Rollback {
			return nil, nil
		}
		return nil, fmt.Errorf("statement not found")
	}
	if h.ID == nil {
		h.ID = &info.ID
	}
	if h.EntID == nil {
		id, err := uuid.Parse(info.EntID)
		if err != nil {
			return nil, err
		}
		h.EntID = &id
	}

	h.Reqs = []*goodstatementcrud.Req{&h.Req}
	if _, err := h.DeleteGoodStatements(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
