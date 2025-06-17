package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/constraint"
)

func (s *Server) AdminDeleteTopMostConstraint(ctx context.Context, in *npool.AdminDeleteTopMostConstraintRequest) (*npool.AdminDeleteTopMostConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(&in.ID, true),
		constraint1.WithEntID(&in.EntID, true),
		constraint1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteTopMostConstraintResponse{
		Info: info,
	}, nil
}
