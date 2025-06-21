//nolint:nolintlint,dupl
package appcoin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	appcoin1 "github.com/NpoolPlatform/kunman/gateway/chain/app/coin"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoins(ctx context.Context, in *npool.GetCoinsRequest) (*npool.GetCoinsResponse, error) {
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithAppID(&in.AppID, true),
		appcoin1.WithForPay(in.ForPay, false),
		appcoin1.WithOffset(in.GetOffset()),
		appcoin1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppCoins(ctx context.Context, in *npool.GetAppCoinsRequest) (*npool.GetAppCoinsResponse, error) {
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithAppID(&in.TargetAppID, true),
		appcoin1.WithOffset(in.GetOffset()),
		appcoin1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
