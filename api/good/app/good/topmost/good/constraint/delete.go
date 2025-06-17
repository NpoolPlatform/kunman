package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/good/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/constraint"
)

func (s *Server) DeleteTopMostGoodConstraint(ctx context.Context, in *npool.DeleteTopMostGoodConstraintRequest) (*npool.DeleteTopMostGoodConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(&in.ID, true),
		constraint1.WithEntID(&in.EntID, true),
		constraint1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTopMostGoodConstraintResponse{
		Info: info,
	}, nil
}
