package description

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	description1 "github.com/NpoolPlatform/kunman/gateway/chain/app/coin/description"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/coin/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// nolint
func (s *Server) CreateCoinDescription(ctx context.Context, in *npool.CreateCoinDescriptionRequest) (*npool.CreateCoinDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithAppID(&in.AppID, true),
		description1.WithCoinTypeID(&in.CoinTypeID, true),
		description1.WithUsedFor(&in.UsedFor, true),
		description1.WithTitle(&in.Title, true),
		description1.WithMessage(&in.Message, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoinDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinDescriptionResponse{
		Info: info,
	}, nil
}

// nolint
func (s *Server) CreateAppCoinDescription(ctx context.Context, in *npool.CreateAppCoinDescriptionRequest) (*npool.CreateAppCoinDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithAppID(&in.TargetAppID, true),
		description1.WithCoinTypeID(&in.CoinTypeID, true),
		description1.WithUsedFor(&in.UsedFor, true),
		description1.WithTitle(&in.Title, true),
		description1.WithMessage(&in.Message, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoinDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppCoinDescriptionResponse{
		Info: info,
	}, nil
}
