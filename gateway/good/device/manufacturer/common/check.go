package common

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	manufacturermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/manufacturer"
	manufacturermw "github.com/NpoolPlatform/kunman/middleware/good/device/manufacturer"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type ManufacturerCheckHandler struct {
	ManufacturerID *string
}

func (h *ManufacturerCheckHandler) CheckManufacturerWithManufacturerID(ctx context.Context, manufacturerID string) error {
	conds := &manufacturermwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: manufacturerID},
	}
	handler, err := manufacturermw.NewHandler(
		ctx,
		manufacturermw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistManufacturerConds(ctx)
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
