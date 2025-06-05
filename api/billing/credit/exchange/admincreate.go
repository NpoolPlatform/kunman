package exchange

import (
	"context"

	exchange1 "github.com/NpoolPlatform/kunman/gateway/billing/credit/exchange"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/credit/exchange"
)

func (s *Server) AdminCreateExchange(ctx context.Context, in *npool.AdminCreateExchangeRequest) (*npool.AdminCreateExchangeResponse, error) {
	handler, err := exchange1.NewHandler(
		ctx,
		exchange1.WithAppID(&in.TargetAppID, true),
		exchange1.WithUsageType(&in.UsageType, true),
		exchange1.WithCredit(&in.Credit, true),
		exchange1.WithExchangeThreshold(&in.ExchangeThreshold, true),
		exchange1.WithPath(&in.Path, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateExchange",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateExchange(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateExchange",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateExchangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateExchangeResponse{
		Info: info,
	}, nil
}
