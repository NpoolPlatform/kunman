package outofgas

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/outofgas"
	outofgasmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/outofgas"
	outofgasmw "github.com/NpoolPlatform/kunman/middleware/order/outofgas"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	outOfGases []*outofgasmwpb.OutOfGas
	infos      []*npool.OutOfGas
	apps       map[string]*appmwpb.App
	users      map[string]*usermwpb.User
	appGoods   map[string]*appgoodmwpb.Good
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = ordergwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, outOfGas := range h.outOfGases {
			appIDs = append(appIDs, outOfGas.AppID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = ordergwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, outOfGas := range h.outOfGases {
			userIDs = append(userIDs, outOfGas.UserID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getAppGoods(ctx context.Context) (err error) {
	h.appGoods, err = ordergwcommon.GetAppGoods(ctx, func() (appGoodIDs []string) {
		for _, outOfGas := range h.outOfGases {
			appGoodIDs = append(appGoodIDs, outOfGas.AppGoodID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) formalize() {
	for _, outOfGas := range h.outOfGases {
		app, ok := h.apps[outOfGas.AppID]
		if !ok {
			continue
		}
		user, ok := h.users[outOfGas.UserID]
		if !ok {
			continue
		}
		appGood, ok := h.appGoods[outOfGas.AppGoodID]
		if !ok {
			continue
		}
		h.infos = append(h.infos, &npool.OutOfGas{
			ID:           outOfGas.ID,
			EntID:        outOfGas.EntID,
			AppID:        outOfGas.AppID,
			AppName:      app.Name,
			UserID:       outOfGas.UserID,
			EmailAddress: user.EmailAddress,
			PhoneNO:      user.PhoneNO,
			GoodID:       outOfGas.GoodID,
			GoodType:     outOfGas.GoodType,
			GoodName:     appGood.GoodName,
			AppGoodID:    outOfGas.AppGoodID,
			AppGoodName:  appGood.AppGoodName,
			OrderID:      outOfGas.OrderID,
			StartAt:      outOfGas.StartAt,
			EndAt:        outOfGas.EndAt,
			CreatedAt:    outOfGas.CreatedAt,
			UpdatedAt:    outOfGas.UpdatedAt,
		})
	}
}

func (h *Handler) GetOutOfGas(ctx context.Context) (*npool.OutOfGas, error) {
	outOfGasHandler, err := outofgasmw.NewHandler(
		ctx,
		outofgasmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := outOfGasHandler.GetOutOfGas(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid outofgas")
	}
	if h.OrderID != nil && info.OrderID != *h.OrderID {
		return nil, wlog.Errorf("invalid outofgas")
	}

	handler := &queryHandler{
		Handler:    h,
		outOfGases: []*outofgasmwpb.OutOfGas{info},
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
	return handler.infos[0], nil
}

func (h *Handler) GetOutOfGases(ctx context.Context) ([]*npool.OutOfGas, uint32, error) {
	conds := &outofgasmwpb.Conds{}
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

	outOfGasHandler, err := outofgasmw.NewHandler(
		ctx,
		outofgasmw.WithConds(conds),
		outofgasmw.WithOffset(h.Offset),
		outofgasmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	infos, total, err := outOfGasHandler.GetOutOfGases(ctx)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:    h,
		outOfGases: infos,
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
