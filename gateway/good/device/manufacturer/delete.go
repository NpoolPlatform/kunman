package manufacturer

import (
	"context"

	manufacturermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteManufacturer(ctx context.Context) (*manufacturermwpb.Manufacturer, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkManufacturer(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetManufacturer(ctx)
	if err != nil {
		return nil, err
	}
	if err := manufacturermwcli.DeleteManufacturer(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
