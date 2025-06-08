package common

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
)

type AppUserCheckHandler struct {
	AppID  *string
	UserID *string
}

func (h *AppUserCheckHandler) CheckAppWithAppID(ctx context.Context, appID string) error {
	exist, err := appmwcli.ExistApp(ctx, appID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid app")
	}
	return nil
}

func (h *AppUserCheckHandler) CheckApp(ctx context.Context) error {
	return h.CheckAppWithAppID(ctx, *h.AppID)
}

func (h *AppUserCheckHandler) CheckUserWithUserID(ctx context.Context, userID string) error {
	exist, err := usermwcli.ExistUser(ctx, *h.AppID, userID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid user")
	}
	return nil
}

func (h *AppUserCheckHandler) CheckUser(ctx context.Context) error {
	return h.CheckUserWithUserID(ctx, *h.UserID)
}
