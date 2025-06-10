package poster

import (
	"context"
	"fmt"

	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/poster"
	appgoodpostermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/poster"
	appgoodpostermw "github.com/NpoolPlatform/kunman/middleware/good/app/good/poster"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	posters []*appgoodpostermwpb.Poster
	infos   []*npool.Poster
	apps    map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, poster := range h.posters {
			appIDs = append(appIDs, poster.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, poster := range h.posters {
		info := &npool.Poster{
			ID:          poster.ID,
			EntID:       poster.EntID,
			AppID:       poster.AppID,
			GoodID:      poster.GoodID,
			GoodName:    poster.GoodName,
			GoodType:    poster.GoodType,
			AppGoodID:   poster.AppGoodID,
			AppGoodName: poster.AppGoodName,
			Poster:      poster.Poster,
			Index:       poster.Index,
			CreatedAt:   poster.CreatedAt,
			UpdatedAt:   poster.UpdatedAt,
		}
		app, ok := h.apps[poster.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetPoster(ctx context.Context) (*npool.Poster, error) {
	posterHandler, err := appgoodpostermw.NewHandler(
		ctx,
		appgoodpostermw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	poster, err := posterHandler.GetPoster(ctx)
	if err != nil {
		return nil, err
	}
	if poster == nil {
		return nil, fmt.Errorf("invalid poster")
	}

	handler := &queryHandler{
		Handler: h,
		posters: []*appgoodpostermwpb.Poster{poster},
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

func (h *Handler) GetPosters(ctx context.Context) ([]*npool.Poster, uint32, error) {
	conds := &appgoodpostermwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}

	posterHandler, err := appgoodpostermw.NewHandler(
		ctx,
		appgoodpostermw.WithConds(conds),
		appgoodpostermw.WithOffset(h.Offset),
		appgoodpostermw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	posters, total, err := posterHandler.GetPosters(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(posters) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler: h,
		posters: posters,
		apps:    map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
