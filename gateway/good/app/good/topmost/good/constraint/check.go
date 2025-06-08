package constraint

import (
	"context"
	"fmt"

	topmostgoodconstraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good/constraint"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostgoodconstraintmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good/constraint"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkConstraint(ctx context.Context) error {
	exist, err := topmostgoodconstraintmwcli.ExistTopMostGoodConstraintConds(ctx, &topmostgoodconstraintmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmostgoodconstraint")
	}
	return nil
}
