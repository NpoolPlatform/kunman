package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"
	apponeshotmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription/oneshot"
	apponeshotmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription/oneshot"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	appOneShots []*apponeshotmwpb.OneShot
	infos       []*npool.AppOneShot
	apps        map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, appOneShot := range h.appOneShots {
			appIDs = append(appIDs, appOneShot.AppID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) formalize() {
	for _, appOneShot := range h.appOneShots {
		info := &npool.AppOneShot{
			ID:        appOneShot.ID,
			EntID:     appOneShot.EntID,
			AppID:     appOneShot.AppID,
			GoodID:    appOneShot.GoodID,
			AppGoodID: appOneShot.AppGoodID,

			USDPrice: appOneShot.USDPrice,

			GoodType: appOneShot.GoodType,
			GoodName: appOneShot.GoodName,

			AppGoodName: appOneShot.AppGoodName,
			Banner:      appOneShot.Banner,
			// EnableSetCommission: appOneShot.EnableSetCommission,

			// Descriptions:  appOneShot.Descriptions,
			// Posters:       appOneShot.Posters,
			// DisplayNames:  appOneShot.DisplayNames,
			// DisplayColors: appOneShot.DisplayColors,

			Quota: appOneShot.Quota,

			// Labels: appOneShot.Labels,

			CreatedAt: appOneShot.CreatedAt,
			UpdatedAt: appOneShot.UpdatedAt,
		}
		app, ok := h.apps[appOneShot.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetOneShot(ctx context.Context) (*npool.AppOneShot, error) {
	prHandler, err := apponeshotmw.NewHandler(
		ctx,
		apponeshotmw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, err
	}

	appOneShot, err := prHandler.GetOneShot(ctx)
	if err != nil {
		return nil, err
	}
	if appOneShot == nil {
		return nil, wlog.Errorf("invalid apponeshot")
	}

	handler := &queryHandler{
		Handler:     h,
		appOneShots: []*apponeshotmwpb.OneShot{appOneShot},
		apps:        map[string]*appmwpb.App{},
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

func (h *Handler) GetOneShots(ctx context.Context) ([]*npool.AppOneShot, uint32, error) {
	conds := &apponeshotmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	prHandler, err := apponeshotmw.NewHandler(
		ctx,
		apponeshotmw.WithConds(conds),
		apponeshotmw.WithOffset(h.Offset),
		apponeshotmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	appOneShots, total, err := prHandler.GetOneShots(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(appOneShots) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:     h,
		appOneShots: appOneShots,
		apps:        map[string]*appmwpb.App{},
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
