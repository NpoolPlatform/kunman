package ledger

import (
	"context"
	"errors"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	ledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type unlockHandler struct {
	*lockopHandler
	lop *ledgeropHandler
}

func (h *unlockHandler) unlockBalances(ctx context.Context) error {
	for _, lock := range h.locks {
		ledger := h.lop.ledger(lock.LedgerID)
		if ledger == nil {
			return fmt.Errorf("invalid ledger")
		}
		spendable := lock.Amount
		locked := decimal.NewFromInt(0).Sub(spendable)
		stm, err := ledgercrud.UpdateSetWithValidate(ledger, &ledgercrud.Req{
			Locked:    &locked,
			Spendable: &spendable,
		})
		if err != nil {
			return err
		}
		if _, err := stm.Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

//nolint:gocyclo
func (h *Handler) UnlockBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	handler := &unlockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.LedgerLockState_LedgerLockCanceled.Enum(),
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.getLocks(ctx, tx.LedgerLock); err != nil {
			if ent.IsNotFound(err) && h.Rollback != nil && *h.Rollback {
				return nil
			}
			return err
		}
		if len(handler.locks) == 0 {
			return nil
		}
		if h.Rollback != nil && *h.Rollback {
			handler.state = types.LedgerLockState_LedgerLockRollback.Enum()
		}
		handler.lop.ledgerIDs = []uuid.UUID{handler.locks[0].LedgerID}

		if err := handler.lop.getLedgers(ctx, tx); err != nil {
			return err
		}
		if err := handler.unlockBalances(ctx); err != nil {
			return err
		}
		if err := handler.updateLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if h.Rollback == nil || !*h.Rollback {
			return nil, err
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		if errors.Is(err, ledgercrud.ErrLedgerInconsistent) {
			return nil, nil
		}
		return nil, err
	}

	h.EntID = &handler.lop.ledgers[0].EntID
	return h.GetLedger(ctx)
}

//nolint:gocyclo
func (h *Handler) UnlockBalances(ctx context.Context) ([]*ledgermwpb.Ledger, error) {
	handler := &unlockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.LedgerLockState_LedgerLockCanceled.Enum(),
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.getLocks(ctx, tx.LedgerLock); err != nil {
			if ent.IsNotFound(err) && h.Rollback != nil && *h.Rollback {
				return nil
			}
			return err
		}
		if len(handler.locks) == 0 {
			return nil
		}
		if h.Rollback != nil && *h.Rollback {
			handler.state = types.LedgerLockState_LedgerLockRollback.Enum()
		}
		for _, lock := range handler.locks {
			handler.lop.ledgerIDs = append(handler.lop.ledgerIDs, lock.LedgerID)
		}

		if err := handler.lop.getLedgers(ctx, tx); err != nil {
			return err
		}
		if err := handler.unlockBalances(ctx); err != nil {
			return err
		}
		if err := handler.updateLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if h.Rollback == nil || !*h.Rollback {
			return nil, err
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		if errors.Is(err, ledgercrud.ErrLedgerInconsistent) {
			return nil, nil
		}
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
