package ledger

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	ledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*lockopHandler
	lop *ledgeropHandler
}

func (h *lockHandler) lockBalance(ctx context.Context) error {
	spendable := decimal.NewFromInt(0).Sub(*h.Locked)
	stm, err := ledgercrud.UpdateSetWithValidate(h.lop.ledgers[0], &ledgercrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CurrencyID: h.CurrencyID,
		Locked:     h.Locked,
		Spendable:  &spendable,
	})
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) LockBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	handler := &lockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.getLocks(ctx, tx.LedgerLock); err != nil {
			return err
		}
		if len(handler.locks) > 0 {
			return fmt.Errorf("invalid lockid")
		}

		if err := handler.lop.getLedgers(ctx, tx); err != nil {
			return err
		}
		h.EntID = &handler.lop.ledgers[0].EntID
		if err := handler.lockBalance(ctx); err != nil {
			return err
		}
		if err := handler.createLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetLedger(ctx)
}

func (h *lockHandler) lockBalances(ctx context.Context) error {
	for _, balance := range h.Balances {
		ledger := h.lop.currencyLedger(balance.CurrencyID)
		if ledger == nil {
			return fmt.Errorf("invalid ledger")
		}
		spendable := decimal.NewFromInt(0).Sub(balance.Amount)
		stm, err := ledgercrud.UpdateSetWithValidate(ledger, &ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CurrencyID: &balance.CurrencyID,
			Locked:     &balance.Amount,
			Spendable:  &spendable,
		})
		if err != nil {
			return err
		}
		if ledger, err = stm.Save(ctx); err != nil {
			return err
		}
		h.lop.updateLedger(ledger)
	}
	return nil
}

func (h *Handler) LockBalancesWithTx(ctx context.Context, tx *ent.Tx) ([]*ledgermwpb.Ledger, error) {
	handler := &lockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	if err := handler.getLocks(ctx, tx.LedgerLock); err != nil {
		return nil, err
	}
	if len(handler.locks) > 0 {
		return nil, fmt.Errorf("invalid lockid")
	}

	if err := handler.lop.getLedgers(ctx, tx); err != nil {
		return nil, err
	}
	for _, balance := range h.Balances {
		ledger := handler.lop.currencyLedger(balance.CurrencyID)
		if ledger == nil {
			return nil, fmt.Errorf("invalid ledger")
		}
		balance.LedgerID = ledger.EntID
	}
	if err := handler.lockBalances(ctx); err != nil {
		return nil, err
	}
	if err := handler.createLocks(ctx, tx); err != nil {
		return nil, err
	}

	h.Conds = &ledgercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: handler.lop.ledgerIDs},
	}
	h.Offset = 0
	h.Limit = int32(len(handler.lop.ledgerIDs))
	infos, _, err := h.GetLedgersWithTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func (h *Handler) LockBalances(ctx context.Context) (infos []*ledgermwpb.Ledger, err error) {
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		infos, err = h.LockBalancesWithTx(_ctx, tx)
		return err
	})
	return infos, wlog.WrapError(err)
}
