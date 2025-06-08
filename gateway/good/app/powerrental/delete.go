package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeletePowerRental(ctx context.Context) (*npool.AppPowerRental, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetPowerRental(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid powerrental")
	}
	if err := apppowerrentalmwcli.DeletePowerRental(ctx, h.ID, h.EntID, h.AppGoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
