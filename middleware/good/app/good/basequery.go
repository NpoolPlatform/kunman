package good

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodBaseSelect
}

func (h *baseQueryHandler) selectGood(stm *ent.AppGoodBaseQuery) *ent.AppGoodBaseSelect {
	return stm.Select(entgoodbase.FieldID)
}

func (h *baseQueryHandler) queryGood(cli *ent.Client) {
	h.stmSelect = h.selectGood(
		cli.AppGoodBase.
			Query().
			Where(
				entappgoodbase.EntID(*h.EntID),
				entappgoodbase.DeletedAt(0),
			),
	)
}

func (h *baseQueryHandler) queryGoods(cli *ent.Client) (*ent.AppGoodBaseSelect, error) {
	stm, err := goodbasecrud.SetQueryConds(cli.AppGoodBase.Query(), h.AppGoodBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectGood(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldID),
			t1.C(entappgoodbase.FieldID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldAppID),
			t1.C(entappgoodbase.FieldEntID),
			t1.C(entappgoodbase.FieldGoodID),
			sql.As(t1.C(entappgoodbase.FieldPurchasable), "app_good_purchasable"),
			t1.C(entappgoodbase.FieldEnableProductPage),
			t1.C(entappgoodbase.FieldProductPage),
			sql.As(t1.C(entappgoodbase.FieldOnline), "app_good_online"),
			t1.C(entappgoodbase.FieldVisible),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			t1.C(entappgoodbase.FieldDisplayIndex),
			t1.C(entappgoodbase.FieldBanner),
		)
}

func (h *baseQueryHandler) queryJoinGoodBase(s *sql.Selector) error {
	t1 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		)
	if h.GoodBaseConds.GoodType != nil {
		_type, ok := h.GoodBaseConds.GoodType.Val.(types.GoodType)
		if !ok {
			return wlog.Errorf("invalid goodtype")
		}
		s.OnP(
			sql.EQ(t1.C(entgoodbase.FieldGoodType), _type.String()),
		)
	}
	if h.GoodBaseConds.GoodTypes != nil {
		_types, ok := h.GoodBaseConds.GoodTypes.Val.([]types.GoodType)
		if !ok {
			return wlog.Errorf("invalid goodtypes")
		}
		s.OnP(
			sql.In(t1.C(entgoodbase.FieldGoodType), func() (__types []interface{}) {
				for _, _type := range _types {
					__types = append(__types, interface{}(_type.String()))
				}
				return
			}()...),
		)
	}
	s.AppendSelect(
		t1.C(entgoodbase.FieldGoodType),
		t1.C(entgoodbase.FieldBenefitType),
		sql.As(t1.C(entgoodbase.FieldName), "good_name"),
		t1.C(entgoodbase.FieldServiceStartAt),
		t1.C(entgoodbase.FieldStartMode),
		t1.C(entgoodbase.FieldState),
		t1.C(entgoodbase.FieldTestOnly),
		t1.C(entgoodbase.FieldBenefitIntervalHours),
		sql.As(t1.C(entgoodbase.FieldPurchasable), "good_purchasable"),
		sql.As(t1.C(entgoodbase.FieldOnline), "good_online"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinGoodBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodBase", "Error", err)
		}
	})
}
