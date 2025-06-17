package compensate

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/compensate"
	compensate1 "github.com/NpoolPlatform/kunman/gateway/order/compensate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminGetCompensates(ctx context.Context, in *npool.AdminGetCompensatesRequest) (*npool.AdminGetCompensatesResponse, error) {
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithAppID(in.TargetAppID, false),
		compensate1.WithGoodID(in.GoodID, false),
		compensate1.WithOffset(in.Offset),
		compensate1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetCompensates",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetCompensatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCompensates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetCompensates",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetCompensatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminGetCompensatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
