package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"
)

func (s *Server) AdminUpdateTopMost(ctx context.Context, in *npool.AdminUpdateTopMostRequest) (*npool.AdminUpdateTopMostResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithID(&in.ID, true),
		topmost1.WithEntID(&in.EntID, true),
		topmost1.WithAppID(&in.TargetAppID, true),
		topmost1.WithTitle(in.Title, false),
		topmost1.WithMessage(in.Message, false),
		topmost1.WithTargetURL(in.TargetUrl, false),
		topmost1.WithStartAt(in.StartAt, false),
		topmost1.WithEndAt(in.EndAt, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateTopMostResponse{
		Info: info,
	}, nil
}
