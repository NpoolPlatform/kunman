package common

import (
	"context"
	"fmt"

	appmw "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
)

type AppUserCheckHandler struct {
	AppID  *string
	UserID *string
}

func (h *AppUserCheckHandler) CheckAppWithAppID(ctx context.Context, appID string) error {
	handler, err := appmw.NewHandler(
		ctx,
		appmw.WithEntID(&appID, true),
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
	return nil
}

func (h *AppUserCheckHandler) CheckApp(ctx context.Context) error {
	return h.CheckAppWithAppID(ctx, *h.AppID)
}

func (h *AppUserCheckHandler) CheckUserWithUserID(ctx context.Context, userID string) error {
	handler, err := usermw.NewHandler(
		ctx,
		usermw.WithEntID(&userID, true),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistUser(ctx)
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
