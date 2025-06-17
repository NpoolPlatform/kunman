package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/kunman/gateway/good/app/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
)

func (s *Server) AdminGetAppPowerRentals(ctx context.Context, in *npool.AdminGetAppPowerRentalsRequest) (*npool.AdminGetAppPowerRentalsResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(&in.TargetAppID, true),
		powerrental1.WithOffset(in.Offset),
		powerrental1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppPowerRentals",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppPowerRentalsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPowerRentals(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppPowerRentals",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppPowerRentalsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetAppPowerRentalsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
