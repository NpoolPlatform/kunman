package powerrental

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	powerrental1 "github.com/NpoolPlatform/kunman/gateway/order/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminGetPowerRentalOrders(ctx context.Context, in *npool.AdminGetPowerRentalOrdersRequest) (*npool.AdminGetPowerRentalOrdersResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(in.TargetAppID, false),
		powerrental1.WithGoodID(in.GoodID, false),
		powerrental1.WithOffset(in.Offset),
		powerrental1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetPowerRentalOrders",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetPowerRentalOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetPowerRentalOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetPowerRentalOrders",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetPowerRentalOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminGetPowerRentalOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
