package send

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

type Handler struct {
	Subject     string
	Content     string
	From        string
	To          string
	ToCCs       []string
	ReplyTos    []string
	AccountType basetypes.SignMethod
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

func WithSubject(subject string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Subject = subject
		return nil
	}
}

func WithContent(content string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Content = content
		return nil
	}
}

func WithFrom(from string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.From = from
		return nil
	}
}

func WithTo(to string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.To = to
		return nil
	}
}

func WithToCCs(toCCs []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if len(toCCs) == 0 {
			return nil
		}
		h.ToCCs = toCCs
		return nil
	}
}

func WithReplyTos(replyTos []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if len(replyTos) == 0 {
			return nil
		}
		h.ReplyTos = replyTos
		return nil
	}
}

func WithAccountType(accountType basetypes.SignMethod) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch accountType {
		case basetypes.SignMethod_Email:
		case basetypes.SignMethod_Mobile:
		default:
			return fmt.Errorf("accountType %v invalid", accountType)
		}
		h.AccountType = accountType
		return nil
	}
}
