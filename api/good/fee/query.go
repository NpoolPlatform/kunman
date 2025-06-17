package fee

import (
	"context"

	fee1 "github.com/NpoolPlatform/kunman/gateway/good/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/fee"
)

func (s *Server) GetFee(ctx context.Context, in *npool.GetFeeRequest) (*npool.GetFeeResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFee",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFee",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFeeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFees(ctx context.Context, in *npool.GetFeesRequest) (*npool.GetFeesResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithOffset(in.Offset),
		fee1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFees",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetFees(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFees",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFeesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
