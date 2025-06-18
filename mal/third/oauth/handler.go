package oauth

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

type Handler struct {
	ClientID     string
	ClientSecret string
	ClientName   basetypes.SignMethod
	Code         string
	AccessToken  string
	RedirectURI  string
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

func WithClientID(clientID string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ClientID = clientID
		return nil
	}
}

func WithClientSecret(clientSecret string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ClientSecret = clientSecret
		return nil
	}
}

func WithRedirectURI(redirectURI string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RedirectURI = redirectURI
		return nil
	}
}

func WithCode(code string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Code = code
		return nil
	}
}

func WithAccessToken(accessToken string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AccessToken = accessToken
		return nil
	}
}

func WithClientName(clientName basetypes.SignMethod) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch clientName {
		case basetypes.SignMethod_Twitter:
		case basetypes.SignMethod_Github:
		case basetypes.SignMethod_Facebook:
		case basetypes.SignMethod_Linkedin:
		case basetypes.SignMethod_Wechat:
		case basetypes.SignMethod_Google:
		default:
			return fmt.Errorf("clientName %v invalid", clientName)
		}
		h.ClientName = clientName
		return nil
	}
}
