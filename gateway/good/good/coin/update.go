package goodcoin

import (
	"context"

	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateGoodCoin(ctx context.Context) (*npool.GoodCoin, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkGoodCoin(ctx); err != nil {
		return nil, err
	}
	if err := goodcoinmwcli.UpdateGoodCoin(ctx, &goodcoinmwpb.GoodCoinReq{
		ID:    h.ID,
		EntID: h.EntID,
		Main:  h.Main,
		Index: h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetGoodCoin(ctx)
}
