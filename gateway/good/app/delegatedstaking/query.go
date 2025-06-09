package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
	goodcoingwpb "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
	goodcoinrewardgwpb "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin/reward"
	appdelegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/delegatedstaking"
	appdelegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/app/delegatedstaking"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	appDelegatedStakings []*appdelegatedstakingmwpb.DelegatedStaking
	infos                []*npool.AppDelegatedStaking
	apps                 map[string]*appmwpb.App
	appCoins             map[string]map[string]*appcoinmwpb.Coin
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, appDelegatedStaking := range h.appDelegatedStakings {
			appIDs = append(appIDs, appDelegatedStaking.AppID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.appCoins = map[string]map[string]*appcoinmwpb.Coin{}
	appCoinTypeIDs := map[string][]string{}
	for _, appDelegatedStaking := range h.appDelegatedStakings {
		coinTypeIDs, ok := appCoinTypeIDs[appDelegatedStaking.AppID]
		if !ok {
			coinTypeIDs = []string{}
		}
		for _, goodCoin := range appDelegatedStaking.GoodCoins {
			coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
		}
		appCoinTypeIDs[appDelegatedStaking.AppID] = coinTypeIDs
	}
	for appID, coinTypeIDs := range appCoinTypeIDs {
		coins, err := goodgwcommon.GetAppCoins(ctx, appID, coinTypeIDs)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.appCoins[appID] = coins
	}
	return wlog.WrapError(err)
}

//nolint:funlen
func (h *queryHandler) formalize() {
	for _, appDelegatedStaking := range h.appDelegatedStakings {
		info := &npool.AppDelegatedStaking{
			ID:        appDelegatedStaking.ID,
			EntID:     appDelegatedStaking.EntID,
			AppID:     appDelegatedStaking.AppID,
			GoodID:    appDelegatedStaking.GoodID,
			AppGoodID: appDelegatedStaking.AppGoodID,

			GoodType:             appDelegatedStaking.GoodType,
			BenefitType:          appDelegatedStaking.BenefitType,
			GoodName:             appDelegatedStaking.GoodName,
			ServiceStartAt:       appDelegatedStaking.AppGoodServiceStartAt,
			GoodStartMode:        appDelegatedStaking.GoodStartMode,
			TestOnly:             appDelegatedStaking.TestOnly,
			BenefitIntervalHours: appDelegatedStaking.BenefitIntervalHours,
			GoodPurchasable:      appDelegatedStaking.GoodPurchasable,
			GoodOnline:           appDelegatedStaking.GoodOnline,
			State:                appDelegatedStaking.State,

			AppGoodPurchasable:  appDelegatedStaking.AppGoodPurchasable,
			AppGoodOnline:       appDelegatedStaking.AppGoodOnline,
			EnableProductPage:   appDelegatedStaking.EnableProductPage,
			ProductPage:         appDelegatedStaking.ProductPage,
			Visible:             appDelegatedStaking.Visible,
			AppGoodName:         appDelegatedStaking.AppGoodName,
			DisplayIndex:        appDelegatedStaking.DisplayIndex,
			Banner:              appDelegatedStaking.Banner,
			EnableSetCommission: appDelegatedStaking.EnableSetCommission,
			AppGoodStartMode:    appDelegatedStaking.AppGoodStartMode,

			Likes:          appDelegatedStaking.Likes,
			Dislikes:       appDelegatedStaking.Dislikes,
			Score:          appDelegatedStaking.Score,
			ScoreCount:     appDelegatedStaking.ScoreCount,
			RecommendCount: appDelegatedStaking.RecommendCount,
			CommentCount:   appDelegatedStaking.CommentCount,

			LastRewardAt: appDelegatedStaking.LastRewardAt,
			Rewards: func() (rewards []*goodcoinrewardgwpb.RewardInfo) {
				for _, reward := range appDelegatedStaking.Rewards {
					coins, ok := h.appCoins[appDelegatedStaking.AppID]
					if !ok {
						continue
					}
					coin, ok := coins[reward.CoinTypeID]
					if !ok {
						continue
					}
					rewards = append(rewards, &goodcoinrewardgwpb.RewardInfo{
						CoinTypeID:            reward.CoinTypeID,
						CoinName:              coin.Name,
						CoinUnit:              coin.Unit,
						CoinENV:               coin.ENV,
						CoinLogo:              coin.Logo,
						RewardTID:             reward.RewardTID,
						NextRewardStartAmount: reward.NextRewardStartAmount,
						LastRewardAmount:      reward.LastRewardAmount,
						LastUnitRewardAmount:  reward.LastUnitRewardAmount,
						TotalRewardAmount:     reward.TotalRewardAmount,
						MainCoin:              reward.MainCoin,
					})
				}
				return
			}(),

			GoodCoins: func() (coins []*goodcoingwpb.GoodCoinInfo) {
				for _, goodCoin := range appDelegatedStaking.GoodCoins {
					appCoins, ok := h.appCoins[appDelegatedStaking.AppID]
					if !ok {
						continue
					}
					appCoin, ok := appCoins[goodCoin.CoinTypeID]
					if !ok {
						continue
					}
					coins = append(coins, &goodcoingwpb.GoodCoinInfo{
						CoinTypeID: goodCoin.CoinTypeID,
						CoinName:   appCoin.Name,
						CoinUnit:   appCoin.Unit,
						CoinENV:    appCoin.ENV,
						CoinLogo:   appCoin.Logo,
						Main:       goodCoin.Main,
						Index:      goodCoin.Index,
					})
				}
				return
			}(),
			Descriptions:  appDelegatedStaking.Descriptions,
			Posters:       appDelegatedStaking.Posters,
			DisplayNames:  appDelegatedStaking.DisplayNames,
			DisplayColors: appDelegatedStaking.DisplayColors,
			Labels:        appDelegatedStaking.Labels,

			CreatedAt: appDelegatedStaking.CreatedAt,
			UpdatedAt: appDelegatedStaking.UpdatedAt,
		}

		app, ok := h.apps[appDelegatedStaking.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetDelegatedStaking(ctx context.Context) (*npool.AppDelegatedStaking, error) {
	dsHandler, err := appdelegatedstakingmw.NewHandler(
		ctx,
		appdelegatedstakingmw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, err
	}

	appDelegatedStaking, err := dsHandler.GetDelegatedStaking(ctx)
	if err != nil {
		return nil, err
	}
	if appDelegatedStaking == nil {
		return nil, wlog.Errorf("invalid appdelegatedstaking")
	}

	handler := &queryHandler{
		Handler:              h,
		appDelegatedStakings: []*appdelegatedstakingmwpb.DelegatedStaking{appDelegatedStaking},
		apps:                 map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, err
	}
	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetDelegatedStakings(ctx context.Context) ([]*npool.AppDelegatedStaking, uint32, error) {
	conds := &appdelegatedstakingmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	dsHandler, err := appdelegatedstakingmw.NewHandler(
		ctx,
		appdelegatedstakingmw.WithConds(conds),
		appdelegatedstakingmw.WithOffset(h.Offset),
		appdelegatedstakingmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	appDelegatedStakings, total, err := dsHandler.GetDelegatedStakings(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(appDelegatedStakings) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:              h,
		appDelegatedStakings: appDelegatedStakings,
		apps:                 map[string]*appmwpb.App{},
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
