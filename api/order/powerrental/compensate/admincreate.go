package compensate

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental/compensate"
	compensate1 "github.com/NpoolPlatform/kunman/gateway/order/powerrental/compensate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminCreateCompensate(ctx context.Context, in *npool.AdminCreateCompensateRequest) (*npool.AdminCreateCompensateResponse, error) {
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithGoodID(in.GoodID, false),
		compensate1.WithOrderID(in.OrderID, false),
		compensate1.WithCompensateFromID(&in.CompensateFromID, true),
		compensate1.WithCompensateType(&in.CompensateType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateCompensateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := handler.CreateCompensate(ctx); err != nil {
		logger.Sugar().Errorw(
			"AdminCreateCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateCompensateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminCreateCompensateResponse{}, nil
}
