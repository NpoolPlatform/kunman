package common

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost"
)

type CheckHandler struct {
	goodgwcommon.AppUserCheckHandler
	TopMostID *string
}

func (h *CheckHandler) CheckTopMost(ctx context.Context) error {
	exist, err := topmostmwcli.ExistTopMostConds(ctx, &topmostmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.TopMostID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmost")
	}
	return nil
}
