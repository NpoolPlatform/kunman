package displaycolor

import (
	"context"
	"fmt"

	appgooddisplaycolormwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/color"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgooddisplaycolormwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/display/color"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDisplayColor(ctx context.Context) error {
	exist, err := appgooddisplaycolormwcli.ExistDisplayColorConds(ctx, &appgooddisplaycolormwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid displaycolor")
	}
	return nil
}
