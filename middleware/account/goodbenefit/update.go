package goodbenefit

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	goodbenefitcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/goodbenefit"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entgoodbenefit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/goodbenefit"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/goodbenefit"
)

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) { //nolint
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		goodbenefit, err := tx.GoodBenefit.
			Query().
			Where(
				entgoodbenefit.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if goodbenefit == nil {
			return fmt.Errorf("invalid goodbenefit")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(goodbenefit.AccountID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := accountcrud.UpdateSet(
			account.Update(),
			&accountcrud.Req{
				Active:   h.Active,
				Locked:   h.Locked,
				LockedBy: h.LockedBy,
				Blocked:  h.Blocked,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if _, err := goodbenefitcrud.UpdateSet(
			goodbenefit.Update(),
			&goodbenefitcrud.Req{
				Backup:        h.Backup,
				TransactionID: h.TransactionID,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if h.Backup != nil && *h.Backup {
			return nil
		}

		ids, err := tx.
			GoodBenefit.
			Query().
			Select().
			Modify(func(s *sql.Selector) {
				t := sql.Table(entaccount.Table)
				s.LeftJoin(t).
					On(
						t.C(entaccount.FieldEntID),
						s.C(entgoodbenefit.FieldAccountID),
					).
					OnP(
						sql.EQ(t.C(entaccount.FieldCoinTypeID), account.CoinTypeID),
					).
					OnP(
						sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
					)
				s.Where(
					sql.EQ(t.C(entaccount.FieldCoinTypeID), account.CoinTypeID),
				)
			}).
			Where(
				entgoodbenefit.GoodID(goodbenefit.GoodID),
				entgoodbenefit.IDNEQ(*h.ID),
				entgoodbenefit.Backup(false),
			).
			IDs(_ctx)
		if err != nil {
			return err
		}

		if _, err := tx.
			GoodBenefit.
			Update().
			Where(
				entgoodbenefit.IDIn(ids...),
			).
			SetBackup(true).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
