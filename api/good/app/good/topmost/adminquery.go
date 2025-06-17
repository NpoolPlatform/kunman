package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"
)

func (s *Server) AdminGetTopMosts(ctx context.Context, in *npool.AdminGetTopMostsRequest) (*npool.AdminGetTopMostsResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithAppID(&in.TargetAppID, true),
		topmost1.WithOffset(in.Offset),
		topmost1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetTopMosts",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetTopMostsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetTopMosts(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetTopMosts",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetTopMostsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetTopMostsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
