package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	powerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/powerrental"
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
	if err := powerrentalmwcli.DeletePowerRental(ctx, h.ID, h.EntID, h.GoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
