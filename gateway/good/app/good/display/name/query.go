package displayname

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	appgooddisplaynamemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/name"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/display/name"
	appgooddisplaynamemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/name"
)

type queryHandler struct {
	*Handler
	displayNames []*appgooddisplaynamemwpb.DisplayName
	infos        []*npool.DisplayName
	apps         map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, displayName := range h.displayNames {
			appIDs = append(appIDs, displayName.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, displayName := range h.displayNames {
		info := &npool.DisplayName{
			ID:          displayName.ID,
			EntID:       displayName.EntID,
			AppID:       displayName.AppID,
			GoodID:      displayName.GoodID,
			GoodName:    displayName.GoodName,
			GoodType:    displayName.GoodType,
			AppGoodID:   displayName.AppGoodID,
			AppGoodName: displayName.AppGoodName,
			Name:        displayName.Name,
			Index:       displayName.Index,
			CreatedAt:   displayName.CreatedAt,
			UpdatedAt:   displayName.UpdatedAt,
		}
		app, ok := h.apps[displayName.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetDisplayName(ctx context.Context) (*npool.DisplayName, error) {
	displayName, err := appgooddisplaynamemwcli.GetDisplayName(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if displayName == nil {
		return nil, fmt.Errorf("invalid displayname")
	}

	handler := &queryHandler{
		Handler:      h,
		displayNames: []*appgooddisplaynamemwpb.DisplayName{displayName},
		apps:         map[string]*appmwpb.App{},
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

func (h *Handler) GetDisplayNames(ctx context.Context) ([]*npool.DisplayName, uint32, error) {
	conds := &appgooddisplaynamemwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	displayNames, total, err := appgooddisplaynamemwcli.GetDisplayNames(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(displayNames) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:      h,
		displayNames: displayNames,
		apps:         map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
