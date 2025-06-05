package exchange

import (
	"context"

	exchange1 "github.com/NpoolPlatform/kunman/gateway/billing/credit/exchange"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/credit/exchange"
)

func (s *Server) GetExchange(ctx context.Context, in *npool.GetExchangeRequest) (*npool.GetExchangeResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchange",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetExchange(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchange",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetExchangeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetExchanges(ctx context.Context, in *npool.GetExchangesRequest) (*npool.GetExchangesResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithOffset(in.Offset),
		exchange1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchanges",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetExchanges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchanges",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetExchangesResponse{
		Infos: infos,
	}, nil
}

func (s *Server) GetExchangesCount(ctx context.Context, in *npool.GetExchangesCountRequest) (*npool.GetExchangesCountResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchangesCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangesCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.GetExchangesCount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetExchangesCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetExchangesCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetExchangesCountResponse{
		Total: total,
	}, nil
}
