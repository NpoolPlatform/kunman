package goodstm

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
)

type Handler struct {
	CurrentGoodState *types.GoodState
	NextGoodState    *types.GoodState
	Rollback         *bool
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

//nolint:dupl
func WithCurrentGoodState(state *types.GoodState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return wlog.Errorf("invalid current goodstate")
			}
			return nil
		}
		switch *state {
		case types.GoodState_GoodStatePreWait:
		case types.GoodState_GoodStateWait:
		case types.GoodState_GoodStateCreateGoodUser:
		case types.GoodState_GoodStateCheckHashRate:
		case types.GoodState_GoodStateReady:
		case types.GoodState_GoodStateFail:
		default:
			return wlog.Errorf("invalid current goodstate")
		}
		h.CurrentGoodState = state
		return nil
	}
}

//nolint:dupl
func WithNextGoodState(state *types.GoodState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return wlog.Errorf("invalid next goodstate")
			}
			return nil
		}
		switch *state {
		case types.GoodState_GoodStatePreWait:
		case types.GoodState_GoodStateWait:
		case types.GoodState_GoodStateCreateGoodUser:
		case types.GoodState_GoodStateCheckHashRate:
		case types.GoodState_GoodStateReady:
		case types.GoodState_GoodStateFail:
		default:
			return wlog.Errorf("invalid next goodstate")
		}
		h.NextGoodState = state
		return nil
	}
}

func WithRollback(rollback *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Rollback = rollback
		return nil
	}
}
