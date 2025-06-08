package description

import (
	"context"
	"fmt"

	appgooddescriptionmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/description"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgooddescriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/description"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDescription(ctx context.Context) error {
	exist, err := appgooddescriptionmwcli.ExistDescriptionConds(ctx, &appgooddescriptionmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid description")
	}
	return nil
}
