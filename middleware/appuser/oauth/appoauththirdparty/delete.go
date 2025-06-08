package appoauththirdparty

import (
	"context"
	"fmt"
	"time"

	appoauththirdpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/appoauththirdparty"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entappoauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appoauththirdparty"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/oauth/appoauththirdparty"
)

func (h *Handler) DeleteOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	info, err := h.GetOAuthThirdParty(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		user, err := tx.AppOAuthThirdParty.
			Query().
			Where(
				entappoauththirdparty.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if user == nil {
			return fmt.Errorf("invalid user")
		}

		if _, err := appoauththirdpartycrud.UpdateSet(
			user.Update(),
			&appoauththirdpartycrud.Req{
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
