package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	delegatedstakingmwcli "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"
	delegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/delegatedstaking"
)

func (h *Handler) UpdateDelegatedStaking(ctx context.Context) (*npool.DelegatedStaking, error) {
	handler := checkHandler{
		Handler: h,
	}
	if err := handler.checkDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := delegatedstakingmwcli.UpdateDelegatedStaking(ctx, &delegatedstakingmwpb.DelegatedStakingReq{
		ID:                   h.ID,
		EntID:                h.EntID,
		GoodID:               h.GoodID,
		Name:                 h.Name,
		ServiceStartAt:       h.ServiceStartAt,
		StartMode:            h.StartMode,
		TestOnly:             h.TestOnly,
		BenefitIntervalHours: h.BenefitIntervalHours,
		Purchasable:          h.Purchasable,
		Online:               h.Online,
		ContractCodeURL:      h.ContractCodeURL,
		ContractCodeBranch:   h.ContractCodeBranch,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetDelegatedStaking(ctx)
}
