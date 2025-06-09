package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"
	delegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"
)

func (h *Handler) UpdateDelegatedStaking(ctx context.Context) (*npool.DelegatedStaking, error) {
	handler := checkHandler{
		Handler: h,
	}
	if err := handler.checkDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	dsHandler, err := delegatedstakingmw.NewHandler(
		ctx,
		delegatedstakingmw.WithEntID(h.EntID, true),
		delegatedstakingmw.WithGoodID(h.GoodID, true),
		delegatedstakingmw.WithName(h.Name, false),
		delegatedstakingmw.WithServiceStartAt(h.ServiceStartAt, false),
		delegatedstakingmw.WithStartMode(h.StartMode, false),
		delegatedstakingmw.WithTestOnly(h.TestOnly, false),
		delegatedstakingmw.WithBenefitIntervalHours(h.BenefitIntervalHours, false),
		delegatedstakingmw.WithPurchasable(h.Purchasable, false),
		delegatedstakingmw.WithOnline(h.Online, false),
		delegatedstakingmw.WithContractCodeURL(h.ContractCodeURL, false),
		delegatedstakingmw.WithContractCodeBranch(h.ContractCodeBranch, false),
	)
	if err != nil {
		return nil, err
	}

	if err := dsHandler.UpdateDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetDelegatedStaking(ctx)
}
