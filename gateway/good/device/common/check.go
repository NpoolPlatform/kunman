package common

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
	devicetypemw "github.com/NpoolPlatform/kunman/middleware/good/device"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type DeviceTypeCheckHandler struct {
	DeviceTypeID *string
}

func (h *DeviceTypeCheckHandler) CheckDeviceTypeWithDeviceTypeID(ctx context.Context, deviceTypeID string) error {
	conds := &devicetypemwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: deviceTypeID},
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

func (h *DeviceTypeCheckHandler) CheckDeviceType(ctx context.Context) error {
	return h.CheckDeviceTypeWithDeviceTypeID(ctx, *h.DeviceTypeID)
}
