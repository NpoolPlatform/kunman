package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	malfunctionmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/malfunction"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	malfunctionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/malfunction"
	ordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order"
	powerrentalcompensatemwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental/compensate"
	ordermwcli "github.com/NpoolPlatform/kunman/middleware/order/order"
	powerrentalcompensatemwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental/compensate"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	goodMalfunction   *malfunctionmwpb.Malfunction
	order             *ordermwpb.Order
	compensateSeconds uint32
}

func (h *createHandler) getOrder(ctx context.Context) (err error) {
	h.order, err = ordermwcli.GetOrder(ctx, *h.OrderID)
	return wlog.WrapError(err)
}

func (h *createHandler) getGoodMalfunction(ctx context.Context) (err error) {
	conds := &malfunctionmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.CompensateFromID},
	}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	} else {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: h.order.GoodID}
	}
	h.goodMalfunction, err = malfunctionmwcli.GetMalfunctionOnly(ctx, conds)
	if err != nil {
		return err
	}
	if h.goodMalfunction == nil || h.goodMalfunction.CompensateSeconds <= 0 {
		return wlog.Errorf("invalid goodmalfunction")
	}
	h.compensateSeconds = h.goodMalfunction.CompensateSeconds
	return nil
}

func (h *createHandler) getCompensateType(ctx context.Context) error {
	switch *h.CompensateType {
	case types.CompensateType_CompensateMalfunction:
		return h.getGoodMalfunction(ctx)
	case types.CompensateType_CompensateWalfare:
		fallthrough //nolint
	case types.CompensateType_CompensateStarterDelay:
		return wlog.Errorf("not implemented")
	default:
		return wlog.Errorf("invalid compensatetype")
	}
}

func (h *Handler) CreateCompensate(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.OrderID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid ordergood")
	}
	if h.OrderID != nil {
		if err := handler.getOrder(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	if err := handler.getCompensateType(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	return powerrentalcompensatemwcli.CreateCompensate(ctx, &powerrentalcompensatemwpb.CompensateReq{
		EntID:             h.EntID,
		GoodID:            h.GoodID,
		OrderID:           h.OrderID,
		CompensateFromID:  h.CompensateFromID,
		CompensateType:    h.CompensateType,
		CompensateSeconds: &handler.compensateSeconds,
	})
}
