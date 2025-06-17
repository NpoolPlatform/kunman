package outofgas

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/outofgas"
	outofgas1 "github.com/NpoolPlatform/kunman/gateway/order/outofgas"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) GetOutOfGases(ctx context.Context, in *npool.GetOutOfGasesRequest) (*npool.GetOutOfGasesResponse, error) {
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithAppID(&in.AppID, true),
		outofgas1.WithUserID(in.TargetUserID, false),
		outofgas1.WithAppGoodID(in.AppGoodID, false),
		outofgas1.WithOffset(in.Offset),
		outofgas1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOutOfGases",
			"In", in,
			"Error", err,
		)
		return &npool.GetOutOfGasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOutOfGases(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOutOfGases",
			"In", in,
			"Error", err,
		)
		return &npool.GetOutOfGasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOutOfGasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetMyOutOfGases(ctx context.Context, in *npool.GetMyOutOfGasesRequest) (*npool.GetMyOutOfGasesResponse, error) {
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithAppID(&in.AppID, true),
		outofgas1.WithUserID(&in.UserID, true),
		outofgas1.WithAppGoodID(in.OrderID, false),
		outofgas1.WithOffset(in.Offset),
		outofgas1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyOutOfGases",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyOutOfGasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOutOfGases(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyOutOfGases",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyOutOfGasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMyOutOfGasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
