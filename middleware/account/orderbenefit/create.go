package orderbenefit

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/NpoolPlatform/kunman/message/account/middleware/v1/orderbenefit"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"

	"github.com/google/uuid"
)

func (h *Handler) checkBaseAccount(ctx context.Context) (exist bool, err error) {
	if h.AccountID != nil {
		queryH, err := NewHandler(ctx,
			WithConds(&orderbenefit.Conds{
				AppID:     &v1.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
				UserID:    &v1.StringVal{Op: cruder.EQ, Value: h.UserID.String()},
				AccountID: &v1.StringVal{Op: cruder.EQ, Value: h.AccountID.String()},
			}),
			WithOffset(0),
			WithLimit(1),
		)
		if err != nil {
			return false, err
		}

		baseAccounts, _, err := queryH.GetAccounts(ctx)
		if err != nil {
			return false, err
		}
		if len(baseAccounts) < 1 {
			return false, fmt.Errorf("invalid accountid")
		}

		baseAccount := baseAccounts[0]
		if h.CoinTypeID != nil && baseAccount.CoinTypeID != h.CoinTypeID.String() {
			return false, fmt.Errorf("invalid cointypeid")
		}
		if h.CoinTypeID == nil {
			h.CoinTypeID = func() *uuid.UUID { id, _ := uuid.Parse(baseAccount.CoinTypeID); return &id }()
		}

		if h.accountReq.Address != nil && baseAccount.Address != *h.accountReq.Address {
			return false, fmt.Errorf("invalid address")
		}
		if h.accountReq.Address == nil {
			h.accountReq.Address = &baseAccount.Address
		}

		return true, nil
	} else if h.CoinTypeID == nil || h.accountReq.Address == nil {
		return false, fmt.Errorf("invalid cointypeid or address")
	} else {
		id := uuid.New()
		h.AccountID = &id
	}

	return false, nil
}

func (h *Handler) CreateAccountWithTx(ctx context.Context, tx *ent.Tx) error {
	if h.EntID == nil {
		id := uuid.New()
		h.EntID = &id
	}

	accountExist, err := h.checkBaseAccount(ctx)
	if err != nil {
		return err
	}

	sqlH := h.newSQLHandler()

	if !accountExist {
		if _, err := accountcrud.CreateSet(
			tx.Account.Create(),
			&accountcrud.Req{
				EntID:                  h.AccountID,
				CoinTypeID:             h.CoinTypeID,
				Address:                h.accountReq.Address,
				UsedFor:                h.accountReq.UsedFor,
				PlatformHoldPrivateKey: h.accountReq.PlatformHoldPrivateKey,
			},
		).Save(ctx); err != nil {
			return err
		}
	}

	sql, err := sqlH.genCreateSQL()
	if err != nil {
		return err
	}

	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return fmt.Errorf("fail create account: %v", err)
	}
	return nil
}

func (h *Handler) CreateAccount(ctx context.Context) error {
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return h.CreateAccountWithTx(ctx, tx)
	})
}
