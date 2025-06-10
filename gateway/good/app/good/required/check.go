package required

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	requiredappgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkRequired(ctx context.Context) error {
	conds := &requiredappgoodmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := requiredappgoodmw.NewHandler(
		ctx,
		requiredappgoodmw.WithConds(conds),
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
