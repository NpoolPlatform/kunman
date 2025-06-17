package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/kunman/gateway/good/app/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
)

func (s *Server) GetAppPowerRental(ctx context.Context, in *npool.GetAppPowerRentalRequest) (*npool.GetAppPowerRentalResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetPowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppPowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppPowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppPowerRentalResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppPowerRentals(ctx context.Context, in *npool.GetAppPowerRentalsRequest) (*npool.GetAppPowerRentalsResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(&in.AppID, true),
		powerrental1.WithOffset(in.Offset),
		powerrental1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppPowerRentals",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppPowerRentalsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPowerRentals(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppPowerRentals",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppPowerRentalsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppPowerRentalsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
