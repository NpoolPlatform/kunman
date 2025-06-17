package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/kunman/gateway/good/app/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
)

func (s *Server) AdminDeleteAppPowerRental(ctx context.Context, in *npool.AdminDeleteAppPowerRentalRequest) (*npool.AdminDeleteAppPowerRentalResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(&in.ID, true),
		powerrental1.WithEntID(&in.EntID, true),
		powerrental1.WithAppID(&in.TargetAppID, true),
		powerrental1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeletePowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteAppPowerRentalResponse{
		Info: info,
	}, nil
}
