package devicetype

import (
	"context"

	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
	devicetypemw "github.com/NpoolPlatform/kunman/middleware/good/device"
)

func (h *Handler) GetDeviceType(ctx context.Context) (*devicetypemwpb.DeviceType, error) {
	handler, err := devicetypemw.NewHandler(
		ctx,
		devicetypemw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetDeviceType(ctx)
}

func (h *Handler) GetDeviceTypes(ctx context.Context) ([]*devicetypemwpb.DeviceType, uint32, error) {
	handler, err := devicetypemw.NewHandler(
		ctx,
		devicetypemw.WithConds(&devicetypemwpb.Conds{}),
		devicetypemw.WithOffset(h.Offset),
		devicetypemw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetDeviceTypes(ctx)
}
