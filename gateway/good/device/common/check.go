package common

import (
	"context"
	"fmt"

	devicetypemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	devicetypemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
)

type DeviceTypeCheckHandler struct {
	DeviceTypeID *string
}

func (h *DeviceTypeCheckHandler) CheckDeviceTypeWithDeviceTypeID(ctx context.Context, deviceTypeID string) error {
	exist, err := devicetypemwcli.ExistDeviceTypeConds(ctx, &devicetypemwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: deviceTypeID},
	})
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
