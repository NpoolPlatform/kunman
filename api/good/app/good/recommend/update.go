//nolint:dupl
package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/recommend"
)

func (s *Server) UpdateRecommend(ctx context.Context, in *npool.UpdateRecommendRequest) (*npool.UpdateRecommendResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithID(&in.ID, true),
		recommend1.WithEntID(&in.EntID, true),
		recommend1.WithAppID(&in.AppID, true),
		recommend1.WithRecommenderID(&in.UserID, true),
		recommend1.WithRecommendIndex(in.RecommendIndex, false),
		recommend1.WithMessage(in.Message, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateRecommend(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateRecommendResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateUserRecommend(ctx context.Context, in *npool.UpdateUserRecommendRequest) (*npool.UpdateUserRecommendResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithID(&in.ID, true),
		recommend1.WithEntID(&in.EntID, true),
		recommend1.WithAppID(&in.AppID, true),
		recommend1.WithRecommenderID(&in.TargetUserID, true),
		recommend1.WithHide(in.Hide, false),
		recommend1.WithHideReason(in.HideReason, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateRecommend(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateUserRecommendResponse{
		Info: info,
	}, nil
}
