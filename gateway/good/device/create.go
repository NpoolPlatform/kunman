package devicetype

import (
	"context"

	devicetypemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"

	"github.com/google/uuid"
)

func (h *Handler) CreateDeviceType(ctx context.Context) (*devicetypemwpb.DeviceType, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := devicetypemwcli.CreateDeviceType(ctx, &devicetypemwpb.DeviceTypeReq{
		EntID:            h.EntID,
		Type:             h.Type,
		ManufacturerID:   h.ManufacturerID,
		PowerConsumption: h.PowerConsumption,
		ShipmentAt:       h.ShipmentAt,
	}); err != nil {
		return nil, err
	}
	return h.GetDeviceType(ctx)
}
