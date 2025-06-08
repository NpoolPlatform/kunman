package subscriber

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent"

	subscribercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/subscriber"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/subscriber"
)

func (h *Handler) DeleteSubscriber(ctx context.Context) (*npool.Subscriber, error) {
	info, err := h.GetSubscriber(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := subscribercrud.UpdateSet(
			cli.Subscriber.UpdateOneID(*h.ID),
			&subscribercrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
