package default1

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	defaultmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/default"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/default"
	defaultmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/default"
)

type queryHandler struct {
	*Handler
	defaults []*defaultmwpb.Default
	infos    []*npool.Default
	apps     map[string]*appmwpb.App
	coins    map[string]*coinmwpb.Coin
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, defalut := range h.defaults {
			appIDs = append(appIDs, defalut.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.coins, err = goodgwcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, def := range h.defaults {
			coinTypeIDs = append(coinTypeIDs, def.CoinTypeID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, def := range h.defaults {
		info := &npool.Default{
			ID:          def.ID,
			EntID:       def.EntID,
			AppID:       def.AppID,
			GoodID:      def.GoodID,
			GoodName:    def.GoodName,
			AppGoodID:   def.AppGoodID,
			AppGoodName: def.AppGoodName,
			CoinTypeID:  def.CoinTypeID,
			CreatedAt:   def.CreatedAt,
			UpdatedAt:   def.UpdatedAt,
		}

		app, ok := h.apps[def.AppID]
		if ok {
			info.AppName = app.Name
		}
		coin, ok := h.coins[def.CoinTypeID]
		if ok {
			info.CoinName = coin.Name
			info.CoinLogo = coin.Logo
			info.CoinEnv = coin.ENV
			info.CoinUnit = coin.Unit
		}

		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetDefault(ctx context.Context) (*npool.Default, error) {
	info, err := defaultmwcli.GetDefault(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid default")
	}

	handler := &queryHandler{
		Handler:  h,
		defaults: []*defaultmwpb.Default{info},
		apps:     map[string]*appmwpb.App{},
		coins:    map[string]*coinmwpb.Coin{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetDefaults(ctx context.Context) ([]*npool.Default, uint32, error) {
	infos, total, err := defaultmwcli.GetDefaults(ctx, &defaultmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:  h,
		defaults: infos,
		apps:     map[string]*appmwpb.App{},
		coins:    map[string]*coinmwpb.Coin{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
