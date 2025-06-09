package goodbenefit

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entgoodbenefit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/goodbenefit"

	accountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/account"
	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/goodbenefit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	account1 "github.com/NpoolPlatform/kunman/middleware/account/account"
	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	goodbenefitcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/goodbenefit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateAccount(ctx context.Context) (*npool.Account, error) { //nolint
	// TODO: deduplicate

	handler, err := account1.NewHandler(
		ctx,
		account1.WithConds(&accountmwpb.Conds{
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.CoinTypeID.String()},
			Address:    &basetypes.StringVal{Op: cruder.EQ, Value: *h.Address},
		}),
	)
	if err != nil {
		return nil, err
	}
	exist, err := handler.ExistAccountConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("address exist")
	}

	id1 := uuid.New()
	if h.EntID == nil {
		h.EntID = &id1
	}

	id2 := uuid.New()
	if h.AccountID == nil {
		h.AccountID = &id2
	}

	usedFor := basetypes.AccountUsedFor_GoodBenefit
	privateKey := true

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := accountcrud.CreateSet(
			tx.Account.Create(),
			&accountcrud.Req{
				EntID:                  h.AccountID,
				CoinTypeID:             h.CoinTypeID,
				Address:                h.Address,
				UsedFor:                &usedFor,
				PlatformHoldPrivateKey: &privateKey,
			},
		).Save(_ctx); err != nil {
			return err
		}

		goodbenefit, err := goodbenefitcrud.CreateSet(
			tx.GoodBenefit.Create(),
			&goodbenefitcrud.Req{
				EntID:     h.EntID,
				GoodID:    h.GoodID,
				AccountID: h.AccountID,
				Backup:    h.Backup,
			},
		).Save(_ctx)
		if err != nil {
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
						sql.EQ(t.C(entaccount.FieldCoinTypeID), *h.CoinTypeID),
					).
					OnP(
						sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
					)
				s.Where(
					sql.EQ(t.C(entaccount.FieldCoinTypeID), *h.CoinTypeID),
				)
			}).
			Where(
				entgoodbenefit.GoodID(goodbenefit.GoodID),
				entgoodbenefit.EntIDNEQ(*h.EntID),
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
