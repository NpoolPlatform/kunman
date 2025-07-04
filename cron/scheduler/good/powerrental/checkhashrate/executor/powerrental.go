package executor

import (
	"context"
	"math"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	"github.com/shopspring/decimal"

	miningstockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/stock"
	goodpowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	goodusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/checkhashrate/types"
)

const (
	// 90%~110%
	MaxToleranceScope = 0.1
)

var (
	currentMiningGoodStockState = v1.MiningGoodStockState_MiningGoodStockStateCheckHashRate
	nextMiningGoodStockState    = v1.MiningGoodStockState_MiningGoodStockStateReady
)

type powerRentalHandler struct {
	*goodpowerrentalmwpb.PowerRental

	// for persistent
	miningGoodStockReqs []*miningstockmwpb.MiningGoodStockReq

	// for check
	goodCoinTypeIDs []string

	persistent chan interface{}
	notif      chan interface{}
	done       chan interface{}
}

func (h *powerRentalHandler) checkMiningGoodStockState() error {
	for _, miningGoodStock := range h.MiningGoodStocks {
		if miningGoodStock.State != currentMiningGoodStockState {
			return wlog.Errorf("invalid mininggoodstockstate: %v, mininggoodstock: %v", miningGoodStock.State, miningGoodStock.EntID)
		}
	}
	return nil
}

func (h *powerRentalHandler) getCoinTypeIDs() {
	h.goodCoinTypeIDs = []string{}
	for _, goodCoin := range h.GoodCoins {
		if goodCoin.Main {
			h.goodCoinTypeIDs = append(h.goodCoinTypeIDs, goodCoin.CoinTypeID)
		}
	}
	h.goodCoinTypeIDs = _removeRepeatedElement(h.goodCoinTypeIDs)
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

func (h *powerRentalHandler) checkHashRate(ctx context.Context) error {
	for _, miningGoodStock := range h.MiningGoodStocks {
		handler, err := goodusermw.NewHandler(
			ctx,
			goodusermw.WithEntID(&miningGoodStock.PoolGoodUserID, true),
			goodusermw.WithCoinTypeIDs(h.goodCoinTypeIDs, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		hRate, err := handler.GetGoodUserHashRate(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}

		_hashRate, err := decimal.NewFromString(hRate)
		if err != nil {
			return wlog.WrapError(err)
		}

		_total, err := decimal.NewFromString(miningGoodStock.Total)
		if err != nil {
			return wlog.WrapError(err)
		}

		unit, err := decimal.NewFromString(h.QuantityUnitAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		_total = _total.Mul(unit)

		hashRate := _hashRate.InexactFloat64()
		total := _total.InexactFloat64()
		if math.Abs(hashRate-total) > total*MaxToleranceScope {
			return wlog.Errorf("hash rate not up to total of mininggoodstock")
		}
	}
	return nil
}

func (h *powerRentalHandler) constructMiningGoodStockReqs() {
	_miningGoodStockReqs := []*miningstockmwpb.MiningGoodStockReq{}
	for _, req := range h.MiningGoodStocks {
		_miningGoodStockReqs = append(_miningGoodStockReqs,
			&miningstockmwpb.MiningGoodStockReq{
				EntID: &req.EntID,
				State: &nextMiningGoodStockState,
			},
		)
	}
	h.miningGoodStockReqs = _miningGoodStockReqs
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

	h.getCoinTypeIDs()

	if err = h.checkHashRate(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.checkMiningGoodStockState(); err != nil {
		return wlog.WrapError(err)
	}

	h.constructMiningGoodStockReqs()
	return nil
}
