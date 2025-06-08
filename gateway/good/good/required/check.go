package required

import (
	"context"
	"fmt"

	requiredmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/required"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	requiredmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkRequired(ctx context.Context) error {
	exist, err := requiredmwcli.ExistRequiredConds(ctx, &requiredmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid required")
	}
	return nil
}
