//nolint:dupl
package default1

import (
	"context"

	default1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/default"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/default"
)

func (s *Server) CreateDefault(ctx context.Context, in *npool.CreateDefaultRequest) (*npool.CreateDefaultResponse, error) {
	handler, err := default1.NewHandler(
		ctx,
		default1.WithAppID(&in.AppID, true),
		default1.WithCoinTypeID(&in.CoinTypeID, true),
		default1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDefault(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDefaultResponse{
		Info: info,
	}, nil
}
