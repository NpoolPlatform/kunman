package goodcoin

import (
	"context"

	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constant "github.com/NpoolPlatform/good-gateway/pkg/const"
	goodcoinmwcli "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	goodCoins []*npool.GoodCoin
}

func (h *createHandler) getGoodCoins(ctx context.Context) error {
	h.Limit = constant.DefaultRowLimit
	for {
		goodCoins, _, err := h.GetGoodCoins(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(goodCoins) == 0 {
			return nil
		}
		h.goodCoins = append(h.goodCoins, goodCoins...)
		h.Offset += h.Limit
	}
}

func (h *createHandler) validateCandidateCoin(ctx context.Context) error {
	if len(h.goodCoins) == 0 {
		return nil
	}
	exist, err := coinmwcli.ExistCoinConds(ctx, &coinmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.CoinTypeID},
		ENV:   &basetypes.StringVal{Op: cruder.EQ, Value: h.goodCoins[0].CoinENV},
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if !exist {
		return wlog.Errorf("invalid coin")
	}
	return nil
}

func (h *Handler) CreateGoodCoin(ctx context.Context) (*npool.GoodCoin, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.getGoodCoins(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateCandidateCoin(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := goodcoinmwcli.CreateGoodCoin(ctx, &goodcoinmwpb.GoodCoinReq{
		EntID:      h.EntID,
		GoodID:     h.GoodID,
		CoinTypeID: h.CoinTypeID,
		Main:       h.Main,
		Index:      h.Index,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetGoodCoin(ctx)
}
