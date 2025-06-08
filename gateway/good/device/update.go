package devicetype

import (
	"context"

	devicetypemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
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

	if err := devicetypemwcli.UpdateDeviceType(ctx, &devicetypemwpb.DeviceTypeReq{
		ID:               h.ID,
		Type:             h.Type,
		ManufacturerID:   h.ManufacturerID,
		PowerConsumption: h.PowerConsumption,
		ShipmentAt:       h.ShipmentAt,
	}); err != nil {
		return nil, err
	}
	return h.GetDeviceType(ctx)
}
