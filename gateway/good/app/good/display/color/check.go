package displaycolor

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgooddisplaycolormwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/display/color"
	appgooddisplaycolormw "github.com/NpoolPlatform/kunman/middleware/good/app/good/display/color"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDisplayColor(ctx context.Context) error {
	conds := &appgooddisplaycolormwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}

	mwHandler, err := appgooddisplaycolormw.NewHandler(
		ctx,
		appgooddisplaycolormw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := mwHandler.ExistDisplayColorConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid displaycolor")
	}
	return nil
}
