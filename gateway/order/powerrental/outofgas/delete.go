package outofgas

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	outofgasgwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/outofgas"
	outofgas1 "github.com/NpoolPlatform/order-gateway/pkg/outofgas"
	powerrentaloutofgasmwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental/outofgas"
)

func (h *Handler) DeleteOutOfGas(ctx context.Context) (*outofgasgwpb.OutOfGas, error) {
	h1, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithEntID(h.EntID, true),
		outofgas1.WithOrderID(h.OrderID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h1.GetOutOfGas(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid outofgas")
	}
	if err := powerrentaloutofgasmwcli.DeleteOutOfGas(ctx, &info.ID, &info.EntID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
