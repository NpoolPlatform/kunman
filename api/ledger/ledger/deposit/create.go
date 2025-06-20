package deposit

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	deposit1 "github.com/NpoolPlatform/kunman/gateway/ledger/ledger/deposit"
	npool "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/deposit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppUserDeposit(ctx context.Context, in *npool.CreateAppUserDepositRequest) (
	resp *npool.CreateAppUserDepositResponse,
	err error,
) {
	handler, err := deposit1.NewHandler(
		ctx,
		deposit1.WithAppID(&in.AppID, true),
		deposit1.WithUserID(&in.UserID, true),
		deposit1.WithCoinTypeID(&in.CoinTypeID, true),
		deposit1.WithTargetAppID(&in.TargetAppID, true),
		deposit1.WithTargetUserID(&in.TargetUserID, true),
		deposit1.WithAmount(&in.Amount, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppUserDeposit",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppUserDepositResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDeposit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppUserDeposit",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppUserDepositResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAppUserDepositResponse{
		Info: info,
	}, nil
}
