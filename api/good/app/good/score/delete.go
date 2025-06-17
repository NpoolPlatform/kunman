//nolint:dupl
package score

import (
	"context"

	score1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/score"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"
)

func (s *Server) DeleteScore(ctx context.Context, in *npool.DeleteScoreRequest) (*npool.DeleteScoreResponse, error) {
	handler, err := score1.NewHandler(
		ctx,
		score1.WithID(&in.ID, true),
		score1.WithEntID(&in.EntID, true),
		score1.WithAppID(&in.AppID, true),
		score1.WithUserID(&in.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteScore",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteScore(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteScore",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteScoreResponse{
		Info: info,
	}, nil
}
