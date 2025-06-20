package ledger

import (
	"context"

	"github.com/NpoolPlatform/kunman/gateway/ledger/ledger/handler"
)

type Handler struct {
	*handler.Handler
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	_handler, err := handler.NewHandler(ctx, options...)
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
