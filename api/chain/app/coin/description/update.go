package description

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	description1 "github.com/NpoolPlatform/kunman/gateway/chain/app/coin/description"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/coin/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoinDescription(ctx context.Context, in *npool.UpdateCoinDescriptionRequest) (*npool.UpdateCoinDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(&in.ID, true),
		description1.WithAppID(&in.AppID, true),
		description1.WithTitle(in.Title, false),
		description1.WithMessage(in.Message, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCoinDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinDescriptionResponse{
		Info: info,
	}, nil
}
