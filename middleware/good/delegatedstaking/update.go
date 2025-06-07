//nolint:dupl
package delegatedstaking

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	goodcoinrewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin/reward"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	goodcoinreward1 "github.com/NpoolPlatform/kunman/middleware/good/good/coin/reward"
	rewardhistory1 "github.com/NpoolPlatform/kunman/middleware/good/good/coin/reward/history"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/good/goodbase"
	goodreward1 "github.com/NpoolPlatform/kunman/middleware/good/good/reward"
	goodstm "github.com/NpoolPlatform/kunman/middleware/good/stm"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*delegatedstakingGoodQueryHandler
	sqlDelegatedStaking    string
	sqlGoodReward          string
	sqlGoodBase            string
	sqlCoinRewards         []string
	sqlCoinRewardHistories []string
}

func (h *updateHandler) constructGoodRewardSQL(ctx context.Context) error {
	handler, err := goodreward1.NewHandler(
		ctx,
		goodreward1.WithGoodID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), true),
		goodreward1.WithRewardState(h.RewardReq.RewardState, false),
		goodreward1.WithLastRewardAt(h.RewardReq.LastRewardAt, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlGoodReward, err = handler.ConstructUpdateSQL()
	if err != nil && !wlog.Equal(err, cruder.ErrUpdateNothing) {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) constructGoodBaseSQL(ctx context.Context) error {
	handler, err := goodbase1.NewHandler(
		ctx,
		goodbase1.WithEntID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), true),
		goodbase1.WithGoodType(h.GoodBaseReq.GoodType, false),
		goodbase1.WithBenefitType(h.GoodBaseReq.BenefitType, false),
		goodbase1.WithName(h.GoodBaseReq.Name, false),
		goodbase1.WithServiceStartAt(h.GoodBaseReq.ServiceStartAt, false),
		goodbase1.WithStartMode(h.GoodBaseReq.StartMode, false),
		goodbase1.WithTestOnly(h.GoodBaseReq.TestOnly, false),
		goodbase1.WithBenefitIntervalHours(h.GoodBaseReq.BenefitIntervalHours, false),
		goodbase1.WithPurchasable(h.GoodBaseReq.Purchasable, false),
		goodbase1.WithOnline(h.GoodBaseReq.Online, false),
		goodbase1.WithState(h.GoodBaseReq.State, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlGoodBase, err = handler.ConstructUpdateSQL()
	if err != nil && !wlog.Equal(err, cruder.ErrUpdateNothing) {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) constructCoinRewardSQLs(ctx context.Context) error {
	if h.RewardReq.RewardState == nil {
		return nil
	}

	updateTotal := *h.RewardReq.RewardState == types.BenefitState_BenefitDone
	checkExist := *h.RewardReq.RewardState != types.BenefitState_BenefitWait

	for _, reward := range h.CoinRewardReqs {
		if updateTotal && reward.LastRewardAmount == nil {
			return wlog.Errorf("invalid lastrewardamount")
		}
		handler, err := goodcoinreward1.NewHandler(
			ctx,
			goodcoinreward1.WithGoodID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), true),
			goodcoinreward1.WithCoinTypeID(func() *string { s := reward.CoinTypeID.String(); return &s }(), true),
			goodcoinreward1.WithRewardTID(func() *string {
				if reward.RewardTID == nil {
					return nil
				}
				s := reward.RewardTID.String()
				return &s
			}(), false),
			goodcoinreward1.WithNextRewardStartAmount(func() *string {
				if reward.NextRewardStartAmount == nil {
					return nil
				}
				s := reward.NextRewardStartAmount.String()
				return &s
			}(), false),
			goodcoinreward1.WithRewardAmount(func() *string {
				if reward.LastRewardAmount == nil {
					return nil
				}
				s := reward.LastRewardAmount.String()
				return &s
			}(), false),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		sql, err := handler.ConstructUpdateSQL(updateTotal, *h.RewardReq.LastRewardAt, checkExist)
		if err != nil && !wlog.Equal(err, cruder.ErrUpdateNothing) {
			return wlog.WrapError(err)
		}
		h.sqlCoinRewards = append(h.sqlCoinRewards, sql)
	}
	return nil
}

func (h *updateHandler) constructCoinRewardHistorySQLs(ctx context.Context) error {
	for _, reward := range h.CoinRewardReqs {
		handler, err := rewardhistory1.NewHandler(
			ctx,
			rewardhistory1.WithGoodID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), true),
			rewardhistory1.WithCoinTypeID(func() *string { s := reward.CoinTypeID.String(); return &s }(), true),
			rewardhistory1.WithTID(func() *string {
				if reward.RewardTID == nil {
					return nil
				}
				s := reward.RewardTID.String()
				return &s
			}(), false),
			rewardhistory1.WithRewardDate(h.RewardReq.LastRewardAt, true),
			rewardhistory1.WithAmount(func() *string {
				if reward.LastRewardAmount == nil {
					return nil
				}
				s := reward.LastRewardAmount.String()
				return &s
			}(), false),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.sqlCoinRewardHistories = append(h.sqlCoinRewardHistories, handler.ConstructCreateSQL())
	}
	return nil
}

func (h *updateHandler) constructDelegatedStakingSQL() {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update delegated_stakings "
	if h.ContractCodeURL != nil {
		_sql += fmt.Sprintf("%vcontract_code_url = '%v', ", set, *h.ContractCodeURL)
		set = ""
	}
	if h.ContractCodeBranch != nil {
		_sql += fmt.Sprintf("%vcontract_code_branch = '%v', ", set, *h.ContractCodeBranch)
		set = ""
	}
	if h.ContractState != nil {
		_sql += fmt.Sprintf("%vcontract_state = '%v', ", set, h.ContractState.String())
		set = ""
	}
	if set != "" {
		return
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += fmt.Sprintf("where id = %v ", *h.ID)
	_sql += "and exists ("
	_sql += "select 1 from ("
	_sql += "select * from delegated_stakings as pr "
	_sql += fmt.Sprintf("where pr.good_id = '%v'", *h.GoodID)
	_sql += " limit 1) as tmp)"
	h.sqlDelegatedStaking = _sql
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) (int64, error) {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	return rc.RowsAffected()
}

func (h *updateHandler) updateDelegatedStaking(ctx context.Context, tx *ent.Tx) error {
	if h.sqlDelegatedStaking == "" {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlDelegatedStaking)
	if err != nil || n != 1 {
		return wlog.Errorf("fail update delegatedstaking: %v", err)
	}
	return nil
}

func (h *updateHandler) updateGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.sqlGoodBase == "" {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlGoodBase)
	if err != nil || n != 1 {
		return wlog.Errorf("fail update goodbase: %v", err)
	}
	return nil
}

//nolint:gocyclo,funlen
func (h *updateHandler) validateRewardState() error {
	if h.RewardReq.RewardState == nil {
		return nil
	}
	if *h.RewardReq.RewardState != types.BenefitState_BenefitTransferring {
		h.RewardReq.LastRewardAt = &h.goodReward.LastRewardAt
		for _, req := range h.CoinRewardReqs {
			if req.RewardTID != nil || req.LastRewardAmount != nil {
				return wlog.Errorf("invalid reward")
			}
		}
	} else if h.RewardReq.LastRewardAt == nil {
		return wlog.Errorf("invalid lastrewardat")
	}

	switch *h.RewardReq.RewardState {
	case types.BenefitState_BenefitDone:
		fallthrough // nolint
	case types.BenefitState_BenefitWait:
		coinRewardReqs := []*goodcoinrewardcrud.Req{}
		for _, coinReward := range h.coinRewards {
			coinRewardReq := &goodcoinrewardcrud.Req{
				GoodID:           &coinReward.GoodID,
				CoinTypeID:       &coinReward.CoinTypeID,
				RewardTID:        &coinReward.RewardTid,
				LastRewardAmount: &coinReward.LastRewardAmount,
			}
			for _, reward := range h.CoinRewardReqs {
				if coinReward.CoinTypeID == *reward.CoinTypeID {
					coinRewardReq.NextRewardStartAmount = reward.NextRewardStartAmount
					break
				}
			}
			coinRewardReqs = append(coinRewardReqs, coinRewardReq)
		}
		h.CoinRewardReqs = coinRewardReqs
	}

	switch h.goodReward.RewardState {
	case types.BenefitState_BenefitWait.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitTransferring:
		case types.BenefitState_BenefitFail:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitTransferring.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitBookKeeping:
		case types.BenefitState_BenefitFail:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitBookKeeping.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitUserBookKeeping:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitUserBookKeeping.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitSimulateBookKeeping:
		case types.BenefitState_BenefitDone:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitSimulateBookKeeping.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitDone:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitDone.String():
		fallthrough //nolint
	case types.BenefitState_BenefitFail.String():
		if *h.RewardReq.RewardState != types.BenefitState_BenefitWait {
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	default:
		return wlog.Errorf("invalid rewardstate")
	}
	return nil
}

func (h *updateHandler) validateGoodState(ctx context.Context) error {
	if h.GoodBaseReq.State == nil {
		return nil
	}

	currentState := types.GoodState(types.GoodState_value[h.goodBase.State]).Enum()
	nextState := h.GoodBaseReq.State

	handler, err := goodstm.NewHandler(ctx,
		goodstm.WithCurrentGoodState(currentState, true),
		goodstm.WithNextGoodState(nextState, true),
		goodstm.WithRollback(h.Rollback, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	if _state, err := handler.ValidateUpdateForNewState(); err != nil {
		return wlog.WrapError(err)
	} else if _state != nil {
		h.GoodBaseReq.State = _state
	}

	return nil
}

func (h *updateHandler) updateGoodReward(ctx context.Context, tx *ent.Tx) error {
	if h.RewardReq.RewardState == nil {
		return nil
	}
	if n, err := h.execSQL(ctx, tx, h.sqlGoodReward); err != nil || n != 1 {
		return wlog.Errorf("fail update goodreward: %v", err)
	}
	return nil
}

func (h *updateHandler) updateGoodCoinRewards(ctx context.Context, tx *ent.Tx) error {
	if h.RewardReq.RewardState == nil {
		return nil
	}
	for _, sql := range h.sqlCoinRewards {
		if n, err := h.execSQL(ctx, tx, sql); err != nil || n != 1 {
			return wlog.Errorf("fail update coinreward: %v", err)
		}
	}
	return nil
}

func (h *updateHandler) createCoinRewardHistories(ctx context.Context, tx *ent.Tx) error {
	if h.RewardReq.RewardState == nil {
		return nil
	}
	if *h.RewardReq.RewardState != types.BenefitState_BenefitDone {
		return nil
	}
	// Here reward at should be got from exist record
	for _, sql := range h.sqlCoinRewardHistories {
		if n, err := h.execSQL(ctx, tx, sql); err != nil || n != 1 {
			return wlog.Errorf("fail create coinrewardhistory: %v", err)
		}
	}
	return nil
}

func (h *Handler) UpdateDelegatedStaking(ctx context.Context) error {
	handler := &updateHandler{
		delegatedstakingGoodQueryHandler: &delegatedstakingGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireDelegatedStakingGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	h.ID = &handler.delegatedstaking.ID
	if h.GoodID == nil {
		h.GoodID = &handler.delegatedstaking.GoodID
		h.GoodBaseReq.EntID = h.GoodID
	}
	if err := handler.validateRewardState(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateGoodState(ctx); err != nil {
		return wlog.WrapError(err)
	}

	handler.constructDelegatedStakingSQL()
	if err := handler.constructGoodRewardSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructGoodBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructCoinRewardSQLs(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructCoinRewardHistorySQLs(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateGoodReward(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateGoodCoinRewards(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createCoinRewardHistories(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateDelegatedStaking(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
