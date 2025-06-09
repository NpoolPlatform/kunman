package common

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	goodusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/gooduser"
	goodusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

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
