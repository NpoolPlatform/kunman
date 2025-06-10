package topmost

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	topmostmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost"
	topmostmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkTopMost(ctx context.Context) error {
	conds := &topmostmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := topmostmw.NewHandler(
		ctx,
		topmostmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistTopMostConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmost")
	}
	return nil
}
