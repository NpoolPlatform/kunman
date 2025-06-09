package registetestinfo

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	mpbasetype "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmw "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/coin"
	fractionwithdrawalrulemw "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawalrule"
	poolmw "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/pool"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/config"
	"github.com/NpoolPlatform/kunman/pkg/const/time"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/coin"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/fractionwithdrawalrule"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool"
)

var (
	poolInfos = []*poolmw.Pool{
		{
			MiningPoolType: mpbasetype.MiningPoolType_F2Pool,
			Name:           mpbasetype.MiningPoolType_F2Pool.String(),
			Site:           config.F2PoolSite,
			Logo:           "https://static.f2pool.com/static/images/icon-logo-min-blue.png",
			Description:    "",
		},
	}
	poolEntIDs                   = []*uuid.UUID{}
	coinEntIDs                   = []*uuid.UUID{}
	fractionwithdrawalruleEntIDs = []*uuid.UUID{}

	coinInfos = []*coinmw.Coin{
		{
			MiningPoolType:         mpbasetype.MiningPoolType_F2Pool,
			CoinType:               basetypes.CoinType_CoinTypeBitCoin,
			FeeRatio:               f2pool.CoinType2FeeRatio[basetypes.CoinType_CoinTypeBitCoin],
			FixedRevenueAble:       false,
			LeastTransferAmount:    f2pool.CoinType2LeastTransferAmount[basetypes.CoinType_CoinTypeBitCoin],
			BenefitIntervalSeconds: f2pool.CoinType2BenefitIntervalSeconds[basetypes.CoinType_CoinTypeBitCoin],
		},
	}

	fractionwithdrawalruleInfos = []*fractionwithdrawalrulemw.FractionWithdrawalRule{
		{
			MiningPoolType: mpbasetype.MiningPoolType_F2Pool,
			CoinType:       basetypes.CoinType_CoinTypeBitCoin,
			// 30Day
			WithdrawInterval:      time.SecondsPerDay * 30,
			PayoutThreshold:       "0.005",
			LeastWithdrawalAmount: "0.0005",
			WithdrawFee:           "0",
		},
	}
)

func registePool(ctx context.Context) {
	for _, info := range poolInfos {
		poolH, err := pool.NewHandler(ctx, pool.WithConds(&poolmw.Conds{
			MiningPoolType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(info.MiningPoolType),
			},
		}))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		exist, err := poolH.ExistPoolConds(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		if exist {
			continue
		}

		info.EntID = uuid.NewString()
		poolH, err = pool.NewHandler(ctx,
			pool.WithEntID(&info.EntID, true),
			pool.WithMiningPoolType(&info.MiningPoolType, true),
			pool.WithName(&info.Name, true),
			pool.WithSite(&info.Site, true),
			pool.WithLogo(&info.Logo, true),
			pool.WithDescription(&info.Description, true),
		)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		err = poolH.CreatePool(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
		poolEntIDs = append(poolEntIDs, poolH.EntID)
	}
}

func registeCoinInfo(ctx context.Context) {
	for _, info := range coinInfos {
		// check if exist
		coinH, err := coin.NewHandler(ctx, coin.WithConds(&coinmw.Conds{
			CoinType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(info.CoinType),
			},
			MiningPoolType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(info.MiningPoolType),
			},
		}))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		exist, err := coinH.ExistCoinConds(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		if exist {
			continue
		}

		// get pools
		poolH, err := pool.NewHandler(ctx,
			pool.WithConds(&poolmw.Conds{
				MiningPoolType: &basetypes.Uint32Val{
					Op:    cruder.EQ,
					Value: uint32(info.MiningPoolType),
				},
			}),
			pool.WithOffset(0),
			pool.WithLimit(2),
		)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		poolInfos, _, err := poolH.GetPools(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		if len(poolInfos) == 0 {
			logger.Sugar().Errorf("have no pool of %v", info.MiningPoolType)
			continue
		}

		info.EntID = uuid.NewString()
		info.CoinTypeID = uuid.NewString()
		// create coin
		coinH, err = coin.NewHandler(ctx,
			coin.WithEntID(&info.EntID, true),
			coin.WithCoinTypeID(&info.CoinTypeID, true),
			coin.WithPoolID(&poolInfos[0].EntID, true),
			coin.WithCoinType(&info.CoinType, true),
			coin.WithFeeRatio(&info.FeeRatio, true),
			coin.WithFixedRevenueAble(&info.FixedRevenueAble, true),
			coin.WithLeastTransferAmount(&info.LeastTransferAmount, true),
			coin.WithBenefitIntervalSeconds(&info.BenefitIntervalSeconds, true),
		)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		err = coinH.CreateCoin(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
		coinEntIDs = append(coinEntIDs, coinH.EntID)
	}
}

func registeFractionWithdrawalRule(ctx context.Context) {
	for _, info := range fractionwithdrawalruleInfos {
		coinH, err := coin.NewHandler(ctx,
			coin.WithConds(&coinmw.Conds{
				CoinType: &basetypes.Uint32Val{
					Op:    cruder.EQ,
					Value: uint32(info.CoinType),
				},
				MiningPoolType: &basetypes.Uint32Val{
					Op:    cruder.EQ,
					Value: uint32(info.MiningPoolType),
				},
			}),
			coin.WithOffset(0),
			coin.WithLimit(2))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		// get coins
		coinInfos, _, err := coinH.GetCoins(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		if len(coinInfos) == 0 {
			logger.Sugar().Errorf("have no coin of %v", info.CoinType)
			continue
		}
		// check if exist
		fractionwithdrawalruleH, err := fractionwithdrawalrule.NewHandler(ctx, fractionwithdrawalrule.WithConds(
			&fractionwithdrawalrulemw.Conds{
				PoolCoinTypeID: &basetypes.StringVal{
					Op:    cruder.EQ,
					Value: coinInfos[0].EntID,
				},
			},
		))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		exist, err := fractionwithdrawalruleH.ExistFractionWithdrawalRuleConds(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		if exist {
			continue
		}

		info.EntID = uuid.NewString()
		// create fractionwithdrawal rule
		fractionwithdrawalruleH, err = fractionwithdrawalrule.NewHandler(ctx,
			fractionwithdrawalrule.WithEntID(&info.EntID, true),
			fractionwithdrawalrule.WithPoolCoinTypeID(&coinInfos[0].EntID, true),
			fractionwithdrawalrule.WithWithdrawInterval(&info.WithdrawInterval, true),
			fractionwithdrawalrule.WithPayoutThreshold(&info.PayoutThreshold, true),
			fractionwithdrawalrule.WithLeastWithdrawalAmount(&info.LeastWithdrawalAmount, true),
			fractionwithdrawalrule.WithWithdrawFee(&info.WithdrawFee, true),
		)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		err = fractionwithdrawalruleH.CreateFractionWithdrawalRule(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
		fractionwithdrawalruleEntIDs = append(fractionwithdrawalruleEntIDs, fractionwithdrawalruleH.EntID)
	}
}

func cleanPools(ctx context.Context) {
	for _, entid := range poolEntIDs {
		_entid := entid.String()
		poolH, err := pool.NewHandler(ctx, pool.WithEntID(&_entid, true))
		if err != nil {
			logger.Sugar().Error(err)
		}
		err = poolH.DeletePool(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}

func cleanCoins(ctx context.Context) {
	for _, entid := range coinEntIDs {
		_entid := entid.String()
		coinH, err := coin.NewHandler(ctx, coin.WithEntID(&_entid, true))
		if err != nil {
			logger.Sugar().Error(err)
		}
		err = coinH.DeleteCoin(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}

func cleanFractionWithdrawalrules(ctx context.Context) {
	for _, entid := range fractionwithdrawalruleEntIDs {
		_entid := entid.String()
		fractionwithdrawalruleH, err := fractionwithdrawalrule.NewHandler(ctx, fractionwithdrawalrule.WithEntID(&_entid, true))
		if err != nil {
			logger.Sugar().Error(err)
		}
		err = fractionwithdrawalruleH.DeleteFractionWithdrawalRule(ctx)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}

func InitTestInfo(ctx context.Context) {
	registePool(ctx)
	registeCoinInfo(ctx)
	registeFractionWithdrawalRule(ctx)
}

func CleanTestInfo(ctx context.Context) {
	cleanFractionWithdrawalrules(ctx)
	cleanCoins(ctx)
	cleanPools(ctx)
}
