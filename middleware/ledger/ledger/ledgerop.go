package ledger

import (
	"context"
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entledger "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/ledger"

	"github.com/google/uuid"
)

type ledgeropHandler struct {
	*Handler
	ledgers   []*ent.Ledger
	ledgerIDs []uuid.UUID
}

func (h *ledgeropHandler) getLedgers(ctx context.Context, tx *ent.Tx) error {
	stm := tx.Ledger.Query()
	if len(h.ledgerIDs) > 0 {
		stm.Where(
			entledger.EntIDIn(h.ledgerIDs...),
			entledger.DeletedAt(0),
		)
	} else if len(h.Balances) > 0 {
		currencyIDs := []uuid.UUID{}
		for _, balance := range h.Balances {
			currencyIDs = append(currencyIDs, balance.CurrencyID)
		}
		stm.Where(
			entledger.AppID(*h.AppID),
			entledger.UserID(*h.UserID),
			entledger.CurrencyIDIn(currencyIDs...),
			entledger.DeletedAt(0),
		)
	} else {
		stm.Where(
			entledger.AppID(*h.AppID),
			entledger.UserID(*h.UserID),
			entledger.CurrencyID(*h.CurrencyID),
			entledger.DeletedAt(0),
		)
	}
	ledgers, err := stm.ForUpdate().All(ctx)
	if err != nil {
		return err
	}
	if len(ledgers) == 0 {
		return fmt.Errorf("invalid ledgers")
	}
	h.ledgers = ledgers
	return nil
}

func (h *ledgeropHandler) currencyLedger(currencyID uuid.UUID) *ent.Ledger {
	for _, ledger := range h.ledgers {
		if ledger.CurrencyID == currencyID {
			return ledger
		}
	}
	return nil
}

func (h *ledgeropHandler) updateLedger(ledger *ent.Ledger) {
	for i, _ledger := range h.ledgers {
		if _ledger.EntID == ledger.EntID {
			h.ledgers[i] = ledger
			break
		}
	}
}

func (h *ledgeropHandler) ledger(ledgerID uuid.UUID) *ent.Ledger {
	for _, ledger := range h.ledgers {
		if ledger.EntID == ledgerID {
			return ledger
		}
	}
	return nil
}
