//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appsubscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription"
	appsubscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetAppSubscriptions(ctx context.Context, appGoodIDs []string) (map[string]*appsubscriptionmwpb.Subscription, error) {
	for _, appGoodID := range appGoodIDs {
		if _, err := uuid.Parse(appGoodID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &appsubscriptionmwpb.Conds{
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appGoodIDs},
	}
	handler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithConds(conds),
		appsubscriptionmw.WithOffset(0),
		appsubscriptionmw.WithLimit(int32(len(appGoodIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	appSubscriptions, _, err := handler.GetSubscriptions(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	appSubscriptionMap := map[string]*appsubscriptionmwpb.Subscription{}
	for _, appSubscription := range appSubscriptions {
		appSubscriptionMap[appSubscription.AppGoodID] = appSubscription
	}
	return appSubscriptionMap, nil
}
