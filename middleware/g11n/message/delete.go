package message

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/message"
	messagecrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/message"
)

func (h *Handler) DeleteMessage(ctx context.Context) (*npool.Message, error) {
	info, err := h.GetMessage(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := messagecrud.UpdateSet(
			cli.Message.UpdateOneID(*h.ID),
			&messagecrud.Req{
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
