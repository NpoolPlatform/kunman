package description

import (
	"context"
	"fmt"

	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/description"
	appgooddescriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/description"
	appgooddescriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/description"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	descriptions []*appgooddescriptionmwpb.Description
	infos        []*npool.Description
	apps         map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, description := range h.descriptions {
			appIDs = append(appIDs, description.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, description := range h.descriptions {
		info := &npool.Description{
			ID:          description.ID,
			EntID:       description.EntID,
			AppID:       description.AppID,
			GoodID:      description.GoodID,
			GoodName:    description.GoodName,
			GoodType:    description.GoodType,
			AppGoodID:   description.AppGoodID,
			AppGoodName: description.AppGoodName,
			Description: description.Description,
			Index:       description.Index,
			CreatedAt:   description.CreatedAt,
			UpdatedAt:   description.UpdatedAt,
		}
		app, ok := h.apps[description.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetDescription(ctx context.Context) (*npool.Description, error) {
	descriptionHandler, err := appgooddescriptionmw.NewHandler(
		ctx,
		appgooddescriptionmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	description, err := descriptionHandler.GetDescription(ctx)
	if err != nil {
		return nil, err
	}
	if description == nil {
		return nil, fmt.Errorf("invalid description")
	}

	handler := &queryHandler{
		Handler:      h,
		descriptions: []*appgooddescriptionmwpb.Description{description},
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

func (h *Handler) GetDescriptions(ctx context.Context) ([]*npool.Description, uint32, error) {
	conds := &appgooddescriptionmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}

	descriptionHandler, err := appgooddescriptionmw.NewHandler(
		ctx,
		appgooddescriptionmw.WithConds(conds),
		appgooddescriptionmw.WithOffset(h.Offset),
		appgooddescriptionmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	descriptions, total, err := descriptionHandler.GetDescriptions(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(descriptions) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:      h,
		descriptions: descriptions,
		apps:         map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
