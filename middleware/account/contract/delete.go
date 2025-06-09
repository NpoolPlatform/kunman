package contract

import (
	"context"
	"time"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	contractcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/contract"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entcontract "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/contract"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/contract"
)

func (h *Handler) DeleteAccount(ctx context.Context) (*npool.Account, error) {
	info, err := h.GetAccount(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	if h.ID == nil {
		h.ID = &info.ID
	}

	now := uint32(time.Now().Unix())
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		contract, err := tx.Contract.
			Query().
			Where(
				entcontract.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(contract.AccountID),
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

		if _, err := contractcrud.UpdateSet(
			contract.Update(),
			&contractcrud.Req{
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
