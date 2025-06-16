package api

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/price"
	proxypb "github.com/NpoolPlatform/kunman/message/sphinx/proxy"
	"github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/crud"
	ent "github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db/ent/generated"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTransaction(ctx context.Context, in *proxypb.GetTransactionRequest) (out *proxypb.GetTransactionResponse, err error) {
	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("GetTransaction TransactionID empty")
		return &proxypb.GetTransactionResponse{}, status.Error(codes.InvalidArgument, "TransactionID empty")
	}

	ctx, cancel := context.WithTimeout(ctx, constant.GrpcTimeout)
	defer cancel()

	transInfo, err := crud.GetTransaction(ctx, in.GetTransactionID())
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("GetTransaction TransactionID: %v not found", in.GetTransactionID())
		return &proxypb.GetTransactionResponse{}, status.Errorf(codes.NotFound, "TransactionID: %v not found", in.GetTransactionID())
	}

	if err != nil {
		logger.Sugar().Errorf("GetTransaction call GetTransaction error: %v", err)
		return &proxypb.GetTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &proxypb.GetTransactionResponse{
		Info: &proxypb.TransactionInfo{
			TransactionID: transInfo.TransactionID,
			Name:          transInfo.Name,
			Amount:        price.DBPriceToVisualPrice(transInfo.Amount),
			Payload:       transInfo.Payload,
			From:          transInfo.From,
			To:            transInfo.To,
			Memo:          transInfo.Memo,

			ExitCode:         transInfo.ExitCode,
			CID:              transInfo.Cid,
			TransactionState: proxypb.TransactionState(transInfo.State),

			CreatedAt: transInfo.CreatedAt,
			UpdatedAt: transInfo.UpdatedAt,
		},
	}, nil
}
