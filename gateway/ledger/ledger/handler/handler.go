package handler

import (
	"context"
	"fmt"

	appmw "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/google/uuid"
)

type Handler struct {
	AppID        *string
	UserID       *string
	StartAt      *uint32
	EndAt        *uint32
	SimulateOnly *bool
	Offset       int32
	Limit        int32
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithAppID(appID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appID == nil {
			if must {
				return fmt.Errorf("invalid app id")
			}
			return nil
		}
		if _, err := uuid.Parse(*appID); err != nil {
			return err
		}

		handler, err := appmw.NewHandler(
			ctx,
			appmw.WithEntID(appID, true),
		)
		if err != nil {
			return err
		}

		exist, err := handler.ExistApp(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid app")
		}
		h.AppID = appID
		return nil
	}
}

func WithUserID(userID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			if must {
				return fmt.Errorf("invalid user id")
			}
			return nil
		}
		_, err := uuid.Parse(*userID)
		if err != nil {
			return err
		}
		h.UserID = userID
		return nil
	}
}

func WithStartAt(startAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.StartAt = startAt
		return nil
	}
}

func WithEndAt(endAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EndAt = endAt
		return nil
	}
}

func WithSimulateOnly(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid simulateonly")
			}
			return nil
		}
		h.SimulateOnly = value
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
