package fee

import (
	"context"

	fee1 "github.com/NpoolPlatform/kunman/gateway/good/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/fee"
)

func (s *Server) AdminDeleteFee(ctx context.Context, in *npool.AdminDeleteFeeRequest) (*npool.AdminDeleteFeeResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithID(&in.ID, true),
		fee1.WithEntID(&in.EntID, true),
		fee1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteFeeResponse{
		Info: info,
	}, nil
}
