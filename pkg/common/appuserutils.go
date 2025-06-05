package common

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func GetApps(ctx context.Context, appIDs []string) (map[string]*appmwpb.App, error) {
	for _, appID := range appIDs {
		if _, err := uuid.Parse(appID); err != nil {
			return nil, err
		}
	}

	apps, _, err := appmwcli.GetApps(ctx, &appmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appIDs},
	}, int32(0), int32(len(appIDs)))
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

	users, _, err := usermwcli.GetUsers(ctx, &usermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: userIDs},
	}, 0, int32(len(userIDs)))
	if err != nil {
		return nil, err
	}
	userMap := map[string]*usermwpb.User{}
	for _, user := range users {
		userMap[user.EntID] = user
	}
	return userMap, nil
}
