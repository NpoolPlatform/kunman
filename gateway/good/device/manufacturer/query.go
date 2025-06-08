package manufacturer

import (
	"context"

	manufacturermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	manufacturermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/manufacturer"
)

func (h *Handler) GetManufacturer(ctx context.Context) (*manufacturermwpb.Manufacturer, error) {
	return manufacturermwcli.GetManufacturer(ctx, *h.EntID)
}

func (h *Handler) GetManufacturers(ctx context.Context) ([]*manufacturermwpb.Manufacturer, uint32, error) {
	return manufacturermwcli.GetManufacturers(ctx, &manufacturermwpb.Conds{}, h.Offset, h.Limit)
}
