package appsubscribe

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/subscriber/app/subscribe"
	appsubscribecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/subscriber/app/subscribe"
)

func (h *Handler) DeleteAppSubscribe(ctx context.Context) (*npool.AppSubscribe, error) {
	info, err := h.GetAppSubscribe(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := appsubscribecrud.UpdateSet(
			cli.AppSubscribe.UpdateOneID(*h.ID),
			&appsubscribecrud.Req{
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
