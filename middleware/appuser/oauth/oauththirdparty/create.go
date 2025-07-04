package oauththirdparty

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/oauth/oauththirdparty"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	oauththirdpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/oauththirdparty"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	// TODO: deduplicate

	const limit = 2
	_handler, err := NewHandler(
		ctx,
		WithConds(&npool.Conds{
			ClientName: &basetypes.Int32Val{Op: cruder.EQ, Value: int32(*h.ClientName)},
		}),
		WithLimit(limit),
	)
	if err != nil {
		return nil, err
	}
	infos, _, err := _handler.GetOAuthThirdParties(ctx)
	if err != nil {
		return nil, err
	}
	if infos != nil || len(infos) > 0 {
		return infos[0], nil
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := oauththirdpartycrud.CreateSet(
			tx.OAuthThirdParty.Create(),
			&oauththirdpartycrud.Req{
				EntID:          h.EntID,
				ClientName:     h.ClientName,
				ClientTag:      h.ClientTag,
				ClientLogoURL:  h.ClientLogoURL,
				ClientOAuthURL: h.ClientOAuthURL,
				ResponseType:   h.ResponseType,
				Scope:          h.Scope,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOAuthThirdParty(ctx)
}
