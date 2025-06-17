package malfunction

import (
	"context"

	malfunction1 "github.com/NpoolPlatform/kunman/gateway/good/good/malfunction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
)

func (s *Server) AdminDeleteMalfunction(ctx context.Context, in *npool.AdminDeleteMalfunctionRequest) (*npool.AdminDeleteMalfunctionResponse, error) {
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithID(&in.ID, true),
		malfunction1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteMalfunction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteMalfunctionResponse{
		Info: info,
	}, nil
}
