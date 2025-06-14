package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	malfunctionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/malfunction"
	ordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order"
	malfunctionmw "github.com/NpoolPlatform/kunman/middleware/good/good/malfunction"
	ordermw "github.com/NpoolPlatform/kunman/middleware/order/order"
	powerrentalcompensatemw "github.com/NpoolPlatform/kunman/middleware/order/powerrental/compensate"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	goodMalfunction   *malfunctionmwpb.Malfunction
	order             *ordermwpb.Order
	compensateSeconds uint32
}

func (h *createHandler) getOrder(ctx context.Context) (err error) {
	handler, err := ordermw.NewHandler(
		ctx,
		ordermw.WithEntID(h.OrderID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.order, err = handler.GetOrder(ctx)
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

	handler, err := malfunctionmw.NewHandler(
		ctx,
		malfunctionmw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.goodMalfunction, err = handler.GetMalfunctionOnly(ctx)
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

	compensateHandler, err := powerrentalcompensatemw.NewHandler(
		ctx,
		powerrentalcompensatemw.WithEntID(h.EntID, true),
		powerrentalcompensatemw.WithGoodID(h.GoodID, true),
		powerrentalcompensatemw.WithOrderID(h.OrderID, true),
		powerrentalcompensatemw.WithCompensateFromID(h.CompensateFromID, true),
		powerrentalcompensatemw.WithCompensateType(h.CompensateType, true),
		powerrentalcompensatemw.WithCompensateSeconds(&handler.compensateSeconds, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return compensateHandler.CreateCompensate(ctx)
}
