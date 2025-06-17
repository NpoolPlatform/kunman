//nolint:dupl
package like

import (
	"context"

	like1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/like"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/like"
)

func (s *Server) CreateLike(ctx context.Context, in *npool.CreateLikeRequest) (*npool.CreateLikeResponse, error) {
	handler, err := like1.NewHandler(
		ctx,
		like1.WithAppID(&in.AppID, true),
		like1.WithUserID(&in.UserID, true),
		like1.WithAppGoodID(&in.AppGoodID, true),
		like1.WithLike(&in.Like, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLike",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateLike(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLike",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateLikeResponse{
		Info: info,
	}, nil
}
