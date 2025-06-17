package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/kunman/gateway/good/app/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
)

func (s *Server) UpdateAppPowerRental(ctx context.Context, in *npool.UpdateAppPowerRentalRequest) (*npool.UpdateAppPowerRentalResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(&in.ID, true),
		powerrental1.WithEntID(&in.EntID, true),
		powerrental1.WithAppID(&in.AppID, true),
		powerrental1.WithAppGoodID(&in.AppGoodID, true),

		powerrental1.WithPurchasable(in.Purchasable, false),
		powerrental1.WithEnableProductPage(in.EnableProductPage, false),
		powerrental1.WithProductPage(in.ProductPage, false),
		powerrental1.WithOnline(in.Online, false),
		powerrental1.WithVisible(in.Visible, false),
		powerrental1.WithName(in.Name, false),
		powerrental1.WithDisplayIndex(in.DisplayIndex, false),
		powerrental1.WithBanner(in.Banner, false),

		powerrental1.WithServiceStartAt(in.ServiceStartAt, false),
		powerrental1.WithCancelMode(in.CancelMode, false),
		powerrental1.WithCancelableBeforeStartSeconds(in.CancelableBeforeStartSeconds, false),
		powerrental1.WithEnableSetCommission(in.EnableSetCommission, false),
		powerrental1.WithMinOrderAmount(in.MinOrderAmount, false),
		powerrental1.WithMaxOrderAmount(in.MaxOrderAmount, false),
		powerrental1.WithMaxUserAmount(in.MaxUserAmount, false),
		powerrental1.WithMinOrderDurationSeconds(in.MinOrderDurationSeconds, false),
		powerrental1.WithMaxOrderDurationSeconds(in.MaxOrderDurationSeconds, false),
		powerrental1.WithUnitPrice(in.UnitPrice, false),
		powerrental1.WithSaleStartAt(in.SaleStartAt, false),
		powerrental1.WithSaleEndAt(in.SaleEndAt, false),
		powerrental1.WithSaleMode(in.SaleMode, false),
		powerrental1.WithFixedDuration(in.FixedDuration, false),
		powerrental1.WithPackageWithRequireds(in.PackageWithRequireds, false),
		powerrental1.WithStartMode(in.StartMode, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdatePowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateAppPowerRentalResponse{
		Info: info,
	}, nil
}
