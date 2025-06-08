package default1

import (
	"context"
	"fmt"

	defaultmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/default"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	defaultmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/default"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDefault(ctx context.Context) error {
	exist, err := defaultmwcli.ExistDefaultConds(ctx, &defaultmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid default")
	}
	return nil
}
