package constraint

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	constraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/constraint"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/constraint"
	constraintmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/constraint"
)

type queryHandler struct {
	*Handler
	constraints []*constraintmwpb.TopMostConstraint
	infos       []*npool.TopMostConstraint
	apps        map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, constraint := range h.constraints {
			appIDs = append(appIDs, constraint.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, constraint := range h.constraints {
		info := &npool.TopMostConstraint{
			ID:               constraint.ID,
			EntID:            constraint.EntID,
			AppID:            constraint.AppID,
			TopMostID:        constraint.TopMostID,
			TopMostType:      constraint.TopMostType,
			TopMostTitle:     constraint.TopMostTitle,
			TopMostMessage:   constraint.TopMostMessage,
			TopMostTargetUrl: constraint.TopMostTargetUrl,
			Constraint:       constraint.Constraint,
			TargetValue:      constraint.TargetValue,
			Index:            constraint.Index,
			CreatedAt:        constraint.CreatedAt,
			UpdatedAt:        constraint.UpdatedAt,
		}

		app, ok := h.apps[constraint.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetConstraint(ctx context.Context) (*npool.TopMostConstraint, error) {
	info, err := constraintmwcli.GetTopMostConstraint(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid constraint")
	}

	handler := &queryHandler{
		Handler:     h,
		constraints: []*constraintmwpb.TopMostConstraint{info},
		apps:        map[string]*appmwpb.App{},
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

func (h *Handler) GetConstraints(ctx context.Context) ([]*npool.TopMostConstraint, uint32, error) {
	infos, total, err := constraintmwcli.GetTopMostConstraints(
		ctx,
		&constraintmwpb.Conds{
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
		Handler:     h,
		constraints: infos,
		apps:        map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
