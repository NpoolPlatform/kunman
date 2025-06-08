package simulate

import (
	"context"

	simulatemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental/simulate"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/powerrental/simulate"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteSimulate(ctx context.Context) (*npool.Simulate, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSimulate(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetSimulate(ctx)
	if err != nil {
		return nil, err
	}
	if err := simulatemwcli.DeleteSimulate(ctx, h.ID, h.EntID, h.AppGoodID); err != nil {
		return nil, err
	}
	return info, nil
}
