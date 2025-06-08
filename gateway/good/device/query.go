package devicetype

import (
	"context"

	devicetypemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
)

func (h *Handler) GetDeviceType(ctx context.Context) (*devicetypemwpb.DeviceType, error) {
	return devicetypemwcli.GetDeviceType(ctx, *h.EntID)
}

func (h *Handler) GetDeviceTypes(ctx context.Context) ([]*devicetypemwpb.DeviceType, uint32, error) {
	return devicetypemwcli.GetDeviceTypes(ctx, &devicetypemwpb.Conds{}, h.Offset, h.Limit)
}
