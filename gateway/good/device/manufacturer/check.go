package manufacturer

import (
	"context"
	"fmt"

	manufacturermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
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
