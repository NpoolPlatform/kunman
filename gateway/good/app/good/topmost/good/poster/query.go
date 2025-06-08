package poster

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	postermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good/poster"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good/poster"
	postermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/poster"
)

type queryHandler struct {
	*Handler
	posters []*postermwpb.Poster
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
			ID:               poster.ID,
			EntID:            poster.EntID,
			AppID:            poster.AppID,
			TopMostID:        poster.TopMostID,
			TopMostType:      poster.TopMostType,
			TopMostTitle:     poster.TopMostTitle,
			TopMostMessage:   poster.TopMostMessage,
			TopMostTargetUrl: poster.TopMostTargetUrl,
			TopMostGoodID:    poster.TopMostGoodID,
			AppGoodID:        poster.AppGoodID,
			AppGoodName:      poster.AppGoodName,
			Poster:           poster.Poster,
			Index:            poster.Index,
			CreatedAt:        poster.CreatedAt,
			UpdatedAt:        poster.UpdatedAt,
		}

		app, ok := h.apps[poster.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetPoster(ctx context.Context) (*npool.Poster, error) {
	info, err := postermwcli.GetPoster(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid poster")
	}

	handler := &queryHandler{
		Handler: h,
		posters: []*postermwpb.Poster{info},
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
	infos, total, err := postermwcli.GetPosters(
		ctx,
		&postermwpb.Conds{
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
		posters: infos,
		apps:    map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
