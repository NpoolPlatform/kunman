package constraint

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	topmostgoodconstraintmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good/constraint"
	topmostgoodconstraintmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good/constraint"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkConstraint(ctx context.Context) error {
	conds := &topmostgoodconstraintmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := topmostgoodconstraintmw.NewHandler(
		ctx,
		topmostgoodconstraintmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistConstraintConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmostgoodconstraint")
	}
	return nil
}
