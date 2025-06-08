package displaycolor

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	appgooddisplaycolormwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/color"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/display/color"
	appgooddisplaycolormwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/color"
)

type queryHandler struct {
	*Handler
	displayColors []*appgooddisplaycolormwpb.DisplayColor
	infos         []*npool.DisplayColor
	apps          map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, displayColor := range h.displayColors {
			appIDs = append(appIDs, displayColor.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, displayColor := range h.displayColors {
		info := &npool.DisplayColor{
			ID:          displayColor.ID,
			EntID:       displayColor.EntID,
			AppID:       displayColor.AppID,
			GoodID:      displayColor.GoodID,
			GoodName:    displayColor.GoodName,
			GoodType:    displayColor.GoodType,
			AppGoodID:   displayColor.AppGoodID,
			AppGoodName: displayColor.AppGoodName,
			Color:       displayColor.Color,
			Index:       displayColor.Index,
			CreatedAt:   displayColor.CreatedAt,
			UpdatedAt:   displayColor.UpdatedAt,
		}
		app, ok := h.apps[displayColor.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetDisplayColor(ctx context.Context) (*npool.DisplayColor, error) {
	displayColor, err := appgooddisplaycolormwcli.GetDisplayColor(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if displayColor == nil {
		return nil, fmt.Errorf("invalid displaycolor")
	}

	handler := &queryHandler{
		Handler:       h,
		displayColors: []*appgooddisplaycolormwpb.DisplayColor{displayColor},
		apps:          map[string]*appmwpb.App{},
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

func (h *Handler) GetDisplayColors(ctx context.Context) ([]*npool.DisplayColor, uint32, error) {
	conds := &appgooddisplaycolormwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	displayColors, total, err := appgooddisplaycolormwcli.GetDisplayColors(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(displayColors) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:       h,
		displayColors: displayColors,
		apps:          map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
