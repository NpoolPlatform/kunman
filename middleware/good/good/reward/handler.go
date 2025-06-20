package goodreward

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	goodrewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/reward"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	goodrewardcrud.Req
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

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = &_id
		return nil
	}
}

func WithRewardState(e *types.BenefitState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid rewardstate")
			}
			return nil
		}
		switch *e {
		case types.BenefitState_BenefitWait:
		case types.BenefitState_BenefitTransferring:
		case types.BenefitState_BenefitBookKeeping:
		case types.BenefitState_BenefitUserBookKeeping:
		case types.BenefitState_BenefitSimulateBookKeeping:
		case types.BenefitState_BenefitDone:
		case types.BenefitState_BenefitFail:
		default:
			return wlog.Errorf("invalid rewardstate")
		}
		h.RewardState = e
		return nil
	}
}

func WithLastRewardAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return wlog.Errorf("invalid rewardat")
			}
			return nil
		}
		h.LastRewardAt = n
		return nil
	}
}
