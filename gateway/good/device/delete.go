package devicetype

import (
	"context"

	devicetypemwcli "github.com/NpoolPlatform/kunman/middleware/good/device"
	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDeviceType(ctx context.Context) (*devicetypemwpb.DeviceType, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDeviceType(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetDeviceType(ctx)
	if err != nil {
		return nil, err
	}
	if err := devicetypemwcli.DeleteDeviceType(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
