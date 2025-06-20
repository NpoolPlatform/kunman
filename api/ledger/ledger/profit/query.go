//nolint:dupl
package profit

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/profit"

	handler1 "github.com/NpoolPlatform/kunman/gateway/ledger/ledger/handler"
	profit1 "github.com/NpoolPlatform/kunman/gateway/ledger/ledger/profit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoinProfits(ctx context.Context, in *npool.GetCoinProfitsRequest) (*npool.GetCoinProfitsResponse, error) {
	handler, err := profit1.NewHandler(
		ctx,
		handler1.WithAppID(&in.AppID, true),
		handler1.WithUserID(&in.UserID, true),
		handler1.WithStartAt(in.StartAt, false),
		handler1.WithEndAt(in.EndAt, false),
		handler1.WithOffset(in.Offset),
		handler1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinProfits",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinProfitsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoinProfits(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinProfits",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinProfitsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinProfitsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

//nolint
func (s *Server) GetGoodProfits(ctx context.Context, in *npool.GetGoodProfitsRequest) (*npool.GetGoodProfitsResponse, error) {
	handler, err := profit1.NewHandler(
		ctx,
		handler1.WithAppID(&in.AppID, true),
		handler1.WithUserID(&in.UserID, true),
		handler1.WithStartAt(in.StartAt, false),
		handler1.WithEndAt(in.EndAt, false),
		handler1.WithOffset(in.Offset),
		handler1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodProfits",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodProfitsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetGoodProfits(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodProfits",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodProfitsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodProfitsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetMiningRewards(ctx context.Context, in *npool.GetMiningRewardsRequest) (*npool.GetMiningRewardsResponse, error) {
	handler, err := profit1.NewHandler(
		ctx,
		handler1.WithAppID(&in.AppID, true),
		handler1.WithUserID(&in.UserID, true),
		handler1.WithStartAt(in.StartAt, false),
		handler1.WithEndAt(in.EndAt, false),
		handler1.WithSimulateOnly(in.SimulateOnly, false),
		handler1.WithOffset(in.Offset),
		handler1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMiningRewards",
			"In", in,
			"Error", err,
		)
		return &npool.GetMiningRewardsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetMiningRewards(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMiningRewards",
			"In", in,
			"Error", err,
		)
		return &npool.GetMiningRewardsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMiningRewardsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
