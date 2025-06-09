package deposit

import (
	"context"
	"time"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	depositcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/deposit"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entdeposit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/deposit"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
)

func (h *Handler) DeleteAccount(ctx context.Context) (*npool.Account, error) {
	info, err := h.GetAccount(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		deposit, err := tx.Deposit.
			Query().
			Where(
				entdeposit.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(deposit.AccountID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := accountcrud.UpdateSet(
			account.Update(),
			&accountcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if _, err := depositcrud.UpdateSet(
			deposit.Update(),
			&depositcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
