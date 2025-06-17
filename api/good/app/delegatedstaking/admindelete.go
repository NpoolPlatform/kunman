package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/kunman/gateway/good/app/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
)

func (s *Server) AdminDeleteAppDelegatedStaking(ctx context.Context, in *npool.AdminDeleteAppDelegatedStakingRequest) (*npool.AdminDeleteAppDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithID(&in.ID, true),
		delegatedstaking1.WithEntID(&in.EntID, true),
		delegatedstaking1.WithAppID(&in.TargetAppID, true),
		delegatedstaking1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteAppDelegatedStakingResponse{
		Info: info,
	}, nil
}
