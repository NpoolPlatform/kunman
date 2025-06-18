package verify

import "context"

type Handler struct {
	RecaptchaToken string
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

func WithRecaptchaToken(recaptchaToken string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RecaptchaToken = recaptchaToken
		return nil
	}
}
