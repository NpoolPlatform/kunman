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

func (s *Server) CreateScore(ctx context.Context, in *npool.CreateScoreRequest) (*npool.CreateScoreResponse, error) {
	handler, err := score1.NewHandler(
		ctx,
		score1.WithAppID(&in.AppID, true),
		score1.WithUserID(&in.UserID, true),
		score1.WithAppGoodID(&in.AppGoodID, true),
		score1.WithScore(&in.Score, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateScore",
			"In", in,
			"Error", err,
		)
		return &npool.CreateScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateScore(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateScore",
			"In", in,
			"Error", err,
		)
		return &npool.CreateScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateScoreResponse{
		Info: info,
	}, nil
}
