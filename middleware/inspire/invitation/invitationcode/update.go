package invitationcode

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/invitationcode"
	invitationcodecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) UpdateInvitationCode(ctx context.Context) (*npool.InvitationCode, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := invitationcodecrud.UpdateSet(
			cli.InvitationCode.UpdateOneID(*h.ID),
			&invitationcodecrud.Req{
				Disabled: h.Disabled,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetInvitationCode(ctx)
}
