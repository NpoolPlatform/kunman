package statement

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/statement"

	handler1 "github.com/NpoolPlatform/kunman/gateway/ledger/ledger/handler"
	statement1 "github.com/NpoolPlatform/kunman/gateway/ledger/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetStatements(ctx context.Context, in *npool.GetStatementsRequest) (*npool.GetStatementsResponse, error) {
	handler, err := statement1.NewHandler(
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
			"GetStatements",
			"In", in,
			"Error", err,
		)
		return &npool.GetStatementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetStatements",
			"In", in,
			"Error", err,
		)
		return &npool.GetStatementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStatementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppStatements(ctx context.Context, in *npool.GetAppStatementsRequest) (*npool.GetAppStatementsResponse, error) {
	handler, err := statement1.NewHandler(
		ctx,
		handler1.WithAppID(&in.TargetAppID, true),
		handler1.WithOffset(in.Offset),
		handler1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppStatements",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppStatementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppStatements",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppStatementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppStatementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
