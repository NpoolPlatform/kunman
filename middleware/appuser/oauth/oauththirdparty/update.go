package oauththirdparty

import (
	"context"
	"fmt"

	oauththirdpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/oauththirdparty"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent"
	entoauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/oauththirdparty"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/oauth/oauththirdparty"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

func (h *Handler) UpdateOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	info, err := h.GetOAuthThirdParty(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	if h.ClientName != nil {
		const limit = 2
		handler, err := NewHandler(
			ctx,
			WithConds(&npool.Conds{
				EntID:      &basetypes.StringVal{Op: cruder.NEQ, Value: info.EntID},
				ClientName: &basetypes.Int32Val{Op: cruder.EQ, Value: int32(*h.ClientName)},
			}),
			WithLimit(limit),
		)
		if err != nil {
			return nil, err
		}
		infos, _, err := handler.GetOAuthThirdParties(ctx)
		if err != nil {
			return nil, err
		}
		if infos != nil || len(infos) > 0 {
			return nil, fmt.Errorf("oauththirdparty exist")
		}
	}

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
				ClientName:     h.ClientName,
				ClientTag:      h.ClientTag,
				ClientLogoURL:  h.ClientLogoURL,
				ClientOAuthURL: h.ClientOAuthURL,
				ResponseType:   h.ResponseType,
				Scope:          h.Scope,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOAuthThirdParty(ctx)
}
