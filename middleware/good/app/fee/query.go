package fee

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.AppGoodBaseSelect
	infos    []*npool.Fee
	total    uint32
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppFee(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppFee", "Error", err)
			return
		}
		if err := h.queryJoinGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinGood", "Error", err)
			return
		}
		if err := h.queryJoinFee(s); err != nil {
			logger.Sugar().Errorw("queryJoinFee", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, _ := decimal.NewFromString(info.UnitValue)
		info.UnitValue = amount.String()
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.CancelMode = types.CancelMode(types.CancelMode_value[info.CancelModeStr])
		info.SettlementType = types.GoodSettlementType(types.GoodSettlementType_value[info.SettlementTypeStr])
		info.DurationDisplayType = types.GoodDurationType(types.GoodDurationType_value[info.DurationDisplayTypeStr])
	}
}

func (h *Handler) GetFee(ctx context.Context) (*npool.Fee, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppGoodBase(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetFees(ctx context.Context) ([]*npool.Fee, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryAppGoodBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryAppGoodBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
