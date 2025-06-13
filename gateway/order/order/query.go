package order

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order"
	ordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	ordercli "github.com/NpoolPlatform/kunman/middleware/order/order"
)

type queryHandler struct {
	*Handler
	orders   []*ordermwpb.Order
	infos    []*npool.Order
	users    map[string]*usermwpb.User
	appGoods map[string]*appgoodmwpb.Good
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = ordergwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, order := range h.orders {
			userIDs = append(userIDs, order.UserID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getAppGoods(ctx context.Context) (err error) {
	h.appGoods, err = ordergwcommon.GetAppGoods(ctx, func() (appGoodIDs []string) {
		for _, order := range h.orders {
			appGoodIDs = append(appGoodIDs, order.AppGoodID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) formalize(ctx context.Context) { //nolint
	for _, order := range h.orders {
		info := &npool.Order{
			ID:            order.ID,
			EntID:         order.EntID,
			AppID:         order.AppID,
			UserID:        order.UserID,
			GoodID:        order.GoodID,
			GoodType:      order.GoodType,
			AppGoodID:     order.AppGoodID,
			ParentOrderID: order.ParentOrderID,
			OrderType:     order.OrderType,
			PaymentType:   order.PaymentType,
			CreateMethod:  order.CreateMethod,
			Simulate:      order.Simulate,
			OrderState:    order.OrderState,
			StartMode:     order.StartMode,
			StartAt:       order.StartAt,
			LastBenefitAt: order.LastBenefitAt,
			BenefitState:  order.BenefitState,
			CreatedAt:     order.CreatedAt,
			UpdatedAt:     order.UpdatedAt,
		}
		if user, ok := h.users[order.UserID]; ok {
			info.EmailAddress = user.EmailAddress
			info.PhoneNO = user.PhoneNO
		}
		appGood, ok := h.appGoods[order.AppGoodID]
		if ok {
			info.GoodName = appGood.GoodName
			info.AppGoodName = appGood.AppGoodName
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetOrders(ctx context.Context) ([]*npool.Order, uint32, error) {
	conds := &ordermwpb.Conds{}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	orders, total, err := ordercli.GetOrders(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(orders) == 0 {
		return nil, 0, nil
	}

	handler := &queryHandler{
		Handler:  h,
		orders:   orders,
		infos:    []*npool.Order{},
		users:    map[string]*usermwpb.User{},
		appGoods: map[string]*appgoodmwpb.Good{},
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getAppGoods(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize(ctx)

	return handler.infos, total, nil
}
