package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
	goodcoingwpb "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
	goodcoinrewardgwpb "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin/reward"
	goodstockgwpb "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/stock"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	goodusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
)

type queryHandler struct {
	*Handler
	poolGoodUsers   map[string]*goodusermwpb.GoodUser
	appPowerRentals []*apppowerrentalmwpb.PowerRental
	infos           []*npool.AppPowerRental
	apps            map[string]*appmwpb.App
	appCoins        map[string]map[string]*appcoinmwpb.Coin
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, appPowerRental := range h.appPowerRentals {
			appIDs = append(appIDs, appPowerRental.AppID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.appCoins = map[string]map[string]*appcoinmwpb.Coin{}
	appCoinTypeIDs := map[string][]string{}
	for _, appPowerRental := range h.appPowerRentals {
		coinTypeIDs, ok := appCoinTypeIDs[appPowerRental.AppID]
		if !ok {
			coinTypeIDs = []string{}
		}
		for _, goodCoin := range appPowerRental.GoodCoins {
			coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
		}
		appCoinTypeIDs[appPowerRental.AppID] = coinTypeIDs
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

func (h *queryHandler) getPoolGoodUsers(ctx context.Context) (err error) {
	h.poolGoodUsers, err = goodgwcommon.GetPoolGoodUsers(ctx, func() (poolGoodUserIDs []string) {
		for _, powerRental := range h.appPowerRentals {
			for _, miningGoodStock := range powerRental.MiningGoodStocks {
				if len(miningGoodStock.PoolGoodUserID) > 0 {
					poolGoodUserIDs = append(poolGoodUserIDs, miningGoodStock.PoolGoodUserID)
				}
			}
		}
		return
	}())
	return err
}

//nolint:funlen
func (h *queryHandler) formalize() {
	for _, appPowerRental := range h.appPowerRentals {
		info := &npool.AppPowerRental{
			ID:        appPowerRental.ID,
			EntID:     appPowerRental.EntID,
			AppID:     appPowerRental.AppID,
			GoodID:    appPowerRental.GoodID,
			AppGoodID: appPowerRental.AppGoodID,

			DeviceTypeID:           appPowerRental.DeviceTypeID,
			DeviceType:             appPowerRental.DeviceType,
			DeviceManufacturerName: appPowerRental.DeviceManufacturerName,
			DeviceManufacturerLogo: appPowerRental.DeviceManufacturerLogo,
			DevicePowerConsumption: appPowerRental.DevicePowerConsumption,
			DeviceShipmentAt:       appPowerRental.DeviceShipmentAt,

			VendorLocationID: appPowerRental.VendorLocationID,
			VendorBrand:      appPowerRental.VendorBrand,
			VendorLogo:       appPowerRental.VendorLogo,
			VendorCountry:    appPowerRental.VendorCountry,
			VendorProvince:   appPowerRental.VendorProvince,

			UnitPrice:           appPowerRental.UnitPrice,
			QuantityUnit:        appPowerRental.QuantityUnit,
			QuantityUnitAmount:  appPowerRental.QuantityUnitAmount,
			DeliveryAt:          appPowerRental.DeliveryAt,
			UnitLockDeposit:     appPowerRental.UnitLockDeposit,
			DurationDisplayType: appPowerRental.DurationDisplayType,

			GoodType:             appPowerRental.GoodType,
			BenefitType:          appPowerRental.BenefitType,
			GoodName:             appPowerRental.GoodName,
			ServiceStartAt:       appPowerRental.AppGoodServiceStartAt,
			GoodStartMode:        appPowerRental.GoodStartMode,
			TestOnly:             appPowerRental.TestOnly,
			BenefitIntervalHours: appPowerRental.BenefitIntervalHours,
			GoodPurchasable:      appPowerRental.GoodPurchasable,
			GoodOnline:           appPowerRental.GoodOnline,
			State:                appPowerRental.State,

			StockMode:                    appPowerRental.StockMode,
			AppGoodPurchasable:           appPowerRental.AppGoodPurchasable,
			AppGoodOnline:                appPowerRental.AppGoodOnline,
			EnableProductPage:            appPowerRental.EnableProductPage,
			ProductPage:                  appPowerRental.ProductPage,
			Visible:                      appPowerRental.Visible,
			AppGoodName:                  appPowerRental.AppGoodName,
			DisplayIndex:                 appPowerRental.DisplayIndex,
			Banner:                       appPowerRental.Banner,
			CancelMode:                   appPowerRental.CancelMode,
			CancelableBeforeStartSeconds: appPowerRental.CancelableBeforeStartSeconds,
			EnableSetCommission:          appPowerRental.EnableSetCommission,
			MinOrderAmount:               appPowerRental.MinOrderAmount,
			MaxOrderAmount:               appPowerRental.MaxOrderAmount,
			MaxUserAmount:                appPowerRental.MaxUserAmount,
			MinOrderDurationSeconds:      appPowerRental.MinOrderDurationSeconds,
			MaxOrderDurationSeconds:      appPowerRental.MaxOrderDurationSeconds,
			SaleStartAt:                  appPowerRental.SaleStartAt,
			SaleEndAt:                    appPowerRental.SaleEndAt,
			SaleMode:                     appPowerRental.SaleMode,
			FixedDuration:                appPowerRental.FixedDuration,
			PackageWithRequireds:         appPowerRental.PackageWithRequireds,
			AppGoodStartMode:             appPowerRental.AppGoodStartMode,

			GoodStockID:      appPowerRental.GoodStockID,
			GoodTotal:        appPowerRental.GoodTotal,
			GoodSpotQuantity: appPowerRental.GoodSpotQuantity,

			AppGoodStockID:      appPowerRental.AppGoodStockID,
			AppGoodReserved:     appPowerRental.AppGoodReserved,
			AppGoodSpotQuantity: appPowerRental.AppGoodSpotQuantity,
			AppGoodLocked:       appPowerRental.AppGoodLocked,
			AppGoodInService:    appPowerRental.AppGoodInService,
			AppGoodWaitStart:    appPowerRental.AppGoodWaitStart,
			AppGoodSold:         appPowerRental.AppGoodSold,

			Likes:          appPowerRental.Likes,
			Dislikes:       appPowerRental.Dislikes,
			Score:          appPowerRental.Score,
			ScoreCount:     appPowerRental.ScoreCount,
			RecommendCount: appPowerRental.RecommendCount,
			CommentCount:   appPowerRental.CommentCount,

			LastRewardAt: appPowerRental.LastRewardAt,
			Rewards: func() (rewards []*goodcoinrewardgwpb.RewardInfo) {
				for _, reward := range appPowerRental.Rewards {
					coins, ok := h.appCoins[appPowerRental.AppID]
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
				for _, goodCoin := range appPowerRental.GoodCoins {
					appCoins, ok := h.appCoins[appPowerRental.AppID]
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
			Descriptions:        appPowerRental.Descriptions,
			Posters:             appPowerRental.Posters,
			DisplayNames:        appPowerRental.DisplayNames,
			DisplayColors:       appPowerRental.DisplayColors,
			Requireds:           appPowerRental.Requireds,
			AppMiningGoodStocks: appPowerRental.AppMiningGoodStocks,

			MiningGoodStocks: func() (mininGoodStocks []*goodstockgwpb.MiningGoodStockInfo) {
				for _, stock := range appPowerRental.MiningGoodStocks {
					mininGoodStock := &goodstockgwpb.MiningGoodStockInfo{
						EntID:          stock.EntID,
						GoodStockID:    stock.GoodStockID,
						PoolGoodUserID: stock.PoolGoodUserID,
						Total:          stock.Total,
						SpotQuantity:   stock.SpotQuantity,
					}
					if goodUser, ok := h.poolGoodUsers[stock.PoolGoodUserID]; ok {
						mininGoodStock.MiningPoolID = goodUser.PoolID
						mininGoodStock.MiningPoolName = goodUser.MiningPoolName
						mininGoodStock.MiningPoolLogo = goodUser.MiningPoolLogo
						mininGoodStock.MiningPoolSite = goodUser.MiningPoolSite
						mininGoodStock.MiningPoolReadPageLink = goodUser.ReadPageLink
					}
					mininGoodStocks = append(mininGoodStocks, mininGoodStock)
				}
				return mininGoodStocks
			}(),
			Labels: appPowerRental.Labels,

			CreatedAt: appPowerRental.CreatedAt,
			UpdatedAt: appPowerRental.UpdatedAt,
		}
		if appPowerRental.GoodType == types.GoodType_LegacyPowerRental {
			info.TechniqueFeeRatio = &appPowerRental.TechniqueFeeRatio
		}
		app, ok := h.apps[appPowerRental.AppID]
		if ok {
			info.AppName = app.Name
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetPowerRental(ctx context.Context) (*npool.AppPowerRental, error) {
	appPowerRental, err := apppowerrentalmwcli.GetPowerRental(ctx, *h.AppGoodID)
	if err != nil {
		return nil, err
	}
	if appPowerRental == nil {
		return nil, wlog.Errorf("invalid apppowerrental")
	}

	handler := &queryHandler{
		Handler:         h,
		appPowerRentals: []*apppowerrentalmwpb.PowerRental{appPowerRental},
		apps:            map[string]*appmwpb.App{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, err
	}
	if err := handler.getPoolGoodUsers(ctx); err != nil {
		return nil, err
	}
	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetPowerRentals(ctx context.Context) ([]*npool.AppPowerRental, uint32, error) {
	appPowerRentals, total, err := apppowerrentalmwcli.GetPowerRentals(ctx, &apppowerrentalmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(appPowerRentals) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:         h,
		appPowerRentals: appPowerRentals,
		apps:            map[string]*appmwpb.App{},
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
