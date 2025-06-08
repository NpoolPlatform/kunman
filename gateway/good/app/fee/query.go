package appfee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	appfeemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
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
	info, err := appfeemwcli.GetFee(ctx, *h.AppGoodID)
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
	infos, total, err := appfeemwcli.GetFees(ctx, &appfeemwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit)
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
