package common

import (
	"context"
	"fmt"

	manufacturermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
)

type ManufacturerCheckHandler struct {
	ManufacturerID *string
}

func (h *ManufacturerCheckHandler) CheckManufacturerWithManufacturerID(ctx context.Context, manufacturerID string) error {
	exist, err := manufacturermwcli.ExistManufacturerConds(ctx, &manufacturermwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: manufacturerID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid manufacturer")
	}
	return nil
}

func (h *ManufacturerCheckHandler) CheckManufacturer(ctx context.Context) error {
	return h.CheckManufacturerWithManufacturerID(ctx, *h.ManufacturerID)
}
