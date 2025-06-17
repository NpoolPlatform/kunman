//nolint:dupl
package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/constraint"
)

func (s *Server) AdminUpdateTopMostConstraint(ctx context.Context, in *npool.AdminUpdateTopMostConstraintRequest) (*npool.AdminUpdateTopMostConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(&in.ID, true),
		constraint1.WithEntID(&in.EntID, true),
		constraint1.WithAppID(&in.TargetAppID, true),
		constraint1.WithTargetValue(in.TargetValue, false),
		constraint1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateTopMostConstraintResponse{
		Info: info,
	}, nil
}
