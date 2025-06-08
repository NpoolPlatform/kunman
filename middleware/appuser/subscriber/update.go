package subscriber

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/subscriber"
	subscribercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/subscriber"
)

func (h *Handler) UpdateSubscriber(ctx context.Context) (*npool.Subscriber, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := subscribercrud.UpdateSet(
			cli.Subscriber.UpdateOneID(*h.ID),
			&subscribercrud.Req{
				Registered: h.Registered,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSubscriber(ctx)
}
