package appfee

import (
	"context"

	appfee1 "github.com/NpoolPlatform/kunman/gateway/good/app/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
)

func (s *Server) AdminDeleteAppFee(ctx context.Context, in *npool.AdminDeleteAppFeeRequest) (*npool.AdminDeleteAppFeeResponse, error) {
	handler, err := appfee1.NewHandler(
		ctx,
		appfee1.WithID(&in.ID, true),
		appfee1.WithEntID(&in.EntID, true),
		appfee1.WithAppID(&in.TargetAppID, true),
		appfee1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteAppFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteAppFeeResponse{
		Info: info,
	}, nil
}
