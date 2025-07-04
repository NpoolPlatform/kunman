package executor

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"

	miningstockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/stock"
	goodpowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	poolcoinmwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/coin"
	goodusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/gooduser"
	rootusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/rootuser"
	poolcoinmw "github.com/NpoolPlatform/kunman/middleware/miningpool/coin"
	rootusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/rootuser"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/creategooduser/types"
)

var (
	currentMiningGoodStockState = v1.MiningGoodStockState_MiningGoodStockStateCreateGoodUser
	nextMiningGoodStockState    = v1.MiningGoodStockState_MiningGoodStockStateCheckHashRate
)

type powerRentalHandler struct {
	*goodpowerrentalmwpb.PowerRental

	// for persistent
	miningGoodStockReqs []*miningstockmwpb.MiningGoodStockReq
	goodUserReqs        []*goodusermwpb.GoodUserReq

	// for check
	rootUsers           []*rootusermwpb.RootUser
	rootUserCoinTypeIDs map[string][]string
	goodCoinTypeIDs     []string
	persistent          chan interface{}
	notif               chan interface{}
	done                chan interface{}
}

func (h *powerRentalHandler) checkMiningGoodStockState() error {
	for _, miningGoodStock := range h.MiningGoodStocks {
		if miningGoodStock.State != currentMiningGoodStockState {
			return wlog.Errorf("invalid mininggoodstockstate: %v, mininggoodstock: %v", miningGoodStock.State, miningGoodStock.EntID)
		}
	}
	return nil
}

func (h *powerRentalHandler) getCoinTypeIDs() error {
	if len(h.GoodCoins) == 0 {
		return wlog.Errorf("powerrental good have no coins")
	}

	h.goodCoinTypeIDs = []string{}
	haveMainCoin := false
	for _, goodCoin := range h.GoodCoins {
		h.goodCoinTypeIDs = append(h.goodCoinTypeIDs, goodCoin.CoinTypeID)
		if goodCoin.Main {
			haveMainCoin = true
		}
	}
	if !haveMainCoin {
		return wlog.Errorf("powerrental good have no main coin")
	}

	h.goodCoinTypeIDs = _removeRepeatedElement(h.goodCoinTypeIDs)
	return nil
}

func _removeRepeatedElement(arr []string) []string {
	newArr := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

func (h *powerRentalHandler) getPoolInfos(ctx context.Context) error {
	var err error
	poolRootUserIDs := []string{}
	for _, miningGoodStock := range h.MiningGoodStocks {
		poolRootUserIDs = append(poolRootUserIDs, miningGoodStock.PoolRootUserID)
	}

	conds := &rootusermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: poolRootUserIDs},
		Authed: &basetypes.BoolVal{Op: cruder.EQ, Value: true},
	}
	handler, err := rootusermw.NewHandler(
		ctx,
		rootusermw.WithConds(conds),
		rootusermw.WithOffset(0),
		rootusermw.WithLimit(int32(len(poolRootUserIDs))),
	)
	if err != nil {
		return err
	}

	h.rootUsers, _, err = handler.GetRootUsers(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.rootUserCoinTypeIDs = make(map[string][]string)
	for _, rootUser := range h.rootUsers {
		conds := &poolcoinmwpb.Conds{
			CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: h.goodCoinTypeIDs},
			PoolID:      &basetypes.StringVal{Op: cruder.EQ, Value: rootUser.PoolID},
		}
		handler, err := poolcoinmw.NewHandler(
			ctx,
			poolcoinmw.WithConds(conds),
			poolcoinmw.WithOffset(0),
			poolcoinmw.WithLimit(int32(len(h.goodCoinTypeIDs))),
		)
		if err != nil {
			return err
		}

		coinInfos, _, err := handler.GetCoins(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}

		// for check if coinTypeIDs have been support by miningpool
		coinTypeIDs := []string{}
		for _, coinInfo := range coinInfos {
			coinTypeIDs = append(coinTypeIDs, coinInfo.CoinTypeID)
		}
		h.rootUserCoinTypeIDs[rootUser.EntID] = _removeRepeatedElement(coinTypeIDs)
	}

	return wlog.WrapError(err)
}

func (h *powerRentalHandler) checkPoolRootUsers() error {
	for _, miningGoodStock := range h.MiningGoodStocks {
		exist := false
		for _, rootUser := range h.rootUsers {
			if miningGoodStock.PoolRootUserID == rootUser.EntID {
				exist = true
				break
			}
		}

		if !exist {
			return wlog.Errorf("have no rootuser for mininggoodstock, rootuser: %v", miningGoodStock.PoolRootUserID)
		}

		ids, ok := h.rootUserCoinTypeIDs[miningGoodStock.PoolRootUserID]
		if !ok || len(ids) < len(h.goodCoinTypeIDs) {
			return wlog.Errorf("miningpool not support mininggoodstock, rootuser: %v", miningGoodStock.PoolRootUserID)
		}
	}
	return nil
}

func (h *powerRentalHandler) constructMiningGoodStockReqs() {
	_miningGoodStockReqs := []*miningstockmwpb.MiningGoodStockReq{}
	_gooduserReqs := []*goodusermwpb.GoodUserReq{}
	for _, req := range h.MiningGoodStocks {
		poolGoodUserID := uuid.NewString()
		_miningGoodStockReqs = append(_miningGoodStockReqs,
			&miningstockmwpb.MiningGoodStockReq{
				EntID:          &req.EntID,
				State:          &nextMiningGoodStockState,
				PoolGoodUserID: &poolGoodUserID,
			},
		)
		_gooduserReqs = append(_gooduserReqs, &goodusermwpb.GoodUserReq{
			EntID:       &poolGoodUserID,
			CoinTypeIDs: h.goodCoinTypeIDs,
			RootUserID:  &req.PoolRootUserID,
		})
	}
	h.miningGoodStockReqs = _miningGoodStockReqs
	h.goodUserReqs = _gooduserReqs
}

//nolint:gocritic
func (h *powerRentalHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Powerrental Good", h.PowerRental,
			"PaymentTransferCoins", h.MiningGoodStocks,
			"Error", *err,
		)
	}

	persistentPowerRental := &types.PersistentGoodPowerRental{
		PowerRental:         h.PowerRental,
		MiningGoodStockReqs: h.miningGoodStockReqs,
		GoodUserReqs:        h.goodUserReqs,
	}

	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentPowerRental, h.persistent)
	} else {
		asyncfeed.AsyncFeed(ctx, persistentPowerRental, h.done)
	}
}

func (h *powerRentalHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	if err = h.getCoinTypeIDs(); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.getPoolInfos(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.checkPoolRootUsers(); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.checkMiningGoodStockState(); err != nil {
		return wlog.WrapError(err)
	}

	h.constructMiningGoodStockReqs()
	return nil
}
