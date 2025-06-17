package history

import (
	"context"

	history1 "github.com/NpoolPlatform/kunman/gateway/good/good/coin/reward/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin/reward/history"
)

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithGoodID(in.GoodID, false),
		history1.WithStartAt(in.StartAt, false),
		history1.WithEndAt(in.EndAt, false),
		history1.WithCoinTypeID(in.CoinTypeID, false),
		history1.WithOffset(in.Offset),
		history1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistories",
			"In", in,
			"Error", err,
		)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetHistories(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistories",
			"In", in,
			"Error", err,
		)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetHistoriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
