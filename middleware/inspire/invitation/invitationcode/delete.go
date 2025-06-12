package invitationcode

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/invitationcode"
	invitationcodecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) DeleteInvitationCode(ctx context.Context) (*npool.InvitationCode, error) {
	info, err := h.GetInvitationCode(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, nil
	}
	if h.ID == nil {
		h.ID = &info.ID
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := invitationcodecrud.UpdateSet(
			cli.InvitationCode.UpdateOneID(*h.ID),
			&invitationcodecrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return info, nil
}
