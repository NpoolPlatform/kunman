package displayname

import (
	"context"
	"fmt"

	appgooddisplaynamemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/name"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgooddisplaynamemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/display/name"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDisplayName(ctx context.Context) error {
	exist, err := appgooddisplaynamemwcli.ExistDisplayNameConds(ctx, &appgooddisplaynamemwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid displayname")
	}
	return nil
}
