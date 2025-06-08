package topmost

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

type queryHandler struct {
	*Handler
	topmosts []*topmostmwpb.TopMost
	infos    []*npool.TopMost
	apps     map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, topmost := range h.topmosts {
			appIDs = append(appIDs, topmost.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, topmost := range h.topmosts {
		info := &npool.TopMost{
			ID:          topmost.ID,
			EntID:       topmost.EntID,
			AppID:       topmost.AppID,
			TopMostType: topmost.TopMostType,
			Title:       topmost.Title,
			TargetUrl:   topmost.TargetUrl,
			Message:     topmost.Message,
			StartAt:     topmost.StartAt,
			EndAt:       topmost.EndAt,
			CreatedAt:   topmost.CreatedAt,
			UpdatedAt:   topmost.UpdatedAt,
		}

		app, ok := h.apps[topmost.AppID]
		if ok {
			info.AppName = app.Name
		}

		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetTopMost(ctx context.Context) (*npool.TopMost, error) {
	info, err := topmostmwcli.GetTopMost(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid topmost")
	}

	handler := &queryHandler{
		Handler:  h,
		topmosts: []*topmostmwpb.TopMost{info},
		apps:     map[string]*appmwpb.App{},
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

func (h *Handler) GetTopMosts(ctx context.Context) ([]*npool.TopMost, uint32, error) {
	infos, total, err := topmostmwcli.GetTopMosts(ctx, &topmostmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:  h,
		topmosts: infos,
		apps:     map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
