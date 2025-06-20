package oauth

import (
	"context"
	"fmt"

	facebook "github.com/NpoolPlatform/kunman/mal/third/facebook"
	github "github.com/NpoolPlatform/kunman/mal/third/github"
	google "github.com/NpoolPlatform/kunman/mal/third/google"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/third/middleware/v1/oauth"
)

func (h *Handler) GetAccessToken(ctx context.Context) (*npool.AccessTokenInfo, error) {
	if h.Code == "" {
		return nil, fmt.Errorf("invalid code")
	}

	switch h.ClientName {
	case basetypes.SignMethod_Github:
		info, err := github.GetAccessToken(h.ClientID, h.ClientSecret, h.Code)
		if err != nil {
			return nil, err
		}
		return info, nil
	case basetypes.SignMethod_Google:
		info, err := google.GetAccessToken(h.ClientID, h.ClientSecret, h.Code, h.RedirectURI)
		if err != nil {
			return nil, err
		}
		return info, nil
	case basetypes.SignMethod_Facebook:
		info, err := facebook.GetAccessToken(h.ClientID, h.ClientSecret, h.Code, h.RedirectURI)
		if err != nil {
			return nil, err
		}
		return info, nil
	default:
		return nil, fmt.Errorf("unsupport oauth")
	}
}

func (h *Handler) GetThirdUserInfo(ctx context.Context) (*npool.ThirdUserInfo, error) {
	if h.AccessToken == "" {
		return nil, fmt.Errorf("invalid accesstoken")
	}

	switch h.ClientName {
	case basetypes.SignMethod_Github:
		info, err := github.GetUserInfo(h.AccessToken)
		if err != nil {
			return nil, err
		}
		return info, nil
	case basetypes.SignMethod_Google:
		info, err := google.GetUserInfo(h.AccessToken)
		if err != nil {
			return nil, err
		}
		return info, nil
	case basetypes.SignMethod_Facebook:
		info, err := facebook.GetUserInfo(h.AccessToken)
		if err != nil {
			return nil, err
		}
		return info, nil
	default:
		return nil, fmt.Errorf("unsupport oauth")
	}
}
