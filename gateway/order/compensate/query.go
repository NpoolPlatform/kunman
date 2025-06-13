package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/compensate"
	compensatemwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/compensate"
	compensatemw "github.com/NpoolPlatform/kunman/middleware/order/compensate"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	compensates []*compensatemwpb.Compensate
	infos       []*npool.Compensate
	apps        map[string]*appmwpb.App
	users       map[string]*usermwpb.User
	appGoods    map[string]*appgoodmwpb.Good
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = ordergwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, compensate := range h.compensates {
			appIDs = append(appIDs, compensate.AppID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = ordergwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, compensate := range h.compensates {
			userIDs = append(userIDs, compensate.UserID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getAppGoods(ctx context.Context) (err error) {
	h.appGoods, err = ordergwcommon.GetAppGoods(ctx, func() (appGoodIDs []string) {
		for _, compensate := range h.compensates {
			appGoodIDs = append(appGoodIDs, compensate.AppGoodID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) formalize() {
	for _, compensate := range h.compensates {
		app, ok := h.apps[compensate.AppID]
		if !ok {
			continue
		}
		user, ok := h.users[compensate.UserID]
		if !ok {
			continue
		}
		appGood, ok := h.appGoods[compensate.AppGoodID]
		if !ok {
			continue
		}
		h.infos = append(h.infos, &npool.Compensate{
			ID:                compensate.ID,
			EntID:             compensate.EntID,
			AppID:             compensate.AppID,
			AppName:           app.Name,
			UserID:            compensate.UserID,
			EmailAddress:      user.EmailAddress,
			PhoneNO:           user.PhoneNO,
			GoodID:            compensate.GoodID,
			GoodType:          compensate.GoodType,
			GoodName:          appGood.GoodName,
			AppGoodID:         compensate.AppGoodID,
			AppGoodName:       appGood.AppGoodName,
			OrderID:           compensate.OrderID,
			CompensateFromID:  compensate.CompensateFromID,
			CompensateType:    compensate.CompensateType,
			CompensateSeconds: compensate.CompensateSeconds,
			CreatedAt:         compensate.CreatedAt,
			UpdatedAt:         compensate.UpdatedAt,
			// TODO: add compensate name
		})
	}
}

func (h *Handler) GetCompensate(ctx context.Context) (*npool.Compensate, error) {
	compensateHandler, err := compensatemw.NewHandler(
		ctx,
		compensatemw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := compensateHandler.GetCompensate(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid compensate")
	}
	if h.OrderID != nil && info.OrderID != *h.OrderID {
		return nil, wlog.Errorf("invalid compensate")
	}

	handler := &queryHandler{
		Handler:     h,
		compensates: []*compensatemwpb.Compensate{info},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAppGoods(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}
	return handler.infos[0], nil
}

func (h *Handler) GetCompensates(ctx context.Context) ([]*npool.Compensate, uint32, error) {
	conds := &compensatemwpb.Conds{}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	if h.OrderID != nil {
		conds.OrderID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderID}
	}

	compensateHandler, err := compensatemw.NewHandler(
		ctx,
		compensatemw.WithConds(conds),
		compensatemw.WithOffset(h.Offset),
		compensatemw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	infos, total, err := compensateHandler.GetCompensates(ctx)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:     h,
		compensates: infos,
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

	handler.formalize()
	return handler.infos, total, nil
}
