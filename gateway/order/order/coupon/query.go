package coupon

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	allocatedcouponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order/coupon"
	ordercouponmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/coupon"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	ordercouponcli "github.com/NpoolPlatform/kunman/middleware/order/order/coupon"
)

type queryHandler struct {
	*Handler
	orderCoupons     []*ordercouponmwpb.OrderCoupon
	infos            []*npool.OrderCoupon
	apps             map[string]*appmwpb.App
	users            map[string]*usermwpb.User
	appGoods         map[string]*appgoodmwpb.Good
	allocatedCoupons map[string]*allocatedcouponmwpb.Coupon
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = ordergwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, orderCoupon := range h.orderCoupons {
			appIDs = append(appIDs, orderCoupon.AppID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = ordergwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, orderCoupon := range h.orderCoupons {
			userIDs = append(userIDs, orderCoupon.UserID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getAppGoods(ctx context.Context) (err error) {
	h.appGoods, err = ordergwcommon.GetAppGoods(ctx, func() (appGoodIDs []string) {
		for _, orderCoupon := range h.orderCoupons {
			appGoodIDs = append(appGoodIDs, orderCoupon.AppGoodID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getAllocatedCoupons(ctx context.Context) (err error) {
	h.allocatedCoupons, err = ordergwcommon.GetAllocatedCoupons(ctx, func() (allocatedCouponIDs []string) {
		for _, orderCoupon := range h.orderCoupons {
			allocatedCouponIDs = append(allocatedCouponIDs, orderCoupon.CouponID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize(ctx context.Context) { //nolint
	for _, orderCoupon := range h.orderCoupons {
		info := &npool.OrderCoupon{
			ID:                orderCoupon.ID,
			EntID:             orderCoupon.EntID,
			AppID:             orderCoupon.AppID,
			UserID:            orderCoupon.UserID,
			GoodID:            orderCoupon.GoodID,
			GoodType:          orderCoupon.GoodType,
			AppGoodID:         orderCoupon.AppGoodID,
			OrderID:           orderCoupon.OrderID,
			AllocatedCouponID: orderCoupon.CouponID,
			// TODO: coupon info
			CreatedAt: orderCoupon.CreatedAt,
			UpdatedAt: orderCoupon.UpdatedAt,
		}
		if app, ok := h.apps[orderCoupon.AppID]; ok {
			info.AppName = app.Name
		}
		if user, ok := h.users[orderCoupon.UserID]; ok {
			info.EmailAddress = user.EmailAddress
			info.PhoneNO = user.PhoneNO
		}
		if appGood, ok := h.appGoods[orderCoupon.AppGoodID]; ok {
			info.GoodName = appGood.GoodName
			info.AppGoodName = appGood.AppGoodName
		}
		if coupon, ok := h.allocatedCoupons[orderCoupon.CouponID]; ok {
			info.CouponName = coupon.CouponName
			info.CouponType = coupon.CouponType
			info.Denomination = coupon.Denomination
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetOrderCoupons(ctx context.Context) ([]*npool.OrderCoupon, uint32, error) {
	conds := &ordercouponmwpb.Conds{}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	orderCoupons, total, err := ordercouponcli.GetOrderCoupons(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(orderCoupons) == 0 {
		return nil, 0, nil
	}

	handler := &queryHandler{
		Handler:      h,
		orderCoupons: orderCoupons,
		infos:        []*npool.OrderCoupon{},
		users:        map[string]*usermwpb.User{},
		appGoods:     map[string]*appgoodmwpb.Good{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getAppGoods(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getAllocatedCoupons(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize(ctx)

	return handler.infos, total, nil
}
