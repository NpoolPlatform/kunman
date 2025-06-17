package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/kunman/gateway/good/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/powerrental"
)

func (s *Server) AdminDeletePowerRental(ctx context.Context, in *npool.AdminDeletePowerRentalRequest) (*npool.AdminDeletePowerRentalResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(&in.ID, true),
		powerrental1.WithEntID(&in.EntID, true),
		powerrental1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeletePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeletePowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeletePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeletePowerRentalResponse{
		Info: info,
	}, nil
}
