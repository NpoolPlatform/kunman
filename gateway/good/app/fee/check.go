package appfee

import (
	"context"
	"fmt"

	appfeemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkAppFee(ctx context.Context) error {
	exist, err := appfeemwcli.ExistFeeConds(ctx, &appfeemwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid appfee")
	}
	return nil
}
