package manufacturer

import (
	"context"

	manufacturermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"

	"github.com/google/uuid"
)

func (h *Handler) CreateManufacturer(ctx context.Context) (*manufacturermwpb.Manufacturer, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := manufacturermwcli.CreateManufacturer(ctx, &manufacturermwpb.ManufacturerReq{
		EntID: h.EntID,
		Name:  h.Name,
		Logo:  h.Logo,
	}); err != nil {
		return nil, err
	}
	return h.GetManufacturer(ctx)
}
