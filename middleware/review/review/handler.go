package review

import (
	"context"
	"fmt"

	reviewtypes "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	npool "github.com/NpoolPlatform/kunman/message/review/middleware/v2/review"
	crud "github.com/NpoolPlatform/kunman/middleware/review/crud/review"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Handler struct {
	crud.Req
	Offset    int32
	Limit     int32
	Conds     *crud.Conds
	ObjectIDs []uuid.UUID
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(appID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appID == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_appID, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}
		h.AppID = &_appID
		return nil
	}
}

func WithReviewerID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid reviewerid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ReviewerID = &_id
		return nil
	}
}

func WithObjectID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid objectid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ObjectID = &_id
		return nil
	}
}

func WithObjectIDs(ids []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, id := range ids {
			uid, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			h.ObjectIDs = append(h.ObjectIDs, uid)
		}
		return nil
	}
}

func WithDomain(domain *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if domain == nil {
			if must {
				return fmt.Errorf("invalid domain")
			}
			return nil
		}
		h.Domain = domain
		return nil
	}
}

func WithTrigger(trigger *reviewtypes.ReviewTriggerType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if trigger == nil {
			if must {
				return fmt.Errorf("invalid trigger")
			}
			return nil
		}
		switch *trigger {
		case reviewtypes.ReviewTriggerType_InsufficientFunds:
		case reviewtypes.ReviewTriggerType_InsufficientGas:
		case reviewtypes.ReviewTriggerType_InsufficientFundsGas:
		case reviewtypes.ReviewTriggerType_LargeAmount:
		case reviewtypes.ReviewTriggerType_AutoReviewed:
		default:
			return fmt.Errorf("invalid trigger type")
		}

		h.Trigger = trigger
		return nil
	}
}

func WithObjectType(_type *reviewtypes.ReviewObjectType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid type")
			}
			return nil
		}
		switch *_type {
		case reviewtypes.ReviewObjectType_ObjectKyc:
		case reviewtypes.ReviewObjectType_ObjectWithdrawal:
		case reviewtypes.ReviewObjectType_ObjectRandomCouponCash:
		default:
			return fmt.Errorf("invalid object type")
		}
		h.ObjectType = _type
		return nil
	}
}

func WithState(state *reviewtypes.ReviewState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid state")
			}
			return nil
		}
		switch *state {
		case reviewtypes.ReviewState_Rejected:
		case reviewtypes.ReviewState_Approved:
		default:
			return fmt.Errorf("invalid review state")
		}
		h.State = state
		return nil
	}
}

func WithMessage(message *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if message == nil {
			if must {
				return fmt.Errorf("invalid message")
			}
			return nil
		}
		h.Message = message
		return nil
	}
}

// nolint
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.ReviewerID != nil {
			id, err := uuid.Parse(conds.GetReviewerID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ReviewerID = &cruder.Cond{Op: conds.GetReviewerID().GetOp(), Val: id}
		}
		if conds.ObjectID != nil {
			id, err := uuid.Parse(conds.GetObjectID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ObjectID = &cruder.Cond{Op: conds.GetObjectID().GetOp(), Val: id}
		}
		if conds.Domain != nil {
			h.Conds.Domain = &cruder.Cond{Op: conds.GetDomain().GetOp(), Val: conds.GetDomain().GetValue()}
		}
		if conds.Trigger != nil {
			h.Conds.Trigger = &cruder.Cond{Op: conds.GetTrigger().GetOp(), Val: reviewtypes.ReviewTriggerType(conds.GetTrigger().GetValue())}
		}
		if conds.ObjectType != nil {
			h.Conds.ObjectType = &cruder.Cond{Op: conds.GetObjectType().GetOp(), Val: reviewtypes.ReviewObjectType(conds.GetObjectType().GetValue())}
		}
		if conds.State != nil {
			h.Conds.State = &cruder.Cond{Op: conds.GetState().GetOp(), Val: reviewtypes.ReviewState(conds.GetState().GetValue())}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{Op: conds.GetEntIDs().GetOp(), Val: ids}
		}

		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
