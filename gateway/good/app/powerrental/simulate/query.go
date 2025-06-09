package simulate

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	simulatemwcli "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental/simulate"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"
	simulatemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental/simulate"
)

type queryHandler struct {
	*Handler
	simulates []*simulatemwpb.Simulate
	infos     []*npool.Simulate
	apps      map[string]*appmwpb.App
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, simulate := range h.simulates {
			appIDs = append(appIDs, simulate.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, simulate := range h.simulates {
		info := &npool.Simulate{
			ID:                   simulate.ID,
			EntID:                simulate.EntID,
			AppID:                simulate.AppID,
			GoodID:               simulate.GoodID,
			GoodName:             simulate.GoodName,
			AppGoodID:            simulate.AppGoodID,
			AppGoodName:          simulate.AppGoodName,
			OrderUnits:           simulate.OrderUnits,
			OrderDurationSeconds: simulate.OrderDurationSeconds,
			GoodCoins:            simulate.GoodCoins,
			CreatedAt:            simulate.CreatedAt,
			UpdatedAt:            simulate.UpdatedAt,
		}

		app, ok := h.apps[simulate.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetSimulate(ctx context.Context) (*npool.Simulate, error) {
	info, err := simulatemwcli.GetSimulate(ctx, *h.AppGoodID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid simulate")
	}

	handler := &queryHandler{
		Handler:   h,
		simulates: []*simulatemwpb.Simulate{info},
		apps:      map[string]*appmwpb.App{},
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

func (h *Handler) GetSimulates(ctx context.Context) ([]*npool.Simulate, uint32, error) {
	infos, total, err := simulatemwcli.GetSimulates(ctx, &simulatemwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:   h,
		simulates: infos,
		apps:      map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
