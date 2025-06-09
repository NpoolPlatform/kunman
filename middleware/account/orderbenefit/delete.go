package orderbenefit

import (
	"context"
	"time"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	orderbenefitcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/orderbenefit"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entorderbenefit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/orderbenefit"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"

	"github.com/google/uuid"
)

//nolint:gocyclo
func (h *Handler) DeleteAccountWithTx(ctx context.Context, tx *ent.Tx) error {
	info, err := h.GetAccountWithTx(ctx, tx)
	if err != nil {
		return err
	}

	if info == nil {
		return nil
	}

	if h.AppID != nil && h.AppID.String() != info.AppID {
		return wlog.Errorf("invalid appid")
	}
	if h.UserID != nil && h.UserID.String() != info.UserID {
		return wlog.Errorf("invalid userid")
	}
	if h.OrderID != nil && h.OrderID.String() != info.OrderID {
		return wlog.Errorf("invalid orderid")
	}
	if h.AccountID != nil && h.AccountID.String() != info.AccountID {
		return wlog.Errorf("invalid accountid")
	}
	accountID, err := uuid.Parse(info.AccountID)
	if err != nil {
		return err
	}

	oderbenefitID, err := uuid.Parse(info.EntID)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	now := uint32(time.Now().Unix())
	account, err := tx.Account.
		Query().
		Where(
			entaccount.EntID(accountID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	orderbenefitAcc, err := tx.OrderBenefit.
		Query().
		Where(
			entorderbenefit.EntID(oderbenefitID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if h.AccountID == nil {
		_, err = accountcrud.UpdateSet(
			account.Update(),
			&accountcrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
	}

	_, err = orderbenefitcrud.UpdateSet(
		orderbenefitAcc.Update(),
		&orderbenefitcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteAccount(ctx context.Context) error {
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return h.DeleteAccountWithTx(ctx, tx)
	})
}
