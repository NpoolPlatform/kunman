package ledger

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ledgerlockcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/lock"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entledgerlock "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/ledgerlock"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LockBalance struct {
	LedgerID   uuid.UUID
	CoinTypeID uuid.UUID
	Amount     decimal.Decimal
}

type lockopHandler struct {
	*Handler
	locks []*ent.LedgerLock
	state *types.LedgerLockState
}

func (h *lockopHandler) getLocks(ctx context.Context, cli *ent.LedgerLockClient) error {
	locks, err := cli.
		Query().
		Where(
			entledgerlock.ExLockID(*h.LockID),
			entledgerlock.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return err
	}
	h.locks = locks
	return nil
}

func (h *lockopHandler) createLocks(ctx context.Context, tx *ent.Tx) error {
	if h.Locked != nil {
		if _, err := ledgerlockcrud.CreateSet(
			tx.LedgerLock.Create(),
			&ledgerlockcrud.Req{
				EntID:    h.LockID,
				LedgerID: h.EntID,
				Amount:   h.Locked,
				ExLockID: h.LockID,
			},
		).Save(ctx); err != nil {
			return err
		}
	}
	for _, balance := range h.Balances {
		if _, err := ledgerlockcrud.CreateSet(
			tx.LedgerLock.Create(),
			&ledgerlockcrud.Req{
				ExLockID: h.LockID,
				LedgerID: &balance.LedgerID,
				Amount:   &balance.Amount,
			},
		).Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *lockopHandler) updateLocks(ctx context.Context, tx *ent.Tx) error {
	for i, lock := range h.locks {
		switch lock.LockState {
		case types.LedgerLockState_LedgerLockLocked.String():
			switch *h.state {
			case types.LedgerLockState_LedgerLockSettle:
			case types.LedgerLockState_LedgerLockRollback:
			case types.LedgerLockState_LedgerLockCanceled:
			default:
				return fmt.Errorf("invalid ledgerlockstate")
			}
		default:
			return fmt.Errorf("invalid ledgerlockstate")
		}

		stm := tx.
			LedgerLock.
			UpdateOneID(lock.ID).
			SetLockState(h.state.String())
		if *h.state == types.LedgerLockState_LedgerLockSettle {
			stm.SetStatementID(h.StatementIDs[i])
		}
		if _, err := stm.Save(ctx); err != nil {
			return err
		}
	}
	return nil
}
