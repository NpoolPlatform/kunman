package powerrental

import (
	"time"

	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/shopspring/decimal"
)

type PowerRental interface {
	MinOrderAmount() decimal.Decimal
	MaxOrderAmount() decimal.Decimal
	MinOrderDurationSeconds() uint32
	MaxOrderDurationSeconds() uint32
	GoodServiceStartAt() uint32
	GoodStartMode() types.GoodStartMode
}

type powerRental struct {
	powerRental         *ent.PowerRental
	goodBase            *ent.GoodBase
	appGoodBase         *ent.AppGoodBase
	appPowerRental      *ent.AppPowerRental
	stock               *ent.Stock
	miningGoodStocks    []*ent.MiningGoodStock
	appGoodStock        *ent.AppStock
	appMiningGoodStocks []*ent.AppMiningGoodStock
}

func (pr *powerRental) MinOrderAmount() decimal.Decimal {
	return pr.appPowerRental.MinOrderAmount
}

func (pr *powerRental) MaxOrderAmount() decimal.Decimal {
	return pr.appPowerRental.MaxOrderAmount
}

func (pr *powerRental) MinOrderDurationSeconds() uint32 {
	return pr.appPowerRental.MinOrderDurationSeconds
}

func (pr *powerRental) MaxOrderDurationSeconds() uint32 {
	return pr.appPowerRental.MaxOrderDurationSeconds
}

func (pr *powerRental) GoodServiceStartAt() uint32 {
	now := uint32(time.Now().Unix())
	if now < pr.goodBase.ServiceStartAt {
		return pr.goodBase.ServiceStartAt
	}
	return now
}

func (pr *powerRental) GoodStartMode() types.GoodStartMode {
	return types.GoodStartMode(types.GoodStartMode_value[pr.goodBase.StartMode])
}
