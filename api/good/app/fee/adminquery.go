package appfee

import (
	"context"

	appfee1 "github.com/NpoolPlatform/kunman/gateway/good/app/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
)

func (s *Server) AdminGetAppFees(ctx context.Context, in *npool.AdminGetAppFeesRequest) (*npool.AdminGetAppFeesResponse, error) {
	handler, err := appfee1.NewHandler(
		ctx,
		appfee1.WithAppID(&in.TargetAppID, true),
		appfee1.WithOffset(in.Offset),
		appfee1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppFees",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetAppFees(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppFees",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetAppFeesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
