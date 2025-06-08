package delegatedstaking

import (
	"context"
	"fmt"

	contractmwcli "github.com/NpoolPlatform/account-middleware/pkg/client/contract"
	coinmwcli "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	delegatedstakingmwcli "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"
	goodcoinmwcli "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	contractmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/contract"
	accounttypes "github.com/NpoolPlatform/kunman/message/basetypes/account/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"
	delegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/delegatedstaking"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
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
	coin, err := coinmwcli.GetCoin(ctx, *h.CoinTypeID)
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
	main := true
	if err := goodcoinmwcli.CreateGoodCoin(ctx, &goodcoinmwpb.GoodCoinReq{
		GoodID:     h.GoodID,
		CoinTypeID: h.CoinTypeID,
		Main:       &main,
	}); err != nil {
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
	_, err = contractmwcli.CreateAccount(ctx, &contractmwpb.AccountReq{
		GoodID:               h.GoodID,
		DelegatedStakingID:   h.EntID,
		ContractOperatorType: accounttypes.ContractOperatorType_ContractOwner.Enum(),
		CoinTypeID:           h.CoinTypeID,
		Address:              h.contractDevelopmentAddress,
	})
	if err != nil {
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
	_, err = contractmwcli.CreateAccount(ctx, &contractmwpb.AccountReq{
		GoodID:               h.GoodID,
		DelegatedStakingID:   h.EntID,
		ContractOperatorType: accounttypes.ContractOperatorType_ContractCalculator.Enum(),
		CoinTypeID:           h.CoinTypeID,
		Address:              h.contractCalculateAddress,
	})
	if err != nil {
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
	if err := delegatedstakingmwcli.CreateDelegatedStaking(ctx, &delegatedstakingmwpb.DelegatedStakingReq{
		EntID:                h.EntID,
		GoodID:               h.GoodID,
		GoodType:             h.GoodType,
		Name:                 h.Name,
		ServiceStartAt:       h.ServiceStartAt,
		StartMode:            h.StartMode,
		TestOnly:             h.TestOnly,
		BenefitIntervalHours: h.BenefitIntervalHours,
		Purchasable:          h.Purchasable,
		Online:               h.Online,
		ContractCodeURL:      h.ContractCodeURL,
		ContractCodeBranch:   h.ContractCodeBranch,
		ContractState:        h.ContractState,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.createGoodCoin(ctx); err != nil {
		return nil, err
	}
	return h.GetDelegatedStaking(ctx)
}
