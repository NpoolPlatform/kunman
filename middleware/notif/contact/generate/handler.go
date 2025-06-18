package generate

import (
	"context"
	"fmt"

	handler "github.com/NpoolPlatform/kunman/middleware/notif/contact"
)

type Handler struct {
	*handler.Handler
	Subject    *string
	Body       *string
	SenderName *string
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

func WithSubject(subject *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if subject == nil {
			if must {
				return fmt.Errorf("invalid subject")
			}
			return nil
		}
		if *subject == "" {
			return fmt.Errorf("subject is empty")
		}
		h.Subject = subject
		return nil
	}
}

func WithBody(body *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if body == nil {
			if must {
				return fmt.Errorf("invalid body")
			}
			return nil
		}
		if *body == "" {
			return fmt.Errorf("body is empty")
		}
		h.Body = body
		return nil
	}
}

func WithSenderName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("sender name is empty")
		}
		h.SenderName = name
		return nil
	}
}
