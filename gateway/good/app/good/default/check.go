package default1

import (
	"context"
	"fmt"

	defaultmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/default"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	defaultmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
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
