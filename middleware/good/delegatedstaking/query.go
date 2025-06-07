package delegatedstaking

import (
	"context"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodcoincrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin"
	goodcoinrewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin/reward"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodcoin "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoin"
	entgoodcoinreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoinreward"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/delegatedstaking"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	goodcoinrewardmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount    *ent.GoodBaseSelect
	infos       []*npool.DelegatedStaking
	goodCoins   []*goodcoinmwpb.GoodCoinInfo
	coinRewards []*goodcoinrewardmwpb.RewardInfo
	total       uint32
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinReward(s)
		if err := h.queryJoinGoodCoin(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodCoin", "Error", err)
		}
		if err := h.queryJoinDelegatedStaking(s); err != nil {
			logger.Sugar().Errorw("queryJoinDelegatedStaking", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) getGoodCoins(ctx context.Context, cli *ent.Client) error {
	goodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			if _, err := uuid.Parse(info.GoodID); err != nil {
				continue
			}
			uids = append(uids, uuid.MustParse(info.GoodID))
		}
		return
	}()

	stm, err := goodcoincrud.SetQueryConds(
		cli.GoodCoin.Query(),
		&goodcoincrud.Conds{
			GoodIDs: &cruder.Cond{Op: cruder.IN, Val: goodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entgoodcoin.FieldGoodID,
		entgoodcoin.FieldCoinTypeID,
		entgoodcoin.FieldMain,
		entgoodcoin.FieldIndex,
	).Scan(ctx, &h.goodCoins)
}

func (h *queryHandler) getCoinRewards(ctx context.Context, cli *ent.Client) error {
	goodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			if _, err := uuid.Parse(info.GoodID); err != nil {
				continue
			}
			uids = append(uids, uuid.MustParse(info.GoodID))
		}
		return
	}()

	stm, err := goodcoinrewardcrud.SetQueryConds(
		cli.GoodCoinReward.Query(),
		&goodcoinrewardcrud.Conds{
			GoodIDs: &cruder.Cond{Op: cruder.IN, Val: goodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entgoodcoinreward.FieldGoodID,
		entgoodcoinreward.FieldCoinTypeID,
		entgoodcoinreward.FieldRewardTid,
		entgoodcoinreward.FieldNextRewardStartAmount,
		entgoodcoinreward.FieldLastRewardAmount,
		entgoodcoinreward.FieldLastUnitRewardAmount,
		entgoodcoinreward.FieldTotalRewardAmount,
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(entgoodcoin.Table)
		s.Join(t1).
			On(
				s.C(entgoodcoinreward.FieldGoodID),
				t1.C(entgoodcoin.FieldGoodID),
			).
			On(
				s.C(entgoodcoinreward.FieldCoinTypeID),
				t1.C(entgoodcoin.FieldCoinTypeID),
			).
			OnP(
				sql.EQ(t1.C(entgoodcoin.FieldDeletedAt), 0),
			).
			AppendSelect(
				sql.As(t1.C(entgoodcoin.FieldMain), "main_coin"),
			)
	}).Scan(ctx, &h.coinRewards)
}

func (h *queryHandler) formalize() {
	goodCoins := map[string][]*goodcoinmwpb.GoodCoinInfo{}
	coinRewards := map[string][]*goodcoinrewardmwpb.RewardInfo{}

	for _, goodCoin := range h.goodCoins {
		goodCoins[goodCoin.GoodID] = append(goodCoins[goodCoin.GoodID], goodCoin)
	}
	for _, coinReward := range h.coinRewards {
		coinReward.NextRewardStartAmount = func() string {
			amount, _ := decimal.NewFromString(coinReward.NextRewardStartAmount)
			return amount.String()
		}()
		coinReward.LastRewardAmount = func() string {
			amount, _ := decimal.NewFromString(coinReward.LastRewardAmount)
			return amount.String()
		}()
		coinReward.LastUnitRewardAmount = func() string {
			amount, _ := decimal.NewFromString(coinReward.LastUnitRewardAmount)
			return amount.String()
		}()
		coinReward.TotalRewardAmount = func() string {
			amount, _ := decimal.NewFromString(coinReward.TotalRewardAmount)
			return amount.String()
		}()
		coinRewards[coinReward.GoodID] = append(coinRewards[coinReward.GoodID], coinReward)
	}
	for _, info := range h.infos {
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.BenefitType = types.BenefitType(types.BenefitType_value[info.BenefitTypeStr])
		info.StartMode = types.GoodStartMode(types.GoodStartMode_value[info.StartModeStr])
		info.RewardState = types.BenefitState(types.BenefitState_value[info.RewardStateStr])
		info.ContractState = types.ContractState(types.ContractState_value[info.ContractStateStr])
		info.State = types.GoodState(types.GoodState_value[info.StateStr])
		info.GoodCoins = goodCoins[info.GoodID]
		info.Rewards = coinRewards[info.GoodID]
	}
}

func (h *Handler) GetDelegatedStaking(ctx context.Context) (*npool.DelegatedStaking, error) {
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
		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getGoodCoins(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getCoinRewards(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
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

func (h *Handler) GetDelegatedStakings(ctx context.Context) ([]*npool.DelegatedStaking, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryGoodBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryGoodBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()
		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))

		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getGoodCoins(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getCoinRewards(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
