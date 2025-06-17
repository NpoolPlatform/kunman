package displayname

import (
	"context"

	displayname1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
)

func (s *Server) AdminGetDisplayNames(ctx context.Context, in *npool.AdminGetDisplayNamesRequest) (*npool.AdminGetDisplayNamesResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithAppID(&in.TargetAppID, true),
		displayname1.WithOffset(in.Offset),
		displayname1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetDisplayNames",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetDisplayNamesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDisplayNames(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetDisplayNames",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetDisplayNamesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetDisplayNamesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
