package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/kunman/gateway/good/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"
)

func (s *Server) AdminCreateDelegatedStaking(ctx context.Context, in *npool.AdminCreateDelegatedStakingRequest) (*npool.AdminCreateDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithContractCodeURL(&in.ContractCodeURL, true),
		delegatedstaking1.WithContractCodeBranch(&in.ContractCodeBranch, true),
		delegatedstaking1.WithCoinTypeID(&in.CoinTypeID, true),
		delegatedstaking1.WithGoodType(&in.GoodType, true),
		delegatedstaking1.WithName(&in.Name, true),
		delegatedstaking1.WithServiceStartAt(in.ServiceStartAt, true),
		delegatedstaking1.WithStartMode(&in.StartMode, true),
		delegatedstaking1.WithTestOnly(in.TestOnly, false),
		delegatedstaking1.WithBenefitIntervalHours(in.BenefitIntervalHours, true),
		delegatedstaking1.WithPurchasable(in.Purchasable, false),
		delegatedstaking1.WithOnline(in.Online, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateDelegatedStakingResponse{
		Info: info,
	}, nil
}
