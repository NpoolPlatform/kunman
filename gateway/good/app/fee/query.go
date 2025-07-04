package appfee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	fees  []*appfeemwpb.Fee
	infos []*npool.AppFee
	apps  map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, fee := range h.fees {
			appIDs = append(appIDs, fee.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, fee := range h.fees {
		app, ok := h.apps[fee.AppID]
		if !ok {
			continue
		}
		h.infos = append(h.infos, &npool.AppFee{
			ID:                      fee.ID,
			EntID:                   fee.EntID,
			AppID:                   fee.AppID,
			AppName:                 app.Name,
			GoodID:                  fee.GoodID,
			GoodName:                fee.GoodName,
			AppGoodID:               fee.AppGoodID,
			AppGoodName:             fee.AppGoodName,
			ProductPage:             fee.ProductPage,
			Banner:                  fee.Banner,
			UnitValue:               fee.UnitValue,
			MinOrderDurationSeconds: fee.MinOrderDurationSeconds,
			GoodType:                fee.GoodType,
			SettlementType:          fee.SettlementType,
			DurationDisplayType:     fee.DurationDisplayType,
			CancelMode:              fee.CancelMode,
			CreatedAt:               fee.CreatedAt,
			UpdatedAt:               fee.UpdatedAt,
		})
	}
}

func (h *Handler) GetAppFee(ctx context.Context) (*npool.AppFee, error) {
	feeHandler, err := appfeemw.NewHandler(
		ctx,
		appfeemw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := feeHandler.GetFee(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid fee")
	}
	if info.AppID != *h.AppID {
		return nil, wlog.Errorf("permission denied")
	}
	handler := &queryHandler{
		Handler: h,
		fees:    []*appfeemwpb.Fee{info},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetAppFees(ctx context.Context) ([]*npool.AppFee, uint32, error) {
	conds := &appfeemwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	feeHandler, err := appfeemw.NewHandler(
		ctx,
		appfeemw.WithConds(conds),
		appfeemw.WithOffset(h.Offset),
		appfeemw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	infos, total, err := feeHandler.GetFees(ctx)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, 0, nil
	}
	handler := &queryHandler{
		Handler: h,
		fees:    infos,
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos, total, nil
}
