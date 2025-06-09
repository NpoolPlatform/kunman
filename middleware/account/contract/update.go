package contract

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	contractcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/contract"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entcontract "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/contract"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/contract"
)

type updateHandler struct {
	*Handler
	contract *ent.Contract
}

func (h *updateHandler) getContract(ctx context.Context, tx *ent.Tx, must bool) (err error) {
	stm := tx.Contract.Query()
	if h.ID != nil {
		stm.Where(entcontract.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcontract.EntID(*h.EntID))
	}

	if h.contract, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return err
	}
	return nil
}

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) {
	if h.ID == nil && h.EntID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	handler := &updateHandler{
		Handler: h,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		err := handler.getContract(ctx, tx, true)
		if err != nil {
			return err
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(handler.contract.AccountID),
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

		if _, err := contractcrud.UpdateSet(
			handler.contract.Update(),
			&contractcrud.Req{
				Backup: h.Backup,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if h.Backup != nil && *h.Backup {
			return nil
		}

		ids, err := tx.
			Contract.
			Query().
			Select().
			Modify(func(s *sql.Selector) {
				t := sql.Table(entaccount.Table)
				s.LeftJoin(t).
					On(
						t.C(entaccount.FieldEntID),
						s.C(entcontract.FieldAccountID),
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
				entcontract.GoodID(handler.contract.GoodID),
				entcontract.IDNEQ(handler.contract.ID),
				entcontract.Backup(false),
			).
			IDs(_ctx)
		if err != nil {
			return err
		}

		if _, err := tx.
			Contract.
			Update().
			Where(
				entcontract.IDIn(ids...),
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
