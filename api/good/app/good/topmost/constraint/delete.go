package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/constraint"
)

func (s *Server) DeleteTopMostConstraint(ctx context.Context, in *npool.DeleteTopMostConstraintRequest) (*npool.DeleteTopMostConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(&in.ID, true),
		constraint1.WithEntID(&in.EntID, true),
		constraint1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTopMostConstraintResponse{
		Info: info,
	}, nil
}
