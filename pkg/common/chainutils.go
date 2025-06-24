//nolint:dupl
package common

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	fiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	fiatmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
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
		appcoinmw.WithOffset(0),
		appcoinmw.WithLimit(int32(len(coinTypeIDs))),
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
		coinmw.WithOffset(0),
		coinmw.WithLimit(int32(len(coinTypeIDs))),
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

func GetFiats(ctx context.Context, fiatIDs []string) (map[string]*fiatmwpb.Fiat, error) {
	for _, fiatID := range fiatIDs {
		if _, err := uuid.Parse(fiatID); err != nil {
			return nil, err
		}
	}

	conds := &fiatmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: fiatIDs},
	}
	handler, err := fiatmw.NewHandler(
		ctx,
		fiatmw.WithConds(conds),
		fiatmw.WithOffset(0),
		fiatmw.WithLimit(int32(len(fiatIDs))),
	)
	if err != nil {
		return nil, err
	}

	fiats, _, err := handler.GetFiats(ctx)
	if err != nil {
		return nil, err
	}
	fiatMap := map[string]*fiatmwpb.Fiat{}
	for _, fiat := range fiats {
		fiatMap[fiat.EntID] = fiat
	}
	return fiatMap, nil
}
