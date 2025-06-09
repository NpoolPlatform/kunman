package label

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	appgoodlabelmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/label"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/label"
	appgoodlabelmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/label"
)

type queryHandler struct {
	*Handler
	labels []*appgoodlabelmwpb.Label
	infos  []*npool.Label
	apps   map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, label := range h.labels {
			appIDs = append(appIDs, label.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, label := range h.labels {
		info := &npool.Label{
			ID:           label.ID,
			EntID:        label.EntID,
			AppID:        label.AppID,
			GoodID:       label.GoodID,
			GoodName:     label.GoodName,
			GoodType:     label.GoodType,
			AppGoodID:    label.AppGoodID,
			AppGoodName:  label.AppGoodName,
			Icon:         label.Icon,
			IconBgColor:  label.IconBgColor,
			Label:        label.Label,
			LabelBgColor: label.LabelBgColor,
			Index:        label.Index,
			CreatedAt:    label.CreatedAt,
			UpdatedAt:    label.UpdatedAt,
		}
		app, ok := h.apps[label.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetLabel(ctx context.Context) (*npool.Label, error) {
	label, err := appgoodlabelmwcli.GetLabel(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if label == nil {
		return nil, fmt.Errorf("invalid label")
	}

	handler := &queryHandler{
		Handler: h,
		labels:  []*appgoodlabelmwpb.Label{label},
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

func (h *Handler) GetLabels(ctx context.Context) ([]*npool.Label, uint32, error) {
	conds := &appgoodlabelmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	labels, total, err := appgoodlabelmwcli.GetLabels(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(labels) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler: h,
		labels:  labels,
		apps:    map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
