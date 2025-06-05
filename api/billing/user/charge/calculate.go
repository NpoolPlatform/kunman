package charge

import (
	"context"

	charge1 "github.com/NpoolPlatform/kunman/gateway/billing/user/charge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gateway/v1/user/charge"
)

func (s *Server) CalculateCharge(ctx context.Context, in *npool.CalculateChargeRequest) (*npool.CalculateChargeResponse, error) {
	handler, err := charge1.NewHandler(
		ctx,
		charge1.WithAppID(&in.AppID, true),
		charge1.WithUserID(&in.UserID, true),
		charge1.WithPath(&in.Path, true),
		charge1.WithReqMsg(*in.ReqMsg, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CalculateCharge",
			"In", in,
			"Error", err,
		)
		return &npool.CalculateChargeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	allow, msg, err := handler.CalculateCharge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CalculateCharge",
			"In", in,
			"Error", err,
		)
		return &npool.CalculateChargeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CalculateChargeResponse{
		Allow:   allow,
		Message: &msg,
	}, nil
}
