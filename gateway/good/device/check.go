package devicetype

import (
	"context"
	"fmt"

	devicetypemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
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
