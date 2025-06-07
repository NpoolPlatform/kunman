package delegatedstaking

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappdelegatedstaking "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appdelegatedstaking"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entdelegatedstaking "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/delegatedstaking"
	entextrainfo "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/extrainfo"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entgoodcoin "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoin"
	entgoodreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodreward"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodBaseSelect
}

func (h *baseQueryHandler) selectAppGoodBase(stm *ent.AppGoodBaseQuery) *ent.AppGoodBaseSelect {
	return stm.Select(entappgoodbase.FieldID)
}

func (h *baseQueryHandler) queryAppGoodBase(cli *ent.Client) error {
	if h.AppGoodID == nil {
		return wlog.Errorf("invalid appgoodid")
	}
	h.stmSelect = h.selectAppGoodBase(
		cli.AppGoodBase.
			Query().
			Where(
				entappgoodbase.DeletedAt(0),
				entappgoodbase.EntID(*h.AppGoodID),
			),
	)
	return nil
}

func (h *baseQueryHandler) queryAppGoodBases(cli *ent.Client) (*ent.AppGoodBaseSelect, error) {
	stm, err := appgoodbasecrud.SetQueryConds(cli.AppGoodBase.Query(), h.AppGoodBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectAppGoodBase(stm), nil
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
			t1.C(entappgoodbase.FieldGoodID),
			sql.As(t1.C(entappgoodbase.FieldPurchasable), "app_good_purchasable"),
			sql.As(t1.C(entappgoodbase.FieldOnline), "app_good_online"),
			t1.C(entappgoodbase.FieldEnableProductPage),
			t1.C(entappgoodbase.FieldProductPage),
			t1.C(entappgoodbase.FieldVisible),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			t1.C(entappgoodbase.FieldDisplayIndex),
			t1.C(entappgoodbase.FieldBanner),
			t1.C(entappgoodbase.FieldCreatedAt),
			t1.C(entappgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinGoodBase(s *sql.Selector) {
	t1 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		).
		OnP(
			sql.Or(
				sql.EQ(t1.C(entgoodbase.FieldGoodType), types.GoodType_DelegatedStaking.String()),
			),
		).
		AppendSelect(
			t1.C(entgoodbase.FieldGoodType),
			t1.C(entgoodbase.FieldBenefitType),
			sql.As(t1.C(entgoodbase.FieldName), "good_name"),
			sql.As(t1.C(entgoodbase.FieldServiceStartAt), "good_service_start_at"),
			sql.As(t1.C(entgoodbase.FieldStartMode), "good_start_mode"),
			t1.C(entgoodbase.FieldTestOnly),
			t1.C(entgoodbase.FieldBenefitIntervalHours),
			sql.As(t1.C(entgoodbase.FieldPurchasable), "good_purchasable"),
			sql.As(t1.C(entgoodbase.FieldOnline), "good_online"),
			t1.C(entgoodbase.FieldState),
		)
}

func (h *baseQueryHandler) queryJoinGoodCoin(s *sql.Selector) error {
	t := sql.Table(entgoodcoin.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t.C(entgoodcoin.FieldGoodID),
		).
		Distinct()
	if h.GoodCoinConds.CoinTypeID != nil {
		id, ok := h.GoodCoinConds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid cointypeid")
		}
		s.OnP(
			sql.EQ(t.C(entgoodcoin.FieldCoinTypeID), id),
		)
		s.Where(
			sql.EQ(t.C(entgoodcoin.FieldCoinTypeID), id),
		)
	}
	if h.GoodCoinConds.CoinTypeIDs != nil {
		uids, ok := h.GoodCoinConds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid cointypeids")
		}
		_uids := func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return
		}()
		s.OnP(
			sql.In(t.C(entgoodcoin.FieldCoinTypeID), _uids...),
		)
		s.Where(
			sql.In(t.C(entgoodcoin.FieldCoinTypeID), _uids...),
		)
	}
	return nil
}

func (h *baseQueryHandler) queryJoinGoodReward(s *sql.Selector) {
	t := sql.Table(entgoodreward.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t.C(entgoodreward.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entextrainfo.FieldDeletedAt), 0),
		).
		AppendSelect(
			t.C(entgoodreward.FieldLastRewardAt),
		)
}

func (h *baseQueryHandler) queryJoinExtraInfo(s *sql.Selector) {
	t := sql.Table(entextrainfo.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodbase.FieldEntID),
			t.C(entextrainfo.FieldAppGoodID),
		).
		OnP(
			sql.EQ(t.C(entextrainfo.FieldDeletedAt), 0),
		).
		AppendSelect(
			t.C(entextrainfo.FieldLikes),
			t.C(entextrainfo.FieldDislikes),
			t.C(entextrainfo.FieldScoreCount),
			t.C(entextrainfo.FieldRecommendCount),
			t.C(entextrainfo.FieldCommentCount),
			t.C(entextrainfo.FieldScore),
		)
}

func (h *baseQueryHandler) queryJoinAppDelegatedStaking(s *sql.Selector) error {
	t1 := sql.Table(entappdelegatedstaking.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entappdelegatedstaking.FieldAppGoodID),
		)
	if h.ID != nil {
		s.OnP(
			sql.EQ(t1.C(entappdelegatedstaking.FieldID), *h.ID),
		)
	}
	if h.AppDelegatedStakingConds.ID != nil {
		id, ok := h.AppDelegatedStakingConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(
			sql.EQ(t1.C(entappdelegatedstaking.FieldID), id),
		)
	}
	if h.EntID != nil {
		s.OnP(
			sql.EQ(t1.C(entappdelegatedstaking.FieldEntID), *h.EntID),
		)
	}
	if h.AppDelegatedStakingConds.EntID != nil {
		uid, ok := h.AppDelegatedStakingConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(
			sql.EQ(t1.C(entappdelegatedstaking.FieldEntID), uid),
		)
	}

	s.AppendSelect(
		t1.C(entappdelegatedstaking.FieldID),
		t1.C(entappdelegatedstaking.FieldEntID),
		t1.C(entappdelegatedstaking.FieldAppGoodID),
		t1.C(entappdelegatedstaking.FieldEnableSetCommission),
		sql.As(t1.C(entappdelegatedstaking.FieldServiceStartAt), "app_good_service_start_at"),
		sql.As(t1.C(entappdelegatedstaking.FieldStartMode), "app_good_start_mode"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinDelegatedStaking(s *sql.Selector) error {
	t1 := sql.Table(entdelegatedstaking.Table)

	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entdelegatedstaking.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entdelegatedstaking.FieldDeletedAt), 0),
		)

	if h.DelegatedStakingConds.GoodID != nil {
		uid, ok := h.DelegatedStakingConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t1.C(entdelegatedstaking.FieldGoodID), uid))
	}
	if h.DelegatedStakingConds.GoodIDs != nil {
		uids, ok := h.DelegatedStakingConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(sql.In(t1.C(entdelegatedstaking.FieldGoodID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return
		}()...))
	}
	s.AppendSelect(
		t1.C(entdelegatedstaking.FieldGoodID),
		sql.As(t1.C(entdelegatedstaking.FieldContractCodeURL), "contract_code_url"),
		sql.As(t1.C(entdelegatedstaking.FieldContractCodeBranch), "contract_code_branch"),
		sql.As(t1.C(entdelegatedstaking.FieldContractState), "contract_state"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinGoodBase(s)
		h.queryJoinGoodReward(s)
		h.queryJoinExtraInfo(s)
		if err := h.queryJoinDelegatedStaking(s); err != nil {
			logger.Sugar().Errorw("queryJoinDelegatedStaking", "Error", err)
		}
		if err := h.queryJoinAppDelegatedStaking(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppDelegatedStaking", "Error", err)
		}

		if err := h.queryJoinGoodCoin(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodCoin", "Error", err)
		}
	})
}
