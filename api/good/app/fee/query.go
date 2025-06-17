package appfee

import (
	"context"

	appfee1 "github.com/NpoolPlatform/kunman/gateway/good/app/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
)

func (s *Server) GetAppFee(ctx context.Context, in *npool.GetAppFeeRequest) (*npool.GetAppFeeResponse, error) {
	handler, err := appfee1.NewHandler(
		ctx,
		appfee1.WithAppID(&in.AppID, true),
		appfee1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetAppFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppFee",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppFeeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppFees(ctx context.Context, in *npool.GetAppFeesRequest) (*npool.GetAppFeesResponse, error) {
	handler, err := appfee1.NewHandler(
		ctx,
		appfee1.WithAppID(&in.AppID, true),
		appfee1.WithOffset(in.Offset),
		appfee1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppFees",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetAppFees(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppFees",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppFeesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
