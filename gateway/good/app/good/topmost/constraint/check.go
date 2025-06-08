package constraint

import (
	"context"
	"fmt"

	topmostconstraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/constraint"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostconstraintmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/constraint"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkConstraint(ctx context.Context) error {
	exist, err := topmostconstraintmwcli.ExistTopMostConstraintConds(ctx, &topmostconstraintmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmostconstraint")
	}
	return nil
}
