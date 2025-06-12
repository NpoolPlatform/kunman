package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	cashcontrolcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/app/cashcontrol"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	coupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	inspiretypes "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/app/cashcontrol"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) getCoupon(ctx context.Context) error {
	couponID := h.CouponID.String()
	handler, err := coupon1.NewHandler(
		ctx,
		coupon1.WithEntID(&couponID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	coupon, err := handler.GetCoupon(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if coupon == nil {
		return wlog.Errorf("invalid coupon")
	}
	if coupon.AppID != h.AppID.String() {
		return wlog.Errorf("invalid coupon")
	}
	if coupon.CouponType != inspiretypes.CouponType_FixAmount {
		return wlog.Errorf("invalid coupon type")
	}
	return nil
}

func (h *createHandler) createCashControl(ctx context.Context, cli *ent.Client) error {
	if _, err := cashcontrolcrud.CreateSet(
		cli.CashControl.Create(),
		&cashcontrolcrud.Req{
			EntID:       h.EntID,
			AppID:       h.AppID,
			CouponID:    h.CouponID,
			ControlType: h.ControlType,
			Value:       h.Value,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateCashControl(ctx context.Context) (*npool.CashControl, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.getCoupon(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	h.Conds = &cashcontrolcrud.Conds{
		AppID:       &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		CouponID:    &cruder.Cond{Op: cruder.EQ, Val: *h.CouponID},
		ControlType: &cruder.Cond{Op: cruder.EQ, Val: *h.ControlType},
	}
	exist, err := h.ExistCashControlConds(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if exist {
		return nil, wlog.Errorf("control type %v already exist", *h.ControlType)
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createCashControl(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetCashControl(ctx)
}
