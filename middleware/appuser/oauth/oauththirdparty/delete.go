package oauththirdparty

import (
	"context"
	"fmt"
	"time"

	oauththirdpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/oauththirdparty"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entoauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/oauththirdparty"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/oauth/oauththirdparty"
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
		user, err := tx.OAuthThirdParty.
			Query().
			Where(
				entoauththirdparty.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if user == nil {
			return fmt.Errorf("invalid user")
		}

		if _, err := oauththirdpartycrud.UpdateSet(
			user.Update(),
			&oauththirdpartycrud.Req{
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
