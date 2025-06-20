package order

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"

	order1 "github.com/NpoolPlatform/kunman/gateway/order/order"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AdminGetOrders(ctx context.Context, in *npool.AdminGetOrdersRequest) (*npool.AdminGetOrdersResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithAppID(in.TargetAppID, false),
		order1.WithOffset(in.GetOffset()),
		order1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetOrders",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetOrders",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminGetOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
