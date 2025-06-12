package goodachievement

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/good"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entgoodachievement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/goodachievement"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.GoodAchievementSelect
	infos    []*npool.Achievement
	total    uint32
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalizeString(amount string) string {
	_amount, err := decimal.NewFromString(amount)
	if err != nil {
		return decimal.NewFromInt(0).String()
	}
	return _amount.String()
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.SelfAmountUSD = h.formalizeString(info.SelfAmountUSD)
		info.TotalAmountUSD = h.formalizeString(info.TotalAmountUSD)
		info.SelfCommissionUSD = h.formalizeString(info.SelfCommissionUSD)
		info.TotalCommissionUSD = h.formalizeString(info.TotalCommissionUSD)
		info.SelfUnits = h.formalizeString(info.SelfUnits)
		info.TotalUnits = h.formalizeString(info.TotalUnits)
	}
}

func (h *Handler) GetAchievement(ctx context.Context) (*npool.Achievement, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodAchievement(cli); err != nil {
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

func (h *Handler) GetAchievements(ctx context.Context) ([]*npool.Achievement, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryGoodAchievements(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryGoodAchievements(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entgoodachievement.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
