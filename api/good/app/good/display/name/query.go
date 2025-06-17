package displayname

import (
	"context"

	displayname1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
)

func (s *Server) GetDisplayNames(ctx context.Context, in *npool.GetDisplayNamesRequest) (*npool.GetDisplayNamesResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithAppID(&in.AppID, true),
		displayname1.WithAppGoodID(in.AppGoodID, false),
		displayname1.WithOffset(in.Offset),
		displayname1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayNames",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayNamesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDisplayNames(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayNames",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayNamesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDisplayNamesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
