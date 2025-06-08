package common

import (
	"context"
	"fmt"

	goodcommon "github.com/NpoolPlatform/kunman/gateway/good/good/common"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	appgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type AppGoodCheckHandler struct {
	goodgwcommon.AppUserCheckHandler
	goodcommon.GoodCheckHandler
	AppGoodID *string
}

func (h *AppGoodCheckHandler) CheckAppGoodWithAppGoodID(ctx context.Context, appGoodID string) error {
	exist, err := appgoodmwcli.ExistGoodConds(ctx, &appgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: appGoodID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid appgood")
	}
	return nil
}

func (h *AppGoodCheckHandler) CheckAppGood(ctx context.Context) error {
	return h.CheckAppGoodWithAppGoodID(ctx, *h.AppGoodID)
}
