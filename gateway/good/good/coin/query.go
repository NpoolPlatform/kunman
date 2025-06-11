package goodcoin

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	goodcoinmw "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	goodCoins []*goodcoinmwpb.GoodCoin
	coins     map[string]*coinmwpb.Coin
	infos     []*npool.GoodCoin
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.coins, err = goodgwcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, goodCoin := range h.goodCoins {
			coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, goodCoin := range h.goodCoins {
		coin, ok := h.coins[goodCoin.CoinTypeID]
		if !ok {
			continue
		}
		h.infos = append(h.infos, &npool.GoodCoin{
			ID:         goodCoin.ID,
			EntID:      goodCoin.EntID,
			GoodID:     goodCoin.GoodID,
			GoodName:   goodCoin.GoodName,
			GoodType:   goodCoin.GoodType,
			CoinTypeID: goodCoin.CoinTypeID,
			CoinName:   coin.Name,
			CoinUnit:   coin.Unit,
			CoinENV:    coin.ENV,
			CoinLogo:   coin.Logo,
			Main:       goodCoin.Main,
			Index:      goodCoin.Index,
			CreatedAt:  goodCoin.CreatedAt,
			UpdatedAt:  goodCoin.UpdatedAt,
		})
	}
}

func (h *Handler) GetGoodCoin(ctx context.Context) (*npool.GoodCoin, error) {
	coinHandler, err := goodcoinmw.NewHandler(
		ctx,
		goodcoinmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := coinHandler.GetGoodCoin(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid goodcoin")
	}
	handler := &queryHandler{
		Handler:   h,
		goodCoins: []*goodcoinmwpb.GoodCoin{info},
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, wlog.Errorf("invalid goodcoin")
	}
	return handler.infos[0], nil
}

func (h *Handler) GetGoodCoins(ctx context.Context) ([]*npool.GoodCoin, uint32, error) {
	conds := &goodcoinmwpb.Conds{}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}

	coinHandler, err := goodcoinmw.NewHandler(
		ctx,
		goodcoinmw.WithConds(conds),
		goodcoinmw.WithOffset(h.Offset),
		goodcoinmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	infos, total, err := coinHandler.GetGoodCoins(ctx)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, total, nil
	}
	handler := &queryHandler{
		Handler:   h,
		goodCoins: infos,
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos, total, nil
}
