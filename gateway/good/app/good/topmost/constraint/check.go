package constraint

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	topmostconstraintmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/constraint"
	topmostconstraintmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/constraint"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkConstraint(ctx context.Context) error {
	conds := &topmostconstraintmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := topmostconstraintmw.NewHandler(
		ctx,
		topmostconstraintmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistConstraintConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmostconstraint")
	}
	return nil
}
