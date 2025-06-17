package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/kunman/gateway/good/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"
)

func (s *Server) GetDelegatedStaking(ctx context.Context, in *npool.GetDelegatedStakingRequest) (*npool.GetDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.GetDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.GetDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDelegatedStakingResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDelegatedStakings(ctx context.Context, in *npool.GetDelegatedStakingsRequest) (*npool.GetDelegatedStakingsResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithOffset(in.Offset),
		delegatedstaking1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDelegatedStakings",
			"In", in,
			"Error", err,
		)
		return &npool.GetDelegatedStakingsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDelegatedStakings(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDelegatedStakings",
			"In", in,
			"Error", err,
		)
		return &npool.GetDelegatedStakingsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDelegatedStakingsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
