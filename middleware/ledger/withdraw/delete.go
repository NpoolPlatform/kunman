package withdraw

import (
	"context"
	"fmt"
	"time"

	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	ledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	withdrawcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/withdraw"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entledger "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/ledger"
	entwithdraw "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/withdraw"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) unlockBalance(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Withdraw.
		Query().
		Where(
			entwithdraw.ID(*h.ID),
			entwithdraw.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	ledgerInfo, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(info.AppID),
			entledger.UserID(info.UserID),
			entledger.CurrencyID(info.CoinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	locked := decimal.NewFromInt(0).Sub(info.Amount)
	stm, err := ledgercrud.UpdateSetWithValidate(
		ledgerInfo,
		&ledgercrud.Req{
			AppID:        &info.AppID,
			UserID:       &info.UserID,
			CurrencyID:   &info.CoinTypeID,
			CurrencyType: types.CurrencyType_CurrencyCrypto.Enum(),
			Locked:       &locked,
			Spendable:    &info.Amount,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteWithdraw(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Withdraw.
		Query().
		Where(
			entwithdraw.ID(*h.ID),
			entwithdraw.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := withdrawcrud.UpdateSet(
		info.Update(),
		&withdrawcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteWithdraw(ctx context.Context) (*npool.Withdraw, error) {
	info, err := h.GetWithdraw(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	if info.State != types.WithdrawState_Created {
		return nil, fmt.Errorf("permission denied")
	}
	if h.ID == nil {
		h.ID = &info.ID
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.unlockBalance(ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteWithdraw(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}
