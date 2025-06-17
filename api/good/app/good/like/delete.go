//nolint:dupl
package like

import (
	"context"

	like1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/like"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/like"
)

func (s *Server) DeleteLike(ctx context.Context, in *npool.DeleteLikeRequest) (*npool.DeleteLikeResponse, error) {
	handler, err := like1.NewHandler(
		ctx,
		like1.WithID(&in.ID, true),
		like1.WithEntID(&in.EntID, true),
		like1.WithAppID(&in.AppID, true),
		like1.WithUserID(&in.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLike",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteLike(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLike",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteLikeResponse{
		Info: info,
	}, nil
}
