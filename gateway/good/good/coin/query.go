package goodcoin

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
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
	info, err := goodcoinmwcli.GetGoodCoin(ctx, *h.EntID)
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

	infos, total, err := goodcoinmwcli.GetGoodCoins(ctx, conds, h.Offset, h.Limit)
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
