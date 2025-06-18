package review

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	npool "github.com/NpoolPlatform/kunman/message/review/middleware/v2/review"
	crud "github.com/NpoolPlatform/kunman/middleware/review/crud/review"
	"github.com/NpoolPlatform/kunman/middleware/review/db"
	ent "github.com/NpoolPlatform/kunman/middleware/review/db/ent/generated"

	"github.com/google/uuid"
)

func (h *Handler) validateReview(ctx context.Context) error {
	info, err := h.GetReview(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("review not found")
	}
	switch info.State {
	case types.ReviewState_DefaultReviewState:
	case types.ReviewState_Wait:
	default:
		return fmt.Errorf("current review state can not be updated")
	}

	if info.ReviewerID != uuid.Nil.String() && h.ReviewerID != nil {
		return fmt.Errorf("permission denied")
	}
	if h.State != nil && *h.State == types.ReviewState_Rejected && h.Message == nil {
		return fmt.Errorf("message is empty")
	}
	return nil
}

func (h *Handler) UpdateReview(ctx context.Context) (info *npool.Review, err error) {
	if err := h.validateReview(ctx); err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.Review.UpdateOneID(*h.ID),
			&crud.Req{
				State:      h.State,
				ReviewerID: h.ReviewerID,
				Message:    h.Message,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetReview(ctx)
}
