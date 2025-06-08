package required

import (
	"context"
	"fmt"

	requiredappgoodmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkRequired(ctx context.Context) error {
	exist, err := requiredappgoodmwcli.ExistRequiredConds(ctx, &requiredappgoodmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid required")
	}
	return nil
}
