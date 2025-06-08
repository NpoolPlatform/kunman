package subscriber

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/subscriber"
	subscribercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/subscriber"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateSubscriber(ctx context.Context) (*npool.Subscriber, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	// TODO: deduplicate

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := subscribercrud.SetQueryConds(
			cli.Subscriber.Query(),
			&subscribercrud.Conds{
				AppID:        &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				EmailAddress: &cruder.Cond{Op: cruder.EQ, Val: *h.EmailAddress},
			},
		)
		if err != nil {
			return err
		}

		info, err := stm.Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info != nil {
			h.EntID = &info.EntID
			return nil
		}

		if _, err := subscribercrud.CreateSet(
			cli.Subscriber.Create(),
			&subscribercrud.Req{
				EntID:        h.EntID,
				AppID:        h.AppID,
				EmailAddress: h.EmailAddress,
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
