package devicetype

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
	devicetypemw "github.com/NpoolPlatform/kunman/middleware/good/device"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDeviceType(ctx context.Context) error {
	conds := &devicetypemwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	}
	handler, err := devicetypemw.NewHandler(
		ctx,
		devicetypemw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistDeviceTypeConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid devicetype")
	}
	return nil
}
