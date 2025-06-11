package required

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	requiredmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"
	requiredmw "github.com/NpoolPlatform/kunman/middleware/good/good/required"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkRequired(ctx context.Context) error {
	conds := &requiredmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	}
	handler, err := requiredmw.NewHandler(
		ctx,
		requiredmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistRequiredConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid required")
	}
	return nil
}
