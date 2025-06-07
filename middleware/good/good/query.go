package good

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.GoodBaseSelect
	infos    []*npool.Good
	total    uint32
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.BenefitType = types.BenefitType(types.BenefitType_value[info.BenefitTypeStr])
		info.StartMode = types.GoodStartMode(types.GoodStartMode_value[info.StartModeStr])
		info.State = types.GoodState(types.GoodState_value[info.StateStr])
	}
}

func (h *Handler) GetGood(ctx context.Context) (*npool.Good, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryGood(cli)
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
		return handler.scan(_ctx)
	}); err != nil {
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

func (h *Handler) GetGoods(ctx context.Context) (infos []*npool.Good, total uint32, err error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryGoods(cli); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stmCount, err = handler.queryGoods(cli); err != nil {
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
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	}); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
