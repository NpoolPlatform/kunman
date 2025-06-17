package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/kunman/gateway/good/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"
)

func (s *Server) AdminUpdateDelegatedStaking(ctx context.Context, in *npool.AdminUpdateDelegatedStakingRequest) (*npool.AdminUpdateDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithID(&in.ID, true),
		delegatedstaking1.WithEntID(&in.EntID, true),
		delegatedstaking1.WithGoodID(&in.GoodID, true),

		delegatedstaking1.WithContractCodeURL(in.ContractCodeURL, false),
		delegatedstaking1.WithContractCodeBranch(in.ContractCodeBranch, false),
		delegatedstaking1.WithName(in.Name, false),
		delegatedstaking1.WithServiceStartAt(in.ServiceStartAt, false),
		delegatedstaking1.WithStartMode(in.StartMode, false),
		delegatedstaking1.WithTestOnly(in.TestOnly, false),
		delegatedstaking1.WithBenefitIntervalHours(in.BenefitIntervalHours, false),
		delegatedstaking1.WithPurchasable(in.Purchasable, false),
		delegatedstaking1.WithOnline(in.Online, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateDelegatedStakingResponse{
		Info: info,
	}, nil
}
