package subscription

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/billing/db"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"
	npool "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/subscription"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	infos []*npool.Subscription
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, _ := decimal.NewFromString(info.UsdPrice)
		info.UsdPrice = amount.String()
		info.PackageType = types.PackageType(types.PackageType_value[info.PackageTypeStr])
		info.ResetType = types.ResetType(types.ResetType_value[info.ResetTypeStr])
	}
}

func (h *Handler) GetSubscription(ctx context.Context) (*npool.Subscription, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySubscription(cli); err != nil {
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

func (h *Handler) GetSubscriptions(ctx context.Context) ([]*npool.Subscription, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err = handler.querySubscriptions(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos, nil
}
