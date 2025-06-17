package outofgas

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental/outofgas"
	outofgas1 "github.com/NpoolPlatform/kunman/gateway/order/powerrental/outofgas"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminDeleteOutOfGas(ctx context.Context, in *npool.AdminDeleteOutOfGasRequest) (*npool.AdminDeleteOutOfGasResponse, error) {
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithAppID(&in.TargetAppID, true),
		outofgas1.WithUserID(&in.TargetUserID, true),
		outofgas1.WithID(&in.ID, true),
		outofgas1.WithEntID(&in.EntID, true),
		outofgas1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteOutOfGas",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteOutOfGasResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteOutOfGas(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteOutOfGas",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteOutOfGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminDeleteOutOfGasResponse{
		Info: info,
	}, nil
}
