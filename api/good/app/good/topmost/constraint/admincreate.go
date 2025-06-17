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

func (s *Server) AdminCreateTopMostConstraint(ctx context.Context, in *npool.AdminCreateTopMostConstraintRequest) (*npool.AdminCreateTopMostConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithAppID(&in.TargetAppID, true),
		constraint1.WithTopMostID(&in.TopMostID, true),
		constraint1.WithConstraint(&in.Constraint, true),
		constraint1.WithTargetValue(in.TargetValue, false),
		constraint1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateTopMostConstraintResponse{
		Info: info,
	}, nil
}
