package user

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/user"
	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	usercrud "github.com/NpoolPlatform/kunman/middleware/account/crud/user"

	"github.com/google/uuid"
)

func (h *Handler) CreateAccount(ctx context.Context) (*npool.Account, error) {
	// TODO: deduplicate

	id1 := uuid.New()
	if h.EntID == nil {
		h.EntID = &id1
	}

	id2 := uuid.New()
	if h.AccountID == nil {
		h.AccountID = &id2
	}

	privateKey := true

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := accountcrud.CreateSet(
			tx.Account.Create(),
			&accountcrud.Req{
				EntID:                  h.AccountID,
				CoinTypeID:             h.CoinTypeID,
				Address:                h.Address,
				UsedFor:                h.UsedFor,
				PlatformHoldPrivateKey: &privateKey,
			},
		).Save(ctx); err != nil {
			return err
		}

		if _, err := usercrud.CreateSet(
			tx.User.Create(),
			&usercrud.Req{
				EntID:      h.EntID,
				AppID:      h.AppID,
				UserID:     h.UserID,
				CoinTypeID: h.CoinTypeID,
				AccountID:  h.AccountID,
				UsedFor:    h.UsedFor,
				Labels:     h.Labels,
				Memo:       h.Memo,
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
