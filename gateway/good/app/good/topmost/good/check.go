package topmostgood

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	topmostgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good"
	topmostgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkTopMostGood(ctx context.Context) error {
	conds := &topmostgoodmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := topmostgoodmw.NewHandler(
		ctx,
		topmostgoodmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistTopMostGoodConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmostgood")
	}
	return nil
}
