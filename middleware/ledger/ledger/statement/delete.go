package statement

import (
	"context"
	"fmt"
	"time"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	ledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	profitcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/profit"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/statement"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entstatement "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/statement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) updateProfit(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	statement, err := tx.
		Statement.
		Query().
		Where(
			entstatement.ID(*req.ID),
			entstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	if statement.IoSubType != basetypes.IOSubType_MiningBenefit.String() &&
		statement.IoSubType != basetypes.IOSubType_SimulateMiningBenefit.String() {
		return nil
	}

	stm, err := profitcrud.SetQueryConds(
		tx.Profit.Query(),
		&profitcrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: statement.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: statement.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: statement.CurrencyID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	amount := decimal.NewFromInt(0).Sub(statement.Amount)
	stm1, err := profitcrud.UpdateSetWithValidate(
		info,
		&profitcrud.Req{
			Incoming: &amount,
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

func (h *deleteHandler) updateLedger(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	statement, err := tx.
		Statement.
		Query().
		Where(
			entstatement.ID(*req.ID),
			entstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}
	stm, err := ledgercrud.SetQueryConds(
		tx.Ledger.Query(),
		&ledgercrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: statement.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: statement.UserID},
			CurrencyID: &cruder.Cond{Op: cruder.EQ, Val: statement.CurrencyID},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	incoming := decimal.NewFromInt(0)
	outcoming := decimal.NewFromInt(0)
	ioType := basetypes.IOType(basetypes.IOType_value[statement.IoType])
	switch ioType {
	case basetypes.IOType_Incoming:
		incoming = decimal.NewFromInt(0).Sub(statement.Amount)
	case basetypes.IOType_Outcoming:
		outcoming = decimal.NewFromInt(0).Sub(statement.Amount)
	default:
		return fmt.Errorf("invalid io type %v", statement.IoType)
	}
	spendable := incoming.Sub(outcoming)

	stm1, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			Incoming:  &incoming,
			Outcoming: &outcoming,
			Spendable: &spendable,
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

func (h *deleteHandler) deleteStatement(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Statement.
		Query().
		Where(
			entstatement.ID(*req.ID),
			entstatement.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	now := uint32(time.Now().Unix())
	if _, err := crud.UpdateSet(
		info.Update(),
		&crud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteStatements(ctx context.Context) ([]*npool.Statement, error) { //nolint
	ids := []uint32{}
	entIDs := []uuid.UUID{}
	for _, req := range h.Reqs {
		if req.EntID == nil && req.ID == nil {
			return nil, fmt.Errorf("need id or entid")
		}
		if req.ID != nil {
			ids = append(ids, *req.ID)
			continue
		}
		if req.EntID != nil {
			entIDs = append(entIDs, *req.EntID)
		}
		// TODO: Deal Req with ID and EntID
	}
	infos := []*npool.Statement{}
	// if either EntIDs or IDs is empty, you cannot use EntIDs and IDs as conditional queries at the same time,
	// ent will add 'AND FALSE' at 'Where'
	if len(ids) > 0 {
		h.Conds = &crud.Conds{IDs: &cruder.Cond{Op: cruder.IN, Val: ids}}
		h.Limit = int32(len(ids))
		statements, _, err := h.GetStatements(ctx)
		if err != nil {
			return nil, err
		}
		infos = append(infos, statements...)
	}
	if len(entIDs) > 0 {
		h.Conds = &crud.Conds{EntIDs: &cruder.Cond{Op: cruder.IN, Val: entIDs}}
		h.Limit = int32(len(entIDs))
		statements, _, err := h.GetStatements(ctx)
		if err != nil {
			return nil, err
		}
		infos = append(infos, statements...)
	}
	if len(infos) != len(h.Reqs) {
		return nil, fmt.Errorf("statement not found")
	}

	statementMap := map[string]*npool.Statement{}
	for _, val := range infos {
		statementMap[val.EntID] = val
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.ID == nil {
				statement, ok := statementMap[req.EntID.String()]
				if !ok {
					return fmt.Errorf("statement not found")
				}
				req.ID = &statement.ID
			}
			_fn := func() error {
				if err := handler.updateProfit(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.updateLedger(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.deleteStatement(req, ctx, tx); err != nil {
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

func (h *Handler) DeleteStatement(ctx context.Context) (*npool.Statement, error) {
	info, err := h.GetStatement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
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

	h.Reqs = []*crud.Req{&h.Req}
	if _, err := h.DeleteStatements(ctx); err != nil {
		return nil, err
	}

	return info, nil
}
