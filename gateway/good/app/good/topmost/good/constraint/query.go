package constraint

import (
	"context"
	"fmt"

	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/constraint"
	constraintmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good/constraint"
	constraintmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good/constraint"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	constraints []*constraintmwpb.TopMostGoodConstraint
	infos       []*npool.TopMostGoodConstraint
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
		info := &npool.TopMostGoodConstraint{
			ID:               constraint.ID,
			EntID:            constraint.EntID,
			AppID:            constraint.AppID,
			TopMostID:        constraint.TopMostID,
			TopMostType:      constraint.TopMostType,
			TopMostTitle:     constraint.TopMostTitle,
			TopMostMessage:   constraint.TopMostMessage,
			TopMostTargetUrl: constraint.TopMostTargetUrl,
			Constraint:       constraint.Constraint,
			TopMostGoodID:    constraint.TopMostGoodID,
			AppGoodID:        constraint.AppGoodID,
			AppGoodName:      constraint.AppGoodName,
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

func (h *Handler) GetConstraint(ctx context.Context) (*npool.TopMostGoodConstraint, error) {
	constraintHandler, err := constraintmw.NewHandler(
		ctx,
		constraintmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := constraintHandler.GetConstraint(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid constraint")
	}

	handler := &queryHandler{
		Handler:     h,
		constraints: []*constraintmwpb.TopMostGoodConstraint{info},
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

func (h *Handler) GetConstraints(ctx context.Context) ([]*npool.TopMostGoodConstraint, uint32, error) {
	constraintHandler, err := constraintmw.NewHandler(
		ctx,
		constraintmw.WithConds(
			&constraintmwpb.Conds{
				AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
			},
		),
		constraintmw.WithOffset(h.Offset),
		constraintmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	infos, total, err := constraintHandler.GetConstraints(ctx)
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
