package config

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/good/commission/config"
	commissionconfigcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/app/good/commission/config"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcommissionconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appgoodcommissionconfig"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.AppGoodCommissionConfigSelect
	stmSelect *ent.AppGoodCommissionConfigSelect
	infos     []*npool.AppGoodCommissionConfig
	total     uint32
}

func (h *queryHandler) selectCommissionConfig(stm *ent.AppGoodCommissionConfigQuery) *ent.AppGoodCommissionConfigSelect {
	return stm.Select(
		entcommissionconfig.FieldID,
	)
}

func (h *queryHandler) queryCommissionConfig(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}

	stm := cli.AppGoodCommissionConfig.Query().Where(entcommissionconfig.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcommissionconfig.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcommissionconfig.EntID(*h.EntID))
	}
	h.stmSelect = h.selectCommissionConfig(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcommissionconfig.Table)
	s.LeftJoin(t).
		On(
			s.C(entcommissionconfig.FieldID),
			t.C(entcommissionconfig.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcommissionconfig.FieldEntID), "ent_id"),
			sql.As(t.C(entcommissionconfig.FieldAppID), "app_id"),
			sql.As(t.C(entcommissionconfig.FieldGoodID), "good_id"),
			sql.As(t.C(entcommissionconfig.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entcommissionconfig.FieldSettleType), "settle_type"),
			sql.As(t.C(entcommissionconfig.FieldAmountOrPercent), "amount_or_percent"),
			sql.As(t.C(entcommissionconfig.FieldThresholdAmount), "threshold_amount"),
			sql.As(t.C(entcommissionconfig.FieldStartAt), "start_at"),
			sql.As(t.C(entcommissionconfig.FieldEndAt), "end_at"),
			sql.As(t.C(entcommissionconfig.FieldInvites), "invites"),
			sql.As(t.C(entcommissionconfig.FieldDisabled), "disabled"),
			sql.As(t.C(entcommissionconfig.FieldLevel), "level"),
			sql.As(t.C(entcommissionconfig.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcommissionconfig.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryCommissionConfigs(cli *ent.Client) (*ent.AppGoodCommissionConfigSelect, error) {
	stm, err := commissionconfigcrud.SetQueryConds(cli.AppGoodCommissionConfig.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectCommissionConfig(stm), nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return wlog.WrapError(err)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.SettleType = types.SettleType(types.SettleType_value[info.SettleTypeStr])
		amount, err := decimal.NewFromString(info.AmountOrPercent)
		if err != nil {
			info.AmountOrPercent = decimal.NewFromInt(0).String()
		} else {
			info.AmountOrPercent = amount.String()
		}
		amount, err = decimal.NewFromString(info.ThresholdAmount)
		if err != nil {
			info.ThresholdAmount = decimal.NewFromInt(0).String()
		} else {
			info.ThresholdAmount = amount.String()
		}
	}
}

func (h *Handler) GetCommissionConfig(ctx context.Context) (*npool.AppGoodCommissionConfig, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.AppGoodCommissionConfig{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCommissionConfig(cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryJoin(); err != nil {
			return wlog.WrapError(err)
		}
		return handler.scan(ctx)
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

func (h *Handler) GetCommissionConfigs(ctx context.Context) ([]*npool.AppGoodCommissionConfig, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.AppGoodCommissionConfig{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCommissionConfigs(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryCommissionConfigs(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryJoin(); err != nil {
			return wlog.WrapError(err)
		}
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)
		handler.stmSelect.
			Order(ent.Asc(entcommissionconfig.FieldLevel)).
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
