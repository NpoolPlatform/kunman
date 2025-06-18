package readstate

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/readstate"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement/readstate"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

func (h *Handler) DeleteReadState(ctx context.Context) (*npool.ReadState, error) {
	info, err := h.GetReadState(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil { // dtm required
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := crud.UpdateSet(
			cli.ReadAnnouncement.UpdateOneID(*h.ID),
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
