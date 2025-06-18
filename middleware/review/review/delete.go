package review

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/review/middleware/v2/review"
	reviewcrud "github.com/NpoolPlatform/kunman/middleware/review/crud/review"
	"github.com/NpoolPlatform/kunman/middleware/review/db"
	ent "github.com/NpoolPlatform/kunman/middleware/review/db/ent/generated"
)

func (h *Handler) DeleteReview(ctx context.Context) (info *npool.Review, err error) {
	info, err = h.GetReview(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	if h.ID == nil {
		h.ID = &info.ID
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := reviewcrud.UpdateSet(
			cli.Review.UpdateOneID(*h.ID),
			&reviewcrud.Req{
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
