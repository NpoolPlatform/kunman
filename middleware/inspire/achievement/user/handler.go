package user

import (
	"context"

	common1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/user/common"
)

type Handler struct {
	*common1.Handler
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	_handler, err := common1.NewHandler(ctx, options...)
	if err != nil {
		return nil, err
	}
	h := &Handler{
		Handler: _handler,
	}

	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, h); err != nil {
			return nil, err
		}
	}
	return h, nil
}
