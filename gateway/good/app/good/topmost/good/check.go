package topmostgood

import (
	"context"
	"fmt"

	topmostgoodmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	topmostgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkTopMostGood(ctx context.Context) error {
	exist, err := topmostgoodmwcli.ExistTopMostGoodConds(ctx, &topmostgoodmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmostgood")
	}
	return nil
}
