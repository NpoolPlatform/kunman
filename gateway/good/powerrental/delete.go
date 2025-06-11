package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/powerrental"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
)

func (h *Handler) DeletePowerRental(ctx context.Context) (*npool.PowerRental, error) {
	handler := &checkHandler{
		Handler: h,
	}
	if err := handler.checkPowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetPowerRental(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	prHandler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithID(h.ID, true),
		powerrentalmw.WithEntID(h.EntID, true),
		powerrentalmw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := prHandler.DeletePowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
