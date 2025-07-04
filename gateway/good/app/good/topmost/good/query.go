package topmostgood

import (
	"context"
	"fmt"

	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good"
	topmostgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
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
	goodHandler, err := topmostgoodmw.NewHandler(
		ctx,
		topmostgoodmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := goodHandler.GetTopMostGood(ctx)
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
	goodHandler, err := topmostgoodmw.NewHandler(
		ctx,
		topmostgoodmw.WithConds(
			&topmostgoodmwpb.Conds{
				AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
			},
		),
		topmostgoodmw.WithOffset(h.Offset),
		topmostgoodmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	infos, total, err := goodHandler.GetTopMostGoods(ctx)
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
