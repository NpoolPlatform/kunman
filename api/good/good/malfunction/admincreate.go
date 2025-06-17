package malfunction

import (
	"context"

	malfunction1 "github.com/NpoolPlatform/kunman/gateway/good/good/malfunction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
)

func (s *Server) AdminCreateMalfunction(ctx context.Context, in *npool.AdminCreateMalfunctionRequest) (*npool.AdminCreateMalfunctionResponse, error) {
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithGoodID(&in.GoodID, true),
		malfunction1.WithTitle(&in.Title, true),
		malfunction1.WithMessage(&in.Message, true),
		malfunction1.WithStartAt(&in.StartAt, true),
		malfunction1.WithDurationSeconds(in.DurationSeconds, false),
		malfunction1.WithCompensateSeconds(in.CompensateSeconds, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateMalfunction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateMalfunctionResponse{
		Info: info,
	}, nil
}
