package review

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	npool "github.com/NpoolPlatform/kunman/message/review/middleware/v2/review"

	reviewcrud "github.com/NpoolPlatform/kunman/middleware/review/crud/review"
	"github.com/NpoolPlatform/kunman/middleware/review/db"
	ent "github.com/NpoolPlatform/kunman/middleware/review/db/ent/generated"
	entreview "github.com/NpoolPlatform/kunman/middleware/review/db/ent/generated/review"
)

type queryHandler struct {
	*Handler
	stm   *ent.ReviewSelect
	infos []*npool.Review
	total uint32
}

func (h *queryHandler) selectReviews(stm *ent.ReviewQuery) {
	h.stm = stm.Select(
		entreview.FieldID,
		entreview.FieldEntID,
		entreview.FieldAppID,
		entreview.FieldDomain,
		entreview.FieldReviewerID,
		entreview.FieldObjectID,
		entreview.FieldTrigger,
		entreview.FieldObjectType,
		entreview.FieldState,
		entreview.FieldMessage,
		entreview.FieldCreatedAt,
		entreview.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryReview(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Review.Query().Where(entreview.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entreview.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entreview.EntID(*h.EntID))
	}
	h.selectReviews(stm)
	return nil
}

func (h *queryHandler) queryReviews(ctx context.Context, cli *ent.Client) error {
	stm, err := reviewcrud.SetQueryConds(cli.Review.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectReviews(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Trigger = types.ReviewTriggerType(types.ReviewTriggerType_value[info.TriggerStr])
		info.ObjectType = types.ReviewObjectType(types.ReviewObjectType_value[info.ObjectTypeStr])
		info.State = types.ReviewState(types.ReviewState_value[info.StateStr])
	}
}

func (h *Handler) GetReview(ctx context.Context) (info *npool.Review, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReview(cli); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetReviews(ctx context.Context) ([]*npool.Review, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReviews(_ctx, cli); err != nil {
			return err
		}
		handler.
			stm.
			Order(ent.Desc(entreview.FieldUpdatedAt)).
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
