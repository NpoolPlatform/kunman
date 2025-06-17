//nolint:dupl
package default1

import (
	"context"

	default1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/default"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/default"
)

func (s *Server) DeleteDefault(ctx context.Context, in *npool.DeleteDefaultRequest) (*npool.DeleteDefaultResponse, error) {
	handler, err := default1.NewHandler(
		ctx,
		default1.WithID(&in.ID, true),
		default1.WithEntID(&in.EntID, true),
		default1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDefault",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDefault(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDefault",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDefaultResponse{
		Info: info,
	}, nil
}
