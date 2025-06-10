package like

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	likemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/like"
	likemw "github.com/NpoolPlatform/kunman/middleware/good/app/good/like"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkUserLike(ctx context.Context) error {
	conds := &likemwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	}
	handler, err := likemw.NewHandler(
		ctx,
		likemw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistLikeConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid like")
	}
	return nil
}
