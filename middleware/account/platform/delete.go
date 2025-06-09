package platform

import (
	"context"
	"time"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	platformcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/platform"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entplatform "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/platform"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
)

func (h *Handler) DeleteAccount(ctx context.Context) (*npool.Account, error) {
	info, err := h.GetAccount(ctx)
	if err != nil {
		return nil, err
	}

	now := uint32(time.Now().Unix())
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		platform, err := tx.Platform.
			Query().
			Where(
				entplatform.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(platform.AccountID),
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

		if _, err := platformcrud.UpdateSet(
			platform.Update(),
			&platformcrud.Req{
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
