package appfee

import (
	"context"

	appfee1 "github.com/NpoolPlatform/kunman/gateway/good/app/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
)

func (s *Server) AdminCreateAppFee(ctx context.Context, in *npool.AdminCreateAppFeeRequest) (*npool.AdminCreateAppFeeResponse, error) {
	handler, err := appfee1.NewHandler(
		ctx,
		appfee1.WithAppID(&in.TargetAppID, true),
		appfee1.WithGoodID(&in.GoodID, true),
		appfee1.WithProductPage(in.ProductPage, false),
		appfee1.WithName(&in.Name, true),
		appfee1.WithBanner(in.Banner, false),
		appfee1.WithUnitValue(&in.UnitValue, true),
		appfee1.WithMinOrderDurationSeconds(&in.MinOrderDurationSeconds, true),
		appfee1.WithCancelMode(in.CancelMode, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateAppFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateAppFeeResponse{
		Info: info,
	}, nil
}
