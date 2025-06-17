package score

import (
	"context"

	score1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/score"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"
)

func (s *Server) GetMyScores(ctx context.Context, in *npool.GetMyScoresRequest) (*npool.GetMyScoresResponse, error) {
	handler, err := score1.NewHandler(
		ctx,
		score1.WithAppID(&in.AppID, true),
		score1.WithUserID(&in.UserID, true),
		score1.WithGoodID(in.GoodID, false),
		score1.WithAppGoodID(in.AppGoodID, false),
		score1.WithOffset(in.Offset),
		score1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyScores",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyScoresResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetScores(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyScores",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyScoresResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMyScoresResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetScores(ctx context.Context, in *npool.GetScoresRequest) (*npool.GetScoresResponse, error) {
	handler, err := score1.NewHandler(
		ctx,
		score1.WithAppID(&in.AppID, true),
		score1.WithUserID(in.TargetUserID, false),
		score1.WithGoodID(in.GoodID, false),
		score1.WithAppGoodID(in.AppGoodID, false),
		score1.WithOffset(in.Offset),
		score1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetScores",
			"In", in,
			"Error", err,
		)
		return &npool.GetScoresResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetScores(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetScores",
			"In", in,
			"Error", err,
		)
		return &npool.GetScoresResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetScoresResponse{
		Infos: infos,
		Total: total,
	}, nil
}
