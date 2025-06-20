package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription"
	appsubscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription"
	appsubscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	appSubscriptions []*appsubscriptionmwpb.Subscription
	infos            []*npool.AppSubscription
	apps             map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, appSubscription := range h.appSubscriptions {
			appIDs = append(appIDs, appSubscription.AppID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) formalize() {
	for _, appSubscription := range h.appSubscriptions {
		info := &npool.AppSubscription{
			ID:        appSubscription.ID,
			EntID:     appSubscription.EntID,
			AppID:     appSubscription.AppID,
			GoodID:    appSubscription.GoodID,
			AppGoodID: appSubscription.AppGoodID,

			USDPrice:            appSubscription.USDPrice,
			DurationDisplayType: appSubscription.DurationDisplayType,

			GoodType: appSubscription.GoodType,
			GoodName: appSubscription.GoodName,

			AppGoodName: appSubscription.AppGoodName,
			Banner:      appSubscription.Banner,
			// EnableSetCommission: appSubscription.EnableSetCommission,

			// Descriptions:  appSubscription.Descriptions,
			// Posters:       appSubscription.Posters,
			// DisplayNames:  appSubscription.DisplayNames,
			// DisplayColors: appSubscription.DisplayColors,

			DurationUnits:   appSubscription.DurationUnits,
			DurationQuota:   appSubscription.DurationQuota,
			DailyBonusQuota: appSubscription.DailyBonusQuota,

			// Labels: appSubscription.Labels,

			CreatedAt: appSubscription.CreatedAt,
			UpdatedAt: appSubscription.UpdatedAt,
		}
		app, ok := h.apps[appSubscription.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetSubscription(ctx context.Context) (*npool.AppSubscription, error) {
	prHandler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, err
	}

	appSubscription, err := prHandler.GetSubscription(ctx)
	if err != nil {
		return nil, err
	}
	if appSubscription == nil {
		return nil, wlog.Errorf("invalid appsubscription")
	}

	handler := &queryHandler{
		Handler:          h,
		appSubscriptions: []*appsubscriptionmwpb.Subscription{appSubscription},
		apps:             map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetSubscriptions(ctx context.Context) ([]*npool.AppSubscription, uint32, error) {
	conds := &appsubscriptionmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	prHandler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithConds(conds),
		appsubscriptionmw.WithOffset(h.Offset),
		appsubscriptionmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	appSubscriptions, total, err := prHandler.GetSubscriptions(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(appSubscriptions) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:          h,
		appSubscriptions: appSubscriptions,
		apps:             map[string]*appmwpb.App{},
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
