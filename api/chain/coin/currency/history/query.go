//nolint:nolintlint,dupl
package currencyhistory

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	history1 "github.com/NpoolPlatform/kunman/gateway/chain/coin/currency/history"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin/currency/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCurrencies(ctx context.Context, in *npool.GetCurrenciesRequest) (*npool.GetCurrenciesResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithCoinNames(in.CoinNames),
		history1.WithCoinTypeIDs(in.CoinTypeIDs),
		history1.WithStartAt(in.StartAt),
		history1.WithEndAt(in.EndAt),
		history1.WithOffset(in.GetOffset()),
		history1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCurrencies",
			"In", in,
			"Error", err,
		)
		return &npool.GetCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCurrencies(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCurrencies",
			"In", in,
			"Error", err,
		)
		return &npool.GetCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrenciesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
