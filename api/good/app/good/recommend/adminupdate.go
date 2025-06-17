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

func (s *Server) AdminUpdateRecommend(ctx context.Context, in *npool.AdminUpdateRecommendRequest) (*npool.AdminUpdateRecommendResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithID(&in.ID, true),
		recommend1.WithEntID(&in.EntID, true),
		recommend1.WithAppID(&in.TargetAppID, true),
		recommend1.WithRecommenderID(&in.TargetUserID, true),
		recommend1.WithHide(in.Hide, false),
		recommend1.WithHideReason(in.HideReason, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateRecommend(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateRecommendResponse{
		Info: info,
	}, nil
}
