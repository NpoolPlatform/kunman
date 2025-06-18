package channel

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/channel"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/channel"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

func (h *Handler) DeleteChannel(ctx context.Context) (*npool.Channel, error) {
	info, err := h.GetChannel(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil { // dtm required
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := crud.UpdateSet(
			cli.NotifChannel.UpdateOneID(*h.ID),
			&crud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
