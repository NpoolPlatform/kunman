//nolint:dupl
package required

import (
	"context"

	required1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/required"
)

func (s *Server) AdminCreateRequired(ctx context.Context, in *npool.AdminCreateRequiredRequest) (*npool.AdminCreateRequiredResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithAppID(&in.TargetAppID, true),
		required1.WithMainAppGoodID(&in.MainAppGoodID, true),
		required1.WithRequiredAppGoodID(&in.RequiredAppGoodID, true),
		required1.WithMust(in.Must, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateRequiredResponse{
		Info: info,
	}, nil
}
