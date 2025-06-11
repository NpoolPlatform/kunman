package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
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

	prHandler, err := apppowerrentalmw.NewHandler(
		ctx,
		apppowerrentalmw.WithID(h.ID, true),
		apppowerrentalmw.WithEntID(h.EntID, true),
		apppowerrentalmw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := prHandler.DeletePowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
