package score

import (
	"context"

	score1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/score"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"
)

func (s *Server) AdminGetScores(ctx context.Context, in *npool.AdminGetScoresRequest) (*npool.AdminGetScoresResponse, error) {
	handler, err := score1.NewHandler(
		ctx,
		score1.WithAppID(&in.TargetAppID, true),
		score1.WithGoodID(in.GoodID, false),
		score1.WithAppGoodID(in.AppGoodID, false),
		score1.WithOffset(in.Offset),
		score1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetScores",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetScoresResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetScores(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetScores",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetScoresResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetScoresResponse{
		Infos: infos,
		Total: total,
	}, nil
}
