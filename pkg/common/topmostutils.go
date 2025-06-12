//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"

	"github.com/google/uuid"
)

func GetTopMosts(ctx context.Context, topMostIDs []string) (map[string]*topmostmwpb.TopMost, error) {
	for _, topMostID := range topMostIDs {
		if _, err := uuid.Parse(topMostID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	topMosts, _, err := topmostmwcli.GetTopMosts(ctx, &topmostmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: topMostIDs},
	}, int32(0), int32(len(topMostIDs)))
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	topMostMap := map[string]*topmostmwpb.TopMost{}
	for _, topMost := range topMosts {
		topMostMap[topMost.EntID] = topMost
	}
	return topMostMap, nil
}
