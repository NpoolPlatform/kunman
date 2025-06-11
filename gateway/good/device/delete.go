package devicetype

import (
	"context"

	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
	devicetypemw "github.com/NpoolPlatform/kunman/middleware/good/device"
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

	typeHandler, err := devicetypemw.NewHandler(
		ctx,
		devicetypemw.WithID(h.ID, true),
		devicetypemw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := typeHandler.DeleteDeviceType(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
