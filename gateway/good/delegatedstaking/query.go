package delegatedstaking

import (
	"context"
	"fmt"

	contractmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/contract"
	accounttypes "github.com/NpoolPlatform/kunman/message/basetypes/account/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"
	goodcoingwpb "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
	goodcoinrewardgwpb "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin/reward"
	delegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/delegatedstaking"
	goodusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/gooduser"
	contractmw "github.com/NpoolPlatform/kunman/middleware/account/contract"
	delegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	delegatedstakings []*delegatedstakingmwpb.DelegatedStaking
	coins             map[string]*coinmwpb.Coin
	deploymentAddress map[string]*contractmwpb.Account
	calculateAddress  map[string]*contractmwpb.Account
	poolGoodUsers     map[string]*goodusermwpb.GoodUser
	infos             []*npool.DelegatedStaking
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.coins, err = goodgwcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, delegatedstaking := range h.delegatedstakings {
			for _, goodCoin := range delegatedstaking.GoodCoins {
				coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getDelegatedStakingAddress(ctx context.Context) (err error) {
	delegatedstakingIDs := []string{}
	for _, delegatedstaking := range h.delegatedstakings {
		delegatedstakingIDs = append(delegatedstakingIDs, delegatedstaking.EntID)
	}

	conds := &contractmwpb.Conds{
		DelegatedStakingIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: delegatedstakingIDs},
	}
	handler, err := contractmw.NewHandler(
		ctx,
		contractmw.WithConds(conds),
		contractmw.WithOffset(0),
		contractmw.WithLimit(int32(len(delegatedstakingIDs)*2)),
	)
	if err != nil {
		return err
	}

	accounts, _, err := handler.GetAccounts(ctx)
	if accounts == nil {
		return nil
	}
	for _, accont := range accounts {
		if accont.ContractOperatorType == accounttypes.ContractOperatorType_ContractOwner {
			h.deploymentAddress[accont.DelegatedStakingID] = accont
		}
		if accont.ContractOperatorType == accounttypes.ContractOperatorType_ContractCalculator {
			h.calculateAddress[accont.DelegatedStakingID] = accont
		}
	}
	return err
}

func (h *queryHandler) formalize() {
	for _, delegatedstaking := range h.delegatedstakings {
		deploymentAddress := ""
		deploymentAccount, ok := h.deploymentAddress[delegatedstaking.EntID]
		if ok {
			deploymentAddress = deploymentAccount.Address
		}
		calculateAddress := ""
		calculateAccount, ok := h.calculateAddress[delegatedstaking.EntID]
		if ok {
			calculateAddress = calculateAccount.Address
		}
		info := &npool.DelegatedStaking{
			ID:                        delegatedstaking.ID,
			EntID:                     delegatedstaking.EntID,
			GoodID:                    delegatedstaking.GoodID,
			ContractDeploymentAddress: deploymentAddress,
			ContractCalculateAddress:  calculateAddress,

			GoodType:             delegatedstaking.GoodType,
			BenefitType:          delegatedstaking.BenefitType,
			Name:                 delegatedstaking.Name,
			ServiceStartAt:       delegatedstaking.ServiceStartAt,
			StartMode:            delegatedstaking.StartMode,
			TestOnly:             delegatedstaking.TestOnly,
			BenefitIntervalHours: delegatedstaking.BenefitIntervalHours,
			Purchasable:          delegatedstaking.Purchasable,
			Online:               delegatedstaking.Online,
			State:                delegatedstaking.State,

			ContractCodeURL:    delegatedstaking.ContractCodeURL,
			ContractCodeBranch: delegatedstaking.ContractCodeBranch,
			ContractState:      delegatedstaking.ContractState,

			RewardState:  delegatedstaking.RewardState,
			LastRewardAt: delegatedstaking.LastRewardAt,
			Rewards: func() (rewards []*goodcoinrewardgwpb.RewardInfo) {
				for _, reward := range delegatedstaking.Rewards {
					coin, ok := h.coins[reward.CoinTypeID]
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
				for _, goodCoin := range delegatedstaking.GoodCoins {
					coin, ok := h.coins[goodCoin.CoinTypeID]
					if !ok {
						continue
					}
					coins = append(coins, &goodcoingwpb.GoodCoinInfo{
						CoinTypeID: goodCoin.CoinTypeID,
						CoinName:   coin.Name,
						CoinUnit:   coin.Unit,
						CoinENV:    coin.ENV,
						CoinLogo:   coin.Logo,
						Main:       goodCoin.Main,
						Index:      goodCoin.Index,
					})
				}
				return
			}(),

			CreatedAt: delegatedstaking.CreatedAt,
			UpdatedAt: delegatedstaking.UpdatedAt,
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetDelegatedStaking(ctx context.Context) (*npool.DelegatedStaking, error) {
	dsHandler, err := delegatedstakingmw.NewHandler(
		ctx,
		delegatedstakingmw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return nil, err
	}

	delegatedstaking, err := dsHandler.GetDelegatedStaking(ctx)
	if err != nil {
		return nil, err
	}
	if delegatedstaking == nil {
		return nil, fmt.Errorf("invalid delegatedstaking")
	}

	handler := &queryHandler{
		Handler:           h,
		delegatedstakings: []*delegatedstakingmwpb.DelegatedStaking{delegatedstaking},
		coins:             map[string]*coinmwpb.Coin{},
		deploymentAddress: map[string]*contractmwpb.Account{},
		calculateAddress:  map[string]*contractmwpb.Account{},
		poolGoodUsers:     map[string]*goodusermwpb.GoodUser{},
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, err
	}

	if err := handler.getDelegatedStakingAddress(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetDelegatedStakings(ctx context.Context) ([]*npool.DelegatedStaking, uint32, error) {
	dsHandler, err := delegatedstakingmw.NewHandler(
		ctx,
		delegatedstakingmw.WithConds(&delegatedstakingmwpb.Conds{}),
		delegatedstakingmw.WithOffset(h.Offset),
		delegatedstakingmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	delegatedstakings, total, err := dsHandler.GetDelegatedStakings(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(delegatedstakings) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:           h,
		delegatedstakings: delegatedstakings,
		coins:             map[string]*coinmwpb.Coin{},
		deploymentAddress: map[string]*contractmwpb.Account{},
		calculateAddress:  map[string]*contractmwpb.Account{},
		poolGoodUsers:     map[string]*goodusermwpb.GoodUser{},
	}

	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, err
	}

	if err := handler.getDelegatedStakingAddress(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
