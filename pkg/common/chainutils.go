//nolint:dupl
package common

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetAppCoins(ctx context.Context, appID string, coinTypeIDs []string) (map[string]*appcoinmwpb.Coin, error) {
	for _, coinTypeID := range coinTypeIDs {
		if _, err := uuid.Parse(coinTypeID); err != nil {
			return nil, err
		}
	}

	conds := &appcoinmwpb.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: appID},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}
	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithConds(conds),
	)
	if err != nil {
		return nil, err
	}

	coins, _, err := handler.GetCoins(ctx)
	if err != nil {
		return nil, err
	}
	coinMap := map[string]*appcoinmwpb.Coin{}
	for _, coin := range coins {
		coinMap[coin.CoinTypeID] = coin
	}
	return coinMap, nil
}

func GetCoins(ctx context.Context, coinTypeIDs []string) (map[string]*coinmwpb.Coin, error) {
	for _, coinTypeID := range coinTypeIDs {
		if _, err := uuid.Parse(coinTypeID); err != nil {
			return nil, err
		}
	}

	conds := &coinmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithConds(conds),
	)
	if err != nil {
		return nil, err
	}

	coins, _, err := handler.GetCoins(ctx)
	if err != nil {
		return nil, err
	}
	coinMap := map[string]*coinmwpb.Coin{}
	for _, coin := range coins {
		coinMap[coin.EntID] = coin
	}
	return coinMap, nil
}
