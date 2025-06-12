//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	goodusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/gooduser"
	orderusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	goodusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"
	orderusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/orderuser"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

func GetMiningPoolOrderUsers(ctx context.Context, orderuserIDs []string) (map[string]*orderusermwpb.OrderUser, error) {
	for _, orderuserID := range orderuserIDs {
		if _, err := uuid.Parse(orderuserID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &orderusermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: orderuserIDs},
	}
	handler, err := orderusermw.NewHandler(
		ctx,
		orderusermw.WithConds(conds),
		orderusermw.WithOffset(0),
		orderusermw.WithLimit(int32(len(orderuserIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	coins, _, err := handler.GetOrderUsers(ctx)
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
