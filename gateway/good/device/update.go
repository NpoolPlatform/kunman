package devicetype

import (
	"context"

	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
	devicetypemw "github.com/NpoolPlatform/kunman/middleware/good/device"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateDeviceType(ctx context.Context) (*devicetypemwpb.DeviceType, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDeviceType(ctx); err != nil {
		return nil, err
	}

	typeHandler, err := devicetypemw.NewHandler(
		ctx,
		devicetypemw.WithID(h.ID, true),
		devicetypemw.WithEntID(h.EntID, true),
		devicetypemw.WithType(h.Type, true),
		devicetypemw.WithManufacturerID(h.ManufacturerID, true),
		devicetypemw.WithPowerConsumption(h.PowerConsumption, true),
		devicetypemw.WithShipmentAt(h.ShipmentAt, true),
	)
	if err != nil {
		return nil, err
	}

	if err := typeHandler.UpdateDeviceType(ctx); err != nil {
		return nil, err
	}
	return h.GetDeviceType(ctx)
}
