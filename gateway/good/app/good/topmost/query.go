package topmost

import (
	"context"
	"fmt"

	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"
	topmostmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost"
	topmostmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
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
	topMostHandler, err := topmostmw.NewHandler(
		ctx,
		topmostmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := topMostHandler.GetTopMost(ctx)
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
	conds := &topmostmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	topMostHandler, err := topmostmw.NewHandler(
		ctx,
		topmostmw.WithConds(conds),
		topmostmw.WithOffset(h.Offset),
		topmostmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	infos, total, err := topMostHandler.GetTopMosts(ctx)
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
