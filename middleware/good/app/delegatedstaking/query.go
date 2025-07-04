//nolint:dupl
package delegatedstaking

import (
	"context"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/delegatedstaking"
	appgooddescriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/description"
	appgooddisplaycolormwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/display/color"
	appgooddisplaynamemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/display/name"
	appgoodlabelmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/label"
	appgoodpostermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/poster"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	goodcoinrewardmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward"
	appgooddescriptioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/description"
	appgooddisplaycolorcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/display/color"
	appgooddisplaynamecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/display/name"
	appgoodlabelcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/label"
	appgoodpostercrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/poster"
	requiredappgoodcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/required"
	goodcoincrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin"
	goodcoinrewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin/reward"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappfee "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appfee"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entappgooddescription "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgooddescription"
	entappgooddisplaycolor "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgooddisplaycolor"
	entappgooddisplayname "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgooddisplayname"
	entappgoodlabel "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodlabel"
	entappgoodposter "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodposter"
	entfee "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/fee"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entgoodcoin "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoin"
	entgoodcoinreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoinreward"
	entrequiredappgood "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/requiredappgood"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount         *ent.AppGoodBaseSelect
	infos            []*npool.DelegatedStaking
	goodCoins        []*goodcoinmwpb.GoodCoinInfo
	descriptions     []*appgooddescriptionmwpb.DescriptionInfo
	posters          []*appgoodpostermwpb.PosterInfo
	labels           []*appgoodlabelmwpb.LabelInfo
	displayNames     []*appgooddisplaynamemwpb.DisplayNameInfo
	displayColors    []*appgooddisplaycolormwpb.DisplayColorInfo
	coinRewards      []*goodcoinrewardmwpb.RewardInfo
	requiredAppGoods []*requiredappgoodmwpb.RequiredInfo
	total            uint32
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinGoodBase(s)
		h.queryJoinGoodReward(s)
		h.queryJoinExtraInfo(s)
		if err := h.queryJoinAppDelegatedStaking(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppDelegatedStaking", "Error", err)
		}
		if err := h.queryJoinDelegatedStaking(s); err != nil {
			logger.Sugar().Errorw("queryJoinDelegatedStaking", "Error", err)
		}
		if err := h.queryJoinGoodCoin(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodCoin", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) getGoodCoins(ctx context.Context, cli *ent.Client) error {
	goodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
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

func (h *queryHandler) getDescriptions(ctx context.Context, cli *ent.Client) error {
	appGoodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.AppGoodID))
		}
		return
	}()

	stm, err := appgooddescriptioncrud.SetQueryConds(
		cli.AppGoodDescription.Query(),
		&appgooddescriptioncrud.Conds{
			AppGoodIDs: &cruder.Cond{Op: cruder.IN, Val: appGoodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entappgooddescription.FieldAppGoodID,
		entappgooddescription.FieldDescription,
		entappgooddescription.FieldIndex,
	).Scan(ctx, &h.descriptions)
}

func (h *queryHandler) getPosters(ctx context.Context, cli *ent.Client) error {
	appGoodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.AppGoodID))
		}
		return
	}()

	stm, err := appgoodpostercrud.SetQueryConds(
		cli.AppGoodPoster.Query(),
		&appgoodpostercrud.Conds{
			AppGoodIDs: &cruder.Cond{Op: cruder.IN, Val: appGoodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entappgoodposter.FieldAppGoodID,
		entappgoodposter.FieldPoster,
		entappgoodposter.FieldIndex,
	).Scan(ctx, &h.posters)
}

func (h *queryHandler) getLabels(ctx context.Context, cli *ent.Client) error {
	appGoodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.AppGoodID))
		}
		return
	}()

	stm, err := appgoodlabelcrud.SetQueryConds(
		cli.AppGoodLabel.Query(),
		&appgoodlabelcrud.Conds{
			AppGoodIDs: &cruder.Cond{Op: cruder.IN, Val: appGoodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entappgoodlabel.FieldAppGoodID,
		entappgoodlabel.FieldIcon,
		entappgoodlabel.FieldIconBgColor,
		entappgoodlabel.FieldLabel,
		entappgoodlabel.FieldLabelBgColor,
		entappgoodlabel.FieldIndex,
	).Scan(ctx, &h.labels)
}

func (h *queryHandler) getDisplayNames(ctx context.Context, cli *ent.Client) error {
	appGoodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.AppGoodID))
		}
		return
	}()

	stm, err := appgooddisplaynamecrud.SetQueryConds(
		cli.AppGoodDisplayName.Query(),
		&appgooddisplaynamecrud.Conds{
			AppGoodIDs: &cruder.Cond{Op: cruder.IN, Val: appGoodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entappgooddisplayname.FieldAppGoodID,
		entappgooddisplayname.FieldName,
		entappgooddisplayname.FieldIndex,
	).Scan(ctx, &h.displayNames)
}

func (h *queryHandler) getDisplayColors(ctx context.Context, cli *ent.Client) error {
	appGoodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.AppGoodID))
		}
		return
	}()

	stm, err := appgooddisplaycolorcrud.SetQueryConds(
		cli.AppGoodDisplayColor.Query(),
		&appgooddisplaycolorcrud.Conds{
			AppGoodIDs: &cruder.Cond{Op: cruder.IN, Val: appGoodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entappgooddisplaycolor.FieldAppGoodID,
		entappgooddisplaycolor.FieldColor,
		entappgooddisplaycolor.FieldIndex,
	).Scan(ctx, &h.displayColors)
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

func (h *queryHandler) getRequiredAppGoods(ctx context.Context, cli *ent.Client) error {
	appGoodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.AppGoodID))
		}
		return
	}()

	stm, err := requiredappgoodcrud.SetQueryConds(
		cli.RequiredAppGood.Query(),
		&requiredappgoodcrud.Conds{
			MainAppGoodIDs: &cruder.Cond{Op: cruder.IN, Val: appGoodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	return stm.Select(
		entrequiredappgood.FieldMainAppGoodID,
		entrequiredappgood.FieldRequiredAppGoodID,
		entrequiredappgood.FieldMust,
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(entappgoodbase.Table)

		s.Join(t1).
			On(
				s.C(entrequiredappgood.FieldRequiredAppGoodID),
				t1.C(entappgoodbase.FieldEntID),
			).
			AppendSelect(
				sql.As(t1.C(entappgoodbase.FieldName), "required_app_good_name"),
			)

		t2 := sql.Table(entgoodbase.Table)
		s.Join(t2).
			On(
				t1.C(entappgoodbase.FieldGoodID),
				t2.C(entgoodbase.FieldEntID),
			).
			AppendSelect(
				sql.As(t2.C(entgoodbase.FieldGoodType), "required_good_type"),
			)

		t3 := sql.Table(entappfee.Table)
		s.Join(t3).
			On(
				s.C(entrequiredappgood.FieldRequiredAppGoodID),
				t3.C(entappfee.FieldAppGoodID),
			).
			AppendSelect(
				sql.As(t3.C(entappfee.FieldUnitValue), "required_app_good_unit_value"),
			)

		t4 := sql.Table(entfee.Table)
		s.Join(t4).
			On(
				t1.C(entappgoodbase.FieldGoodID),
				t4.C(entfee.FieldGoodID),
			).
			AppendSelect(
				sql.As(t4.C(entfee.FieldSettlementType), "required_good_settlement_type"),
			)
	}).Scan(ctx, &h.requiredAppGoods)
}

func (h *queryHandler) formalize() {
	goodCoins := map[string][]*goodcoinmwpb.GoodCoinInfo{}
	descriptions := map[string][]*appgooddescriptionmwpb.DescriptionInfo{}
	posters := map[string][]*appgoodpostermwpb.PosterInfo{}
	labels := map[string][]*appgoodlabelmwpb.LabelInfo{}
	displayNames := map[string][]*appgooddisplaynamemwpb.DisplayNameInfo{}
	displayColors := map[string][]*appgooddisplaycolormwpb.DisplayColorInfo{}
	coinRewards := map[string][]*goodcoinrewardmwpb.RewardInfo{}
	requireds := map[string][]*requiredappgoodmwpb.RequiredInfo{}
	for _, goodCoin := range h.goodCoins {
		goodCoins[goodCoin.GoodID] = append(goodCoins[goodCoin.GoodID], goodCoin)
	}
	for _, description := range h.descriptions {
		descriptions[description.AppGoodID] = append(descriptions[description.AppGoodID], description)
	}
	for _, poster := range h.posters {
		posters[poster.AppGoodID] = append(posters[poster.AppGoodID], poster)
	}
	for _, label := range h.labels {
		label.Label = types.GoodLabel(types.GoodLabel_value[label.LabelStr])
		labels[label.AppGoodID] = append(labels[label.AppGoodID], label)
	}
	for _, displayName := range h.displayNames {
		displayNames[displayName.AppGoodID] = append(displayNames[displayName.AppGoodID], displayName)
	}
	for _, displayColor := range h.displayColors {
		displayColors[displayColor.AppGoodID] = append(displayColors[displayColor.AppGoodID], displayColor)
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
	for _, required := range h.requiredAppGoods {
		required.RequiredGoodSettlementType = types.GoodSettlementType(types.GoodSettlementType_value[required.RequiredGoodSettlementTypeStr])
		required.RequiredGoodType = types.GoodType(types.GoodType_value[required.RequiredGoodTypeStr])
		requireds[required.MainAppGoodID] = append(requireds[required.MainAppGoodID], required)
	}
	for _, info := range h.infos {
		info.Score = func() string { amount, _ := decimal.NewFromString(info.Score); return amount.String() }()
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.BenefitType = types.BenefitType(types.BenefitType_value[info.BenefitTypeStr])
		info.GoodStartMode = types.GoodStartMode(types.GoodStartMode_value[info.GoodStartModeStr])
		info.AppGoodStartMode = types.GoodStartMode(types.GoodStartMode_value[info.AppGoodStartModeStr])
		info.State = types.GoodState(types.GoodState_value[info.StateStr])
		info.ContractState = types.ContractState(types.ContractState_value[info.ContractStateStr])
		info.GoodCoins = goodCoins[info.GoodID]
		info.Descriptions = descriptions[info.AppGoodID]
		info.Posters = posters[info.AppGoodID]
		info.Labels = labels[info.AppGoodID]
		info.DisplayNames = displayNames[info.AppGoodID]
		info.DisplayColors = displayColors[info.AppGoodID]
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
		if err := handler.queryAppGoodBase(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getGoodCoins(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getRequiredAppGoods(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getDescriptions(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getPosters(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getLabels(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getDisplayNames(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getDisplayColors(_ctx, cli); err != nil {
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
		handler.stmSelect, err = handler.queryAppGoodBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryAppGoodBases(cli)
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
		if err := handler.getRequiredAppGoods(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getDescriptions(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getPosters(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getLabels(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getDisplayNames(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getDisplayColors(_ctx, cli); err != nil {
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
