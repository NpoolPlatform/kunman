package payment

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"

	accountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/account"
	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	account1 "github.com/NpoolPlatform/kunman/middleware/account/account"
	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	paymentcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/payment"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateAccount(ctx context.Context) (*npool.Account, error) {
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

	usedFor := basetypes.AccountUsedFor_GoodPayment
	privateKey := true

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := accountcrud.CreateSet(
			tx.Account.Create(),
			&accountcrud.Req{
				EntID:                  h.AccountID,
				CoinTypeID:             h.CoinTypeID,
				Address:                h.Address,
				UsedFor:                &usedFor,
				PlatformHoldPrivateKey: &privateKey,
			},
		).Save(ctx); err != nil {
			return err
		}

		if _, err := paymentcrud.CreateSet(
			tx.Payment.Create(),
			&paymentcrud.Req{
				EntID:         h.EntID,
				AccountID:     h.AccountID,
				CollectingTID: h.CollectingTID,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
