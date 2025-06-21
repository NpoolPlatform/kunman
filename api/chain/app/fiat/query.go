//nolint:nolintlint,dupl
package appfiat

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	appfiat1 "github.com/NpoolPlatform/kunman/gateway/chain/app/fiat"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetFiats(ctx context.Context, in *npool.GetFiatsRequest) (*npool.GetFiatsResponse, error) {
	handler, err := appfiat1.NewHandler(
		ctx,
		appfiat1.WithAppID(&in.AppID, true),
		appfiat1.WithOffset(in.GetOffset()),
		appfiat1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetFiatsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFiats(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetFiatsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppFiats(ctx context.Context, in *npool.GetAppFiatsRequest) (*npool.GetAppFiatsResponse, error) {
	handler, err := appfiat1.NewHandler(
		ctx,
		appfiat1.WithAppID(&in.TargetAppID, true),
		appfiat1.WithOffset(in.GetOffset()),
		appfiat1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppFiatsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFiats(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppFiatsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppFiatsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
