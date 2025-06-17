package malfunction

import (
	"context"

	malfunction1 "github.com/NpoolPlatform/kunman/gateway/good/good/malfunction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
)

func (s *Server) AdminUpdateMalfunction(ctx context.Context, in *npool.AdminUpdateMalfunctionRequest) (*npool.AdminUpdateMalfunctionResponse, error) {
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithID(&in.ID, true),
		malfunction1.WithEntID(&in.EntID, true),
		malfunction1.WithTitle(in.Title, false),
		malfunction1.WithMessage(in.Message, false),
		malfunction1.WithStartAt(in.StartAt, false),
		malfunction1.WithDurationSeconds(in.DurationSeconds, false),
		malfunction1.WithCompensateSeconds(in.CompensateSeconds, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateMalfunction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateMalfunctionResponse{
		Info: info,
	}, nil
}
