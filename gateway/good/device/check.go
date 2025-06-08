package devicetype

import (
	"context"
	"fmt"

	devicetypemwcli "github.com/NpoolPlatform/kunman/middleware/good/device"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDeviceType(ctx context.Context) error {
	exist, err := devicetypemwcli.ExistDeviceTypeConds(ctx, &devicetypemwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid devicetype")
	}
	return nil
}
