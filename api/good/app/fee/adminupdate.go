package appfee

import (
	"context"

	appfee1 "github.com/NpoolPlatform/kunman/gateway/good/app/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
)

func (s *Server) AdminUpdateAppFee(ctx context.Context, in *npool.AdminUpdateAppFeeRequest) (*npool.AdminUpdateAppFeeResponse, error) {
	handler, err := appfee1.NewHandler(
		ctx,
		appfee1.WithID(&in.ID, true),
		appfee1.WithEntID(&in.EntID, true),
		appfee1.WithAppID(&in.TargetAppID, true),
		appfee1.WithAppGoodID(&in.AppGoodID, true),
		appfee1.WithProductPage(in.ProductPage, false),
		appfee1.WithName(in.Name, false),
		appfee1.WithBanner(in.Banner, false),
		appfee1.WithUnitValue(in.UnitValue, false),
		appfee1.WithMinOrderDurationSeconds(in.MinOrderDurationSeconds, false),
		appfee1.WithCancelMode(in.CancelMode, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateAppFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateAppFeeResponse{
		Info: info,
	}, nil
}
