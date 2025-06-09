package appfee

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkAppFee(ctx context.Context) error {
	conds := &appfeemwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	}
	handler, err := appfeemw.NewHandler(
		ctx,
		appfeemw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistFeeConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid appfee")
	}
	return nil
}
