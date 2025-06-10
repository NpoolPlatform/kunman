package default1

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	defaultmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/default"
	defaultmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/default"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDefault(ctx context.Context) error {
	conds := &defaultmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := defaultmw.NewHandler(
		ctx,
		defaultmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistDefaultConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid default")
	}
	return nil
}
