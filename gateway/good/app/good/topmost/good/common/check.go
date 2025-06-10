package common

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	topmostgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good"
	topmostgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type CheckHandler struct {
	goodgwcommon.AppUserCheckHandler
	TopMostGoodID *string
}

func (h *CheckHandler) CheckTopMostGood(ctx context.Context) error {
	conds := &topmostgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.TopMostGoodID},
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
