package manufacturer

import (
	"context"
	"fmt"

	manufacturermwcli "github.com/NpoolPlatform/kunman/middleware/good/device/manufacturer"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	manufacturermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/manufacturer"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkManufacturer(ctx context.Context) error {
	exist, err := manufacturermwcli.ExistManufacturerConds(ctx, &manufacturermwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid manufacturer")
	}
	return nil
}
