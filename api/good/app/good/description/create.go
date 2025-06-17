package description

import (
	"context"

	description1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/description"
)

func (s *Server) CreateDescription(ctx context.Context, in *npool.CreateDescriptionRequest) (*npool.CreateDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithAppID(&in.AppID, true),
		description1.WithAppGoodID(&in.AppGoodID, true),
		description1.WithDescription(&in.Description, true),
		description1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDescriptionResponse{
		Info: info,
	}, nil
}
