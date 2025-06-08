package manufacturer

import (
	"context"

	manufacturermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateManufacturer(ctx context.Context) (*manufacturermwpb.Manufacturer, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkManufacturer(ctx); err != nil {
		return nil, err
	}

	if err := manufacturermwcli.UpdateManufacturer(ctx, &manufacturermwpb.ManufacturerReq{
		ID:    h.ID,
		EntID: h.EntID,
		Name:  h.Name,
		Logo:  h.Logo,
	}); err != nil {
		return nil, err
	}
	return h.GetManufacturer(ctx)
}
