package ledger

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger"

	ledger1 "github.com/NpoolPlatform/kunman/gateway/ledger/ledger"
	handler1 "github.com/NpoolPlatform/kunman/gateway/ledger/ledger/handler"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetLedgers(ctx context.Context, in *npool.GetLedgersRequest) (*npool.GetLedgersResponse, error) {
	handler, err := ledger1.NewHandler(
		ctx,
		handler1.WithAppID(&in.AppID, true),
		handler1.WithUserID(&in.UserID, true),
		handler1.WithOffset(in.Offset),
		handler1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLedgers",
			"In", in,
			"Error", err,
		)
		return &npool.GetLedgersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetLedgers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLedgers",
			"In", in,
			"Error", err,
		)
		return &npool.GetLedgersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLedgersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppLedgers(ctx context.Context, in *npool.GetAppLedgersRequest) (*npool.GetAppLedgersResponse, error) {
	handler, err := ledger1.NewHandler(
		ctx,
		handler1.WithAppID(&in.TargetAppID, true),
		handler1.WithOffset(in.Offset),
		handler1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppLedgers",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppLedgersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetLedgers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppLedgers",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppLedgersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppLedgersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
