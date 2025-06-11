package malfunction

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
	malfunctionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/malfunction"
	malfunctionmw "github.com/NpoolPlatform/kunman/middleware/good/good/malfunction"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	malfunctions           []*malfunctionmwpb.Malfunction
	compensateOrderNumbers map[string]uint32
	infos                  []*npool.Malfunction
}

func (h *queryHandler) getCompensateOrderNumbers(ctx context.Context) (err error) {
	h.compensateOrderNumbers, err = goodgwcommon.GetCompensateOrderNumbers(ctx, func() (compensateFromIDs []string) {
		for _, malfunction := range h.malfunctions {
			compensateFromIDs = append(compensateFromIDs, malfunction.EntID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) formalize() {
	for _, info := range h.malfunctions {
		h.infos = append(h.infos, &npool.Malfunction{
			ID:                info.ID,
			EntID:             info.EntID,
			GoodID:            info.GoodID,
			GoodType:          info.GoodType,
			GoodName:          info.GoodName,
			Title:             info.Title,
			Message:           info.Message,
			StartAt:           info.StartAt,
			DurationSeconds:   info.DurationSeconds,
			CompensateSeconds: info.CompensateSeconds,
			CompensatedOrders: h.compensateOrderNumbers[info.EntID],
			CreatedAt:         info.CreatedAt,
			UpdatedAt:         info.UpdatedAt,
		})
	}
}

func (h *Handler) GetMalfunction(ctx context.Context) (*npool.Malfunction, error) {
	malfunctionHandler, err := malfunctionmw.NewHandler(
		ctx,
		malfunctionmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := malfunctionHandler.GetMalfunction(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid malfunction")
	}

	handler := &queryHandler{
		malfunctions: []*malfunctionmwpb.Malfunction{info},
	}
	if err := handler.getCompensateOrderNumbers(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetMalfunctions(ctx context.Context) ([]*npool.Malfunction, uint32, error) {
	conds := &malfunctionmwpb.Conds{}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}

	malfunctionHandler, err := malfunctionmw.NewHandler(
		ctx,
		malfunctionmw.WithConds(conds),
		malfunctionmw.WithOffset(h.Offset),
		malfunctionmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	infos, total, err := malfunctionHandler.GetMalfunctions(ctx)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		malfunctions: infos,
	}

	if err := handler.getCompensateOrderNumbers(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, total, nil
}
