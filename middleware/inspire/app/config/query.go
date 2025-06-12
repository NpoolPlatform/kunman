package config

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"
	appconfigcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/app/config"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entappconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appconfig"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.AppConfigSelect
	stmSelect *ent.AppConfigSelect
	infos     []*npool.AppConfig
	total     uint32
}

func (h *queryHandler) selectAppConfig(stm *ent.AppConfigQuery) *ent.AppConfigSelect {
	return stm.Select(
		entappconfig.FieldID,
	)
}

func (h *queryHandler) queryAppConfig(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}

	stm := cli.AppConfig.Query().Where(entappconfig.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappconfig.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappconfig.EntID(*h.EntID))
	}
	h.stmSelect = h.selectAppConfig(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappconfig.Table)
	s.LeftJoin(t).
		On(
			s.C(entappconfig.FieldID),
			t.C(entappconfig.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappconfig.FieldEntID), "ent_id"),
			sql.As(t.C(entappconfig.FieldAppID), "app_id"),
			sql.As(t.C(entappconfig.FieldSettleMode), "settle_mode"),
			sql.As(t.C(entappconfig.FieldSettleAmountType), "settle_amount_type"),
			sql.As(t.C(entappconfig.FieldSettleInterval), "settle_interval"),
			sql.As(t.C(entappconfig.FieldCommissionType), "commission_type"),
			sql.As(t.C(entappconfig.FieldSettleBenefit), "settle_benefit"),
			sql.As(t.C(entappconfig.FieldStartAt), "start_at"),
			sql.As(t.C(entappconfig.FieldEndAt), "end_at"),
			sql.As(t.C(entappconfig.FieldMaxLevel), "max_level"),
			sql.As(t.C(entappconfig.FieldCreatedAt), "created_at"),
			sql.As(t.C(entappconfig.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryAppConfigs(cli *ent.Client) (*ent.AppConfigSelect, error) {
	stm, err := appconfigcrud.SetQueryConds(cli.AppConfig.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectAppConfig(stm), nil
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
		info.CommissionType = types.CommissionType(types.CommissionType_value[info.CommissionTypeStr])
		info.SettleMode = types.SettleMode(types.SettleMode_value[info.SettleModeStr])
		info.SettleAmountType = types.SettleAmountType(types.SettleAmountType_value[info.SettleAmountTypeStr])
		info.SettleInterval = types.SettleInterval(types.SettleInterval_value[info.SettleIntervalStr])
	}
}

func (h *Handler) GetAppConfig(ctx context.Context) (*npool.AppConfig, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.AppConfig{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppConfig(cli); err != nil {
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

func (h *Handler) GetAppConfigs(ctx context.Context) ([]*npool.AppConfig, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.AppConfig{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryAppConfigs(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryAppConfigs(cli)
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
