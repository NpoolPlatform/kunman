package powerrental

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	powerrental1 "github.com/NpoolPlatform/kunman/gateway/order/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) GetPowerRentalOrder(ctx context.Context, in *npool.GetPowerRentalOrderRequest) (*npool.GetPowerRentalOrderResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(&in.AppID, true),
		powerrental1.WithUserID(&in.UserID, true),
		powerrental1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.GetPowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetPowerRentalOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.GetPowerRentalOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPowerRentalOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) GetPowerRentalOrders(ctx context.Context, in *npool.GetPowerRentalOrdersRequest) (*npool.GetPowerRentalOrdersResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(&in.AppID, true),
		powerrental1.WithUserID(in.TargetUserID, false),
		powerrental1.WithAppGoodID(in.AppGoodID, false),
		powerrental1.WithOffset(in.Offset),
		powerrental1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRentalOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetPowerRentalOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetPowerRentalOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRentalOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetPowerRentalOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPowerRentalOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetMyPowerRentalOrders(ctx context.Context, in *npool.GetMyPowerRentalOrdersRequest) (*npool.GetMyPowerRentalOrdersResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(&in.AppID, true),
		powerrental1.WithUserID(&in.UserID, true),
		powerrental1.WithAppGoodID(in.AppGoodID, false),
		powerrental1.WithOffset(in.Offset),
		powerrental1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyPowerRentalOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyPowerRentalOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetPowerRentalOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyPowerRentalOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyPowerRentalOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMyPowerRentalOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
