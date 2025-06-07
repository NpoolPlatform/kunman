package delegatedstaking

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entdelegatedstaking "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/delegatedstaking"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entgoodcoin "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoin"
	entgoodreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodreward"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodBaseSelect
}

func (h *baseQueryHandler) selectGoodBase(stm *ent.GoodBaseQuery) *ent.GoodBaseSelect {
	return stm.Select(entgoodbase.FieldID)
}

func (h *baseQueryHandler) queryGoodBase(cli *ent.Client) error {
	if h.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}
	h.stmSelect = h.selectGoodBase(
		cli.GoodBase.
			Query().
			Where(
				entgoodbase.DeletedAt(0),
				entgoodbase.EntID(*h.GoodID),
				entgoodbase.Or(
					entgoodbase.GoodType(types.GoodType_DelegatedStaking.String()),
				),
			),
	)
	return nil
}

func (h *baseQueryHandler) queryGoodBases(cli *ent.Client) (*ent.GoodBaseSelect, error) {
	stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	stm.Where(
		entgoodbase.Or(
			entgoodbase.GoodType(types.GoodType_DelegatedStaking.String()),
		),
	)
	return h.selectGoodBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodbase.FieldID),
			t1.C(entgoodbase.FieldID),
		).
		AppendSelect(
			t1.C(entgoodbase.FieldGoodType),
			t1.C(entgoodbase.FieldBenefitType),
			t1.C(entgoodbase.FieldName),
			t1.C(entgoodbase.FieldServiceStartAt),
			t1.C(entgoodbase.FieldStartMode),
			t1.C(entgoodbase.FieldTestOnly),
			t1.C(entgoodbase.FieldBenefitIntervalHours),
			t1.C(entgoodbase.FieldPurchasable),
			t1.C(entgoodbase.FieldOnline),
			t1.C(entgoodbase.FieldState),
			t1.C(entgoodbase.FieldCreatedAt),
		)
}

//nolint:gocyclo
func (h *baseQueryHandler) queryJoinDelegatedStaking(s *sql.Selector) error {
	t1 := sql.Table(entdelegatedstaking.Table)

	s.Join(t1).
		On(
			s.C(entgoodbase.FieldEntID),
			t1.C(entdelegatedstaking.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entdelegatedstaking.FieldDeletedAt), 0),
		)
	if h.DelegatedStakingConds.ID != nil {
		u, ok := h.DelegatedStakingConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entdelegatedstaking.FieldID), u))
	}
	if h.DelegatedStakingConds.IDs != nil {
		ids, ok := h.DelegatedStakingConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(sql.In(t1.C(entdelegatedstaking.FieldID), func() (_ids []interface{}) {
			for _, id := range ids {
				_ids = append(_ids, interface{}(id))
			}
			return
		}()...))
	}
	if h.DelegatedStakingConds.EntID != nil {
		uid, ok := h.DelegatedStakingConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entdelegatedstaking.FieldEntID), uid))
	}
	if h.DelegatedStakingConds.EntIDs != nil {
		uids, ok := h.DelegatedStakingConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(sql.In(t1.C(entdelegatedstaking.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return
		}()...))
	}
	if h.DelegatedStakingConds.ContractState != nil {
		state, ok := h.DelegatedStakingConds.ContractState.Val.(types.ContractState)
		if !ok {
			return wlog.Errorf("invalid contractstate")
		}
		s.OnP(
			sql.EQ(t1.C(entdelegatedstaking.FieldContractState), state.String()),
		)
	}
	if h.DelegatedStakingConds.ContractStates != nil {
		states, ok := h.DelegatedStakingConds.ContractStates.Val.([]types.ContractState)
		if !ok {
			return wlog.Errorf("invalid contractstates")
		}
		s.OnP(
			sql.In(t1.C(entdelegatedstaking.FieldContractState), func() (_states []interface{}) {
				for _, state := range states {
					_states = append(_states, interface{}(state.String()))
				}
				return
			}()...),
		)
	}

	s.AppendSelect(
		t1.C(entdelegatedstaking.FieldID),
		t1.C(entdelegatedstaking.FieldEntID),
		t1.C(entdelegatedstaking.FieldGoodID),
		t1.C(entdelegatedstaking.FieldContractCodeURL),
		t1.C(entdelegatedstaking.FieldContractCodeBranch),
		t1.C(entdelegatedstaking.FieldContractState),
		sql.As(t1.C(entdelegatedstaking.FieldUpdatedAt), "updated_at"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinReward(s *sql.Selector) {
	t := sql.Table(entgoodreward.Table)
	s.Join(t).
		On(
			s.C(entgoodbase.FieldEntID),
			t.C(entgoodreward.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entgoodreward.FieldDeletedAt), 0),
		)
	if h.RewardConds.RewardState != nil {
		s.OnP(
			sql.EQ(t.C(entgoodreward.FieldRewardState), h.RewardConds.RewardState.Val.(types.BenefitState).String()),
		)
	}
	if h.RewardConds.RewardAt != nil {
		switch h.RewardConds.RewardAt.Op {
		case cruder.EQ:
			s.OnP(sql.EQ(t.C(entgoodreward.FieldLastRewardAt), h.RewardConds.RewardAt.Val))
		case cruder.NEQ:
			s.OnP(sql.NEQ(t.C(entgoodreward.FieldLastRewardAt), h.RewardConds.RewardAt.Val))
		}
	}
	s.AppendSelect(
		t.C(entgoodreward.FieldRewardState),
		t.C(entgoodreward.FieldLastRewardAt),
	)
}

func (h *baseQueryHandler) queryJoinGoodCoin(s *sql.Selector) error {
	t := sql.Table(entgoodcoin.Table)
	s.LeftJoin(t).
		On(
			s.C(entgoodbase.FieldEntID),
			t.C(entgoodcoin.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entgoodcoin.FieldDeletedAt), 0),
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

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinReward(s)
		if err := h.queryJoinGoodCoin(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodCoin", "Error", err)
		}
		if err := h.queryJoinDelegatedStaking(s); err != nil {
			logger.Sugar().Errorw("queryJoinDelegatedStaking", "Error", err)
		}
	})
}
