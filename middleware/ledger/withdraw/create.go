package withdraw

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	ledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/withdraw"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entledger "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/ledger"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) lockBalance(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(*h.AppID),
			entledger.UserID(*h.UserID),
			entledger.CoinTypeID(*h.CoinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	spendable := decimal.NewFromInt(0).Sub(*h.Amount)
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     h.Amount,
			Spendable:  &spendable,
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

func (h *createHandler) createWithdraw(ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.CreateSet(
		tx.Withdraw.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateWithdraw(ctx context.Context) (*npool.Withdraw, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.lockBalance(ctx, tx); err != nil {
			return err
		}
		if err := handler.createWithdraw(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetWithdraw(ctx)
}
