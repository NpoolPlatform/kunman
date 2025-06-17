//nolint:dupl
package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/constraint"
)

func (s *Server) UpdateTopMostConstraint(ctx context.Context, in *npool.UpdateTopMostConstraintRequest) (*npool.UpdateTopMostConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(&in.ID, true),
		constraint1.WithEntID(&in.EntID, true),
		constraint1.WithAppID(&in.AppID, true),
		constraint1.WithTargetValue(in.TargetValue, false),
		constraint1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostConstraintResponse{
		Info: info,
	}, nil
}
