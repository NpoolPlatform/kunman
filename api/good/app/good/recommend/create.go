package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/recommend"
)

func (s *Server) CreateRecommend(ctx context.Context, in *npool.CreateRecommendRequest) (*npool.CreateRecommendResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithAppID(&in.AppID, true),
		recommend1.WithRecommenderID(&in.UserID, true),
		recommend1.WithAppGoodID(&in.AppGoodID, true),
		recommend1.WithRecommendIndex(&in.RecommendIndex, true),
		recommend1.WithMessage(&in.Message, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateRecommend(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateRecommendResponse{
		Info: info,
	}, nil
}
