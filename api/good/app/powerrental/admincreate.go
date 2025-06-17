package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/kunman/gateway/good/app/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
)

func (s *Server) AdminCreateAppPowerRental(ctx context.Context, in *npool.AdminCreateAppPowerRentalRequest) (*npool.AdminCreateAppPowerRentalResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(&in.TargetAppID, true),
		powerrental1.WithGoodID(&in.GoodID, true),

		powerrental1.WithPurchasable(in.Purchasable, false),
		powerrental1.WithEnableProductPage(in.EnableProductPage, false),
		powerrental1.WithProductPage(in.ProductPage, false),
		powerrental1.WithOnline(in.Online, false),
		powerrental1.WithVisible(in.Visible, false),
		powerrental1.WithName(&in.Name, true),
		powerrental1.WithDisplayIndex(in.DisplayIndex, false),
		powerrental1.WithBanner(in.Banner, false),

		powerrental1.WithServiceStartAt(&in.ServiceStartAt, true),
		powerrental1.WithCancelMode(in.CancelMode, false),
		powerrental1.WithCancelableBeforeStartSeconds(in.CancelableBeforeStartSeconds, false),
		powerrental1.WithEnableSetCommission(in.EnableSetCommission, false),
		powerrental1.WithMinOrderAmount(&in.MinOrderAmount, true),
		powerrental1.WithMaxOrderAmount(&in.MaxOrderAmount, true),
		powerrental1.WithMaxUserAmount(in.MaxUserAmount, false),
		powerrental1.WithMinOrderDurationSeconds(&in.MinOrderDurationSeconds, true),
		powerrental1.WithMaxOrderDurationSeconds(in.MaxOrderDurationSeconds, false),
		powerrental1.WithUnitPrice(&in.UnitPrice, true),
		powerrental1.WithSaleStartAt(&in.SaleStartAt, true),
		powerrental1.WithSaleEndAt(&in.SaleEndAt, true),
		powerrental1.WithSaleMode(&in.SaleMode, true),
		powerrental1.WithFixedDuration(in.FixedDuration, false),
		powerrental1.WithPackageWithRequireds(in.PackageWithRequireds, false),
		powerrental1.WithStartMode(in.StartMode, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreatePowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateAppPowerRentalResponse{
		Info: info,
	}, nil
}
