package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription/oneshot"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	infos []*npool.OneShot
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, _ := decimal.NewFromString(info.USDPrice)
		info.USDPrice = amount.String()
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
	}
}

func (h *Handler) GetOneShot(ctx context.Context) (*npool.OneShot, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodBase(cli); err != nil {
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

func (h *Handler) GetOneShots(ctx context.Context) ([]*npool.OneShot, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err = handler.queryGoodBases(cli); err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, nil
}
