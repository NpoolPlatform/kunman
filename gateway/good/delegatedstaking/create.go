package delegatedstaking

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	accounttypes "github.com/NpoolPlatform/kunman/message/basetypes/account/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"
	contractmw "github.com/NpoolPlatform/kunman/middleware/account/contract"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	delegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"
	goodcoinmw "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	goodCoinName               *string
	contractDevelopmentAddress *string
	contractCalculateAddress   *string
}

func (h *createHandler) getCoin(ctx context.Context) error {
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithEntID(h.CoinTypeID, true),
	)
	if err != nil {
		return err
	}

	coin, err := handler.GetCoin(ctx)
	if err != nil {
		return err
	}
	if coin == nil {
		return fmt.Errorf("invalid coin")
	}

	h.goodCoinName = &coin.Name

	return nil
}

func (h *createHandler) createGoodCoin(ctx context.Context) error {
	handler, err := goodcoinmw.NewHandler(
		ctx,
		goodcoinmw.WithGoodID(h.GoodID, true),
		goodcoinmw.WithCoinTypeID(h.CoinTypeID, true),
		goodcoinmw.WithMain(func() *bool { b := true; return &b }(), true),
	)
	if err != nil {
		return err
	}

	if err := handler.CreateGoodCoin(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:dupl
func (h *createHandler) createDevelopmentAddress(ctx context.Context) error {
	acc, err := sphinxproxycli.CreateAddress(ctx, *h.goodCoinName)
	if err != nil {
		return err
	}
	if acc == nil {
		return fmt.Errorf("fail create address")
	}
	h.contractDevelopmentAddress = &acc.Address

	handler, err := contractmw.NewHandler(
		ctx,
		contractmw.WithGoodID(h.GoodID, true),
		contractmw.WithDelegatedStakingID(h.EntID, true),
		contractmw.WithContractOperatorType(accounttypes.ContractOperatorType_ContractOwner.Enum(), true),
		contractmw.WithCoinTypeID(h.CoinTypeID, true),
		contractmw.WithAddress(h.contractDevelopmentAddress, true),
	)
	if err != nil {
		return err
	}

	if err := handler.CreateAccount(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:dupl
func (h *createHandler) createCalculateAddress(ctx context.Context) error {
	acc, err := sphinxproxycli.CreateAddress(ctx, *h.goodCoinName)
	if err != nil {
		return err
	}
	if acc == nil {
		return fmt.Errorf("fail create address")
	}
	h.contractCalculateAddress = &acc.Address

	handler, err := contractmw.NewHandler(
		ctx,
		contractmw.WithGoodID(h.GoodID, true),
		contractmw.WithDelegatedStakingID(h.EntID, true),
		contractmw.WithContractOperatorType(accounttypes.ContractOperatorType_ContractCalculator.Enum(), true),
		contractmw.WithCoinTypeID(h.CoinTypeID, true),
		contractmw.WithAddress(h.contractCalculateAddress, true),
	)
	if err != nil {
		return err
	}

	if err := handler.CreateAccount(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateDelegatedStaking(ctx context.Context) (*npool.DelegatedStaking, error) {
	if h.GoodID == nil {
		h.GoodID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.getCoin(ctx); err != nil {
		return nil, err
	}
	if err := handler.createDevelopmentAddress(ctx); err != nil {
		return nil, err
	}
	if err := handler.createCalculateAddress(ctx); err != nil {
		return nil, err
	}

	dsHandler, err := delegatedstakingmw.NewHandler(
		ctx,
		delegatedstakingmw.WithEntID(h.EntID, true),
		delegatedstakingmw.WithGoodID(h.GoodID, true),
		delegatedstakingmw.WithGoodType(h.GoodType, true),
		delegatedstakingmw.WithName(h.Name, true),
		delegatedstakingmw.WithServiceStartAt(h.ServiceStartAt, true),
		delegatedstakingmw.WithStartMode(h.StartMode, true),
		delegatedstakingmw.WithTestOnly(h.TestOnly, true),
		delegatedstakingmw.WithBenefitIntervalHours(h.BenefitIntervalHours, true),
		delegatedstakingmw.WithPurchasable(h.Purchasable, true),
		delegatedstakingmw.WithOnline(h.Online, true),
		delegatedstakingmw.WithContractCodeURL(h.ContractCodeURL, true),
		delegatedstakingmw.WithContractCodeBranch(h.ContractCodeBranch, true),
		delegatedstakingmw.WithContractState(h.ContractState, true),
	)
	if err != nil {
		return nil, err
	}

	if err := dsHandler.CreateDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.createGoodCoin(ctx); err != nil {
		return nil, err
	}
	return h.GetDelegatedStaking(ctx)
}
