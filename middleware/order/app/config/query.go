package appconfig

import (
	"context"

	"entgo.io/ent/dialect/sql"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entappconfig "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/appconfig"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/app/config"
	appconfigcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/app/config"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppConfigSelect
	stmCount  *ent.AppConfigSelect
	infos     []*npool.AppConfig
	total     uint32
}

func (h *queryHandler) selectAppConfig(stm *ent.AppConfigQuery) *ent.AppConfigSelect {
	return stm.Select(entappconfig.FieldID)
}

func (h *queryHandler) queryAppConfig(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil && h.AppID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.AppConfig.Query().Where(entappconfig.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappconfig.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappconfig.EntID(*h.EntID))
	}
	if h.AppID != nil {
		stm.Where(entappconfig.AppID(*h.AppID))
	}
	h.stmSelect = h.selectAppConfig(stm)
	return nil
}

func (h *queryHandler) queryAppConfigs(cli *ent.Client) (*ent.AppConfigSelect, error) {
	stm, err := appconfigcrud.SetQueryConds(cli.AppConfig.Query(), h.AppConfigConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectAppConfig(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappconfig.Table)
	s.AppendSelect(
		t.C(entappconfig.FieldID),
		t.C(entappconfig.FieldEntID),
		t.C(entappconfig.FieldAppID),
		t.C(entappconfig.FieldEnableSimulateOrder),
		t.C(entappconfig.FieldSimulateOrderCouponMode),
		t.C(entappconfig.FieldSimulateOrderCouponProbability),
		t.C(entappconfig.FieldSimulateOrderCashableProfitProbability),
		t.C(entappconfig.FieldMaxUnpaidOrders),
		t.C(entappconfig.FieldMaxTypedCouponsPerOrder),
		t.C(entappconfig.FieldCreatedAt),
		t.C(entappconfig.FieldUpdatedAt),
	)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.SimulateOrderCouponProbability = func() string {
			amount, _ := decimal.NewFromString(info.SimulateOrderCouponProbability)
			return amount.String()
		}()
		info.SimulateOrderCashableProfitProbability = func() string {
			amount, _ := decimal.NewFromString(info.SimulateOrderCashableProfitProbability)
			return amount.String()
		}()
		info.SimulateOrderCouponMode = basetypes.SimulateOrderCouponMode(
			basetypes.SimulateOrderCouponMode_value[info.SimulateOrderCouponModeStr],
		)
	}
}

func (h *Handler) GetAppConfig(ctx context.Context) (*npool.AppConfig, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppConfig(cli); err != nil {
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

func (h *Handler) GetAppConfigs(ctx context.Context) (infos []*npool.AppConfig, total uint32, err error) {
	handler := &queryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryAppConfigs(cli); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stmCount, err = handler.queryAppConfigs(cli); err != nil {
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
			Order(ent.Desc(entappconfig.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
