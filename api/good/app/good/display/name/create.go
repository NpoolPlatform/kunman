package displayname

import (
	"context"

	displayname1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
)

func (s *Server) CreateDisplayName(ctx context.Context, in *npool.CreateDisplayNameRequest) (*npool.CreateDisplayNameResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithAppID(&in.AppID, true),
		displayname1.WithAppGoodID(&in.AppGoodID, true),
		displayname1.WithName(&in.Name, true),
		displayname1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDisplayName(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDisplayNameResponse{
		Info: info,
	}, nil
}
