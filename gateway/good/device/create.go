package devicetype

import (
	"context"

	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
	devicetypemw "github.com/NpoolPlatform/kunman/middleware/good/device"

	"github.com/google/uuid"
)

func (h *Handler) CreateDeviceType(ctx context.Context) (*devicetypemwpb.DeviceType, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := devicetypemw.NewHandler(
		ctx,
		devicetypemw.WithEntID(h.EntID, true),
		devicetypemw.WithType(h.Type, true),
		devicetypemw.WithManufacturerID(h.ManufacturerID, true),
		devicetypemw.WithPowerConsumption(h.PowerConsumption, true),
		devicetypemw.WithShipmentAt(h.ShipmentAt, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateDeviceType(ctx); err != nil {
		return nil, err
	}
	return h.GetDeviceType(ctx)
}
