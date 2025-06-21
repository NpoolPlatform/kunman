package tx

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	tx1 "github.com/NpoolPlatform/kunman/gateway/chain/tx"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/tx"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateTx(ctx context.Context, in *npool.UpdateTxRequest) (*npool.UpdateTxResponse, error) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithID(&in.ID, true),
		tx1.WithEntID(&in.EntID, true),
		tx1.WithState(in.State, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTx",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateTx(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTx",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTxResponse{
		Info: info,
	}, nil
}
