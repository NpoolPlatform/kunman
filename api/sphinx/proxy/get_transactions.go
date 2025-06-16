package api

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/price"
	pconst "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/message/const"
	pluginpb "github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	proxypb "github.com/NpoolPlatform/kunman/message/sphinx/proxy"
	"github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/crud"
	ent "github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db/ent/generated"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetTransactions ..
func (s *Server) GetTransactions(ctx context.Context, in *proxypb.GetTransactionsRequest) (out *proxypb.GetTransactionsResponse, err error) {
	pluginInfo := pconst.GetPluginInfo(ctx)
	ctx, cancel := context.WithTimeout(ctx, constant.GrpcTimeout)
	defer cancel()

	// TODO: debug plugin env info include(position and ip)
	if in.GetTransactionState() == proxypb.TransactionState_TransactionStateUnKnow {
		return &proxypb.GetTransactionsResponse{},
			status.Error(codes.InvalidArgument, "Invalid argument TransactionState must not empty")
	}

	if _, ok := proxypb.TransactionState_name[int32(in.GetTransactionState())]; !ok {
		return &proxypb.GetTransactionsResponse{},
			status.Error(codes.InvalidArgument, "Invalid argument TransactionState not support")
	}

	if in.GetCoinType() != sphinxplugin.CoinType_CoinTypeUnKnow {
		if _, ok := sphinxplugin.CoinType_name[int32(in.GetCoinType())]; !ok {
			return &proxypb.GetTransactionsResponse{},
				status.Error(codes.InvalidArgument, "Invalid argument CoinType not support")
		}
	}

	if in.GetENV() != "" && in.GetENV() != "main" && in.GetENV() != "test" {
		return &proxypb.GetTransactionsResponse{},
			status.Error(codes.InvalidArgument, "Invalid argument ENV only support main|test")
	}

	transInfos, err := crud.GetTransactions(ctx, crud.GetTransactionsParam{
		CoinType:         in.GetCoinType(),
		TransactionState: in.GetTransactionState(),
	})
	if ent.IsNotFound(err) {
		logger.Sugar().Info("GetTransactions no wait transaction")
		return &proxypb.GetTransactionsResponse{}, nil
	}

	if err != nil {
		logger.Sugar().Errorf("GetTransactions call GetTransactions error: %v", err)
		return &proxypb.GetTransactionsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	infos := make([]*proxypb.TransactionInfo, 0, len(transInfos))
	for _, info := range transInfos {
		infos = append(infos, &proxypb.TransactionInfo{
			TransactionID:    info.TransactionID,
			TransactionState: proxypb.TransactionState(info.State),
			Name:             info.Name,
			Amount:           price.DBPriceToVisualPrice(info.Amount),
			Payload:          info.Payload,
			From:             info.From,
			To:               info.To,
			Memo:             info.Memo,
		})
	}

	if len(infos) > 0 {
		logger.Sugar().Debugf(
			"%v get tasks,CoinType:%v CoinTransactionState:%v Rows:%v",
			pluginInfo,
			in.GetCoinType(),
			in.GetTransactionState(),
			len(infos),
		)
	}

	return &proxypb.GetTransactionsResponse{
		Infos: infos,
		Total: 0, // total no need
	}, nil
}
