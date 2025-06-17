package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/constraint"
)

func (s *Server) AdminGetTopMostConstraints(ctx context.Context, in *npool.AdminGetTopMostConstraintsRequest) (*npool.AdminGetTopMostConstraintsResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithAppID(&in.TargetAppID, true),
		constraint1.WithOffset(in.Offset),
		constraint1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetTopMostConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetTopMostConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetConstraints(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetTopMostConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetTopMostConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetTopMostConstraintsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
