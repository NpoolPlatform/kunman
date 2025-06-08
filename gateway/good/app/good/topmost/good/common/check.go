package common

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

type CheckHandler struct {
	goodgwcommon.AppUserCheckHandler
	TopMostGoodID *string
}

func (h *CheckHandler) CheckTopMostGood(ctx context.Context) error {
	exist, err := topmostgoodmwcli.ExistTopMostGoodConds(ctx, &topmostgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.TopMostGoodID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmostgood")
	}
	return nil
}
