package topmostgood

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

type queryHandler struct {
	*Handler
	goods []*topmostgoodmwpb.TopMostGood
	infos []*npool.TopMostGood
	apps  map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, good := range h.goods {
			appIDs = append(appIDs, good.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, good := range h.goods {
		info := &npool.TopMostGood{
			ID:             good.ID,
			EntID:          good.EntID,
			AppID:          good.AppID,
			GoodID:         good.GoodID,
			GoodName:       good.GoodName,
			AppGoodID:      good.AppGoodID,
			AppGoodName:    good.AppGoodName,
			TopMostID:      good.TopMostID,
			TopMostType:    good.TopMostType,
			TopMostTitle:   good.TopMostTitle,
			TopMostMessage: good.TopMostMessage,
			DisplayIndex:   good.DisplayIndex,
			UnitPrice:      good.UnitPrice,
			CreatedAt:      good.CreatedAt,
			UpdatedAt:      good.UpdatedAt,
		}

		app, ok := h.apps[good.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	info, err := topmostgoodmwcli.GetTopMostGood(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid topmostgood")
	}

	handler := &queryHandler{
		Handler: h,
		goods:   []*topmostgoodmwpb.TopMostGood{info},
		apps:    map[string]*appmwpb.App{},
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

func (h *Handler) GetTopMostGoods(ctx context.Context) ([]*npool.TopMostGood, uint32, error) {
	infos, total, err := topmostgoodmwcli.GetTopMostGoods(
		ctx,
		&topmostgoodmwpb.Conds{
			AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		},
		h.Offset,
		h.Limit,
	)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler: h,
		goods:   infos,
		apps:    map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
