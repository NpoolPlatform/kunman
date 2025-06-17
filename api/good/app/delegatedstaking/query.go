package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/kunman/gateway/good/app/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
)

func (s *Server) GetAppDelegatedStaking(ctx context.Context, in *npool.GetAppDelegatedStakingRequest) (*npool.GetAppDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppDelegatedStakingResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppDelegatedStakings(ctx context.Context, in *npool.GetAppDelegatedStakingsRequest) (*npool.GetAppDelegatedStakingsResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithAppID(&in.AppID, true),
		delegatedstaking1.WithOffset(in.Offset),
		delegatedstaking1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppDelegatedStakings",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppDelegatedStakingsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDelegatedStakings(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppDelegatedStakings",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppDelegatedStakingsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppDelegatedStakingsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
