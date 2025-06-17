package location

import (
	"context"

	location1 "github.com/NpoolPlatform/kunman/gateway/good/vender/location"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/vender/location"
)

func (s *Server) AdminDeleteLocation(ctx context.Context, in *npool.AdminDeleteLocationRequest) (*npool.AdminDeleteLocationResponse, error) {
	handler, err := location1.NewHandler(
		ctx,
		location1.WithID(&in.ID, true),
		location1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteLocation",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteLocation(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteLocation",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteLocationResponse{
		Info: info,
	}, nil
}
