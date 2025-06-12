//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	topmostmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost"
	topmostmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetTopMosts(ctx context.Context, topMostIDs []string) (map[string]*topmostmwpb.TopMost, error) {
	for _, topMostID := range topMostIDs {
		if _, err := uuid.Parse(topMostID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &topmostmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: topMostIDs},
	}
	handler, err := topmostmw.NewHandler(
		ctx,
		topmostmw.WithConds(conds),
		topmostmw.WithOffset(0),
		topmostmw.WithLimit(int32(len(topMostIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	topMosts, _, err := handler.GetTopMosts(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	topMostMap := map[string]*topmostmwpb.TopMost{}
	for _, topMost := range topMosts {
		topMostMap[topMost.EntID] = topMost
	}
	return topMostMap, nil
}
