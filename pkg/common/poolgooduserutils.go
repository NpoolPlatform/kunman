package common

import (
	"context"

	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	goodusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"

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

	goodUsers, _, err := goodusermwcli.GetGoodUsers(ctx, &goodusermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: poolGoodUserIDs},
	}, int32(0), int32(len(poolGoodUserIDs)))
	if err != nil {
		return nil, err
	}
	goodUserMap := map[string]*goodusermwpb.GoodUser{}
	for _, goodUser := range goodUsers {
		goodUserMap[goodUser.EntID] = goodUser
	}
	return goodUserMap, nil
}
