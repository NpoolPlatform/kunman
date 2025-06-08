package common

import (
	"context"

	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appmw "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetApps(ctx context.Context, appIDs []string) (map[string]*appmwpb.App, error) {
	for _, appID := range appIDs {
		if _, err := uuid.Parse(appID); err != nil {
			return nil, err
		}
	}

	conds := &appmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appIDs},
	}
	handler, err := appmw.NewHandler(
		ctx,
		appmw.WithConds(conds),
		appmw.WithOffset(0),
		appmw.WithLimit(int32(len(appIDs))),
	)
	if err != nil {
		return nil, err
	}

	apps, _, err := handler.GetApps(ctx)
	if err != nil {
		return nil, err
	}
	appMap := map[string]*appmwpb.App{}
	for _, app := range apps {
		appMap[app.EntID] = app
	}
	return appMap, nil
}

func GetUsers(ctx context.Context, userIDs []string) (map[string]*usermwpb.User, error) {
	for _, userID := range userIDs {
		if _, err := uuid.Parse(userID); err != nil {
			return nil, err
		}
	}

	conds := &usermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: userIDs},
	}
	handler, err := usermw.NewHandler(
		ctx,
		usermw.WithConds(conds),
		usermw.WithOffset(0),
		usermw.WithLimit(int32(len(userIDs))),
	)
	if err != nil {
		return nil, err
	}

	users, _, err := handler.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	userMap := map[string]*usermwpb.User{}
	for _, user := range users {
		userMap[user.EntID] = user
	}
	return userMap, nil
}
