package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/constraint"
)

func (s *Server) GetTopMostConstraints(ctx context.Context, in *npool.GetTopMostConstraintsRequest) (*npool.GetTopMostConstraintsResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithAppID(&in.AppID, true),
		constraint1.WithOffset(in.Offset),
		constraint1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetConstraints(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostConstraintsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
