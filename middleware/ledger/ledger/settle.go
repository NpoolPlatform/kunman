package ledger

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	ledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	statementcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/statement"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type settleHandler struct {
	*lockopHandler
	lop *ledgeropHandler
}

func (h *settleHandler) settleBalances(ctx context.Context) error {
	for _, lock := range h.locks {
		outcoming := lock.Amount
		locked := decimal.NewFromInt(0).Sub(outcoming)

		ledger := h.lop.ledger(lock.LedgerID)
		if ledger == nil {
			return wlog.Errorf("invalid ledger")
		}

		stm, err := ledgercrud.UpdateSetWithValidate(ledger, &ledgercrud.Req{
			Locked:    &locked,
			Outcoming: &outcoming,
		})
		if err != nil {
			return wlog.WrapError(err)
		}
		if ledger, err = stm.Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
		h.lop.updateLedger(ledger)
	}
	return nil
}

func (h *settleHandler) createStatements(ctx context.Context, tx *ent.Tx) error {
	// TODO: work around of settle balances

	statementCreatables := map[uuid.UUID]struct{}{}

	for i, lock := range h.locks {
		if err := func() error {
			ledger := h.lop.ledger(lock.LedgerID)
			if ledger == nil {
				return wlog.Errorf("invalid ledger")
			}

			ioType := types.IOType_Outcoming
			stm, err := statementcrud.SetQueryConds(tx.Statement.Query(), &statementcrud.Conds{
				AppID:      &cruder.Cond{Op: cruder.EQ, Val: ledger.AppID},
				UserID:     &cruder.Cond{Op: cruder.EQ, Val: ledger.UserID},
				CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: ledger.CoinTypeID},
				IOType:     &cruder.Cond{Op: cruder.EQ, Val: ioType},
				IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: *h.IOSubType},
				IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: *h.IOExtra},
			})
			if err != nil {
				return wlog.WrapError(err)
			}
			exist, err := stm.Exist(ctx)
			if err != nil {
				return wlog.WrapError(err)
			}
			// Workaround: if we have same coins in this batch settle, we just check the first one
			if _, ok := statementCreatables[ledger.CoinTypeID]; !ok && exist {
				return wlog.Errorf("statement already exist")
			}

			statementCreatables[ledger.CoinTypeID] = struct{}{}

			if _, err := statementcrud.CreateSet(tx.Statement.Create(), &statementcrud.Req{
				EntID:      &h.StatementIDs[i],
				AppID:      &ledger.AppID,
				UserID:     &ledger.UserID,
				CoinTypeID: &ledger.CoinTypeID,
				IOType:     &ioType,
				IOSubType:  h.IOSubType,
				IOExtra:    h.IOExtra,
				Amount:     &lock.Amount,
			}).Save(ctx); err != nil {
				return wlog.WrapError(err)
			}
			return nil
		}(); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *Handler) SettleBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	handler := &settleHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.LedgerLockState_LedgerLockSettle.Enum(),
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	if err := handler.getLocks(ctx); err != nil {
		return nil, err
	}

	handler.lop.ledgerIDs = []uuid.UUID{handler.locks[0].LedgerID}
	h.StatementIDs = []uuid.UUID{*h.StatementID}
	if len(h.StatementIDs) != len(handler.locks) {
		return nil, wlog.Errorf("mismatched statementids")
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error { //nolint:dupl
		if err := handler.lop.getLedgers(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.settleBalances(ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createStatements(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateLocks(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.ID = &handler.lop.ledgers[0].ID
	return h.GetLedger(ctx)
}

func (h *Handler) SettleBalances(ctx context.Context) ([]*ledgermwpb.Ledger, error) {
	handler := &settleHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.LedgerLockState_LedgerLockSettle.Enum(),
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	if err := handler.getLocks(ctx); err != nil {
		return nil, err
	}
	for _, lock := range handler.locks {
		handler.lop.ledgerIDs = append(handler.lop.ledgerIDs, lock.LedgerID)
	}
	if len(h.StatementIDs) != len(handler.locks) {
		return nil, wlog.Errorf("mismatched statementids")
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error { //nolint:dupl
		if err := handler.lop.getLedgers(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.settleBalances(ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createStatements(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateLocks(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &ledgercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: handler.lop.ledgerIDs},
	}
	h.Offset = 0
	h.Limit = int32(len(handler.lop.ledgerIDs))
	infos, _, err := h.GetLedgers(ctx)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
