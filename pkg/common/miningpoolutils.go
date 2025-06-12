//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/gooduser"
	goodusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	orderusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
	orderusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/orderuser"
	"github.com/google/uuid"
)

func GetMiningPoolOrderUsers(ctx context.Context, orderuserIDs []string) (map[string]*orderusermwpb.OrderUser, error) {
	for _, orderuserID := range orderuserIDs {
		if _, err := uuid.Parse(orderuserID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	coins, _, err := orderusermwcli.GetOrderUsers(ctx, &orderusermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: orderuserIDs},
	}, int32(0), int32(len(orderuserIDs)))
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	orderuserMap := map[string]*orderusermwpb.OrderUser{}
	for _, coin := range coins {
		orderuserMap[coin.EntID] = coin
	}

	return orderuserMap, nil
}

func GetPoolGoodUsers(ctx context.Context, _poolGoodUserIDs []string) (map[string]*goodusermwpb.GoodUser, error) {
	poolGoodUserIDs := []string{}
	for _, poolGoodUserID := range _poolGoodUserIDs {
		if poolGoodUserID == "" {
			continue
		}
		if _, err := uuid.Parse(poolGoodUserID); err != nil {
			return nil, err
		}
		poolGoodUserIDs = append(poolGoodUserIDs, poolGoodUserID)
	}

	conds := &goodusermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: poolGoodUserIDs},
	}
	handler, err := goodusermw.NewHandler(
		ctx,
		goodusermw.WithConds(conds),
		goodusermw.WithOffset(0),
		goodusermw.WithLimit(int32(len(poolGoodUserIDs))),
	)
	if err != nil {
		return nil, err
	}

	goodUsers, _, err := handler.GetGoodUsers(ctx)
	if err != nil {
		return nil, err
	}
	goodUserMap := map[string]*goodusermwpb.GoodUser{}
	for _, goodUser := range goodUsers {
		goodUserMap[goodUser.EntID] = goodUser
	}
	return goodUserMap, nil
}
