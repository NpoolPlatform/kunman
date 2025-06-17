package label

import (
	"context"

	label1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/label"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/label"
)

func (s *Server) AdminCreateLabel(ctx context.Context, in *npool.AdminCreateLabelRequest) (*npool.AdminCreateLabelResponse, error) {
	handler, err := label1.NewHandler(
		ctx,
		label1.WithAppID(&in.TargetAppID, true),
		label1.WithAppGoodID(&in.AppGoodID, true),
		label1.WithIcon(in.Icon, false),
		label1.WithIconBgColor(in.IconBgColor, false),
		label1.WithLabel(&in.Label, true),
		label1.WithLabelBgColor(in.LabelBgColor, false),
		label1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateLabel(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateLabelResponse{
		Info: info,
	}, nil
}
