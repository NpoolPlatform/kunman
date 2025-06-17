package oneshot

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	apponeshotmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription/oneshot"
	apponeshotmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription/oneshot"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkOneShot(ctx context.Context) error {
	conds := &apponeshotmwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	}
	handler, err := apponeshotmw.NewHandler(
		ctx,
		apponeshotmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistOneShotConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid apponeshot")
	}
	return nil
}
