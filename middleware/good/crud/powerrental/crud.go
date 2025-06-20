package good

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entpowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/powerrental"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID               *uuid.UUID
	GoodID              *uuid.UUID
	DeviceTypeID        *uuid.UUID
	VendorLocationID    *uuid.UUID
	UnitPrice           *decimal.Decimal
	QuantityUnit        *string
	QuantityUnitAmount  *decimal.Decimal
	DeliveryAt          *uint32
	UnitLockDeposit     *decimal.Decimal
	DurationDisplayType *types.GoodDurationType
	StockMode           *types.GoodStockMode
	DeletedAt           *uint32
}

func CreateSet(c *ent.PowerRentalCreate, req *Req) *ent.PowerRentalCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.DeviceTypeID != nil {
		c.SetDeviceTypeID(*req.DeviceTypeID)
	}
	if req.VendorLocationID != nil {
		c.SetVendorLocationID(*req.VendorLocationID)
	}
	if req.UnitPrice != nil {
		c.SetUnitPrice(*req.UnitPrice)
	}
	if req.QuantityUnit != nil {
		c.SetQuantityUnit(*req.QuantityUnit)
	}
	if req.QuantityUnitAmount != nil {
		c.SetQuantityUnitAmount(*req.QuantityUnitAmount)
	}
	if req.DeliveryAt != nil {
		c.SetDeliveryAt(*req.DeliveryAt)
	}
	if req.UnitLockDeposit != nil {
		c.SetUnitLockDeposit(*req.UnitLockDeposit)
	}
	if req.DurationDisplayType != nil {
		c.SetDurationDisplayType(req.DurationDisplayType.String())
	}
	if req.StockMode != nil {
		c.SetStockMode(req.StockMode.String())
	}
	return c
}

func UpdateSet(u *ent.PowerRentalUpdateOne, req *Req) *ent.PowerRentalUpdateOne {
	if req.UnitPrice != nil {
		u.SetUnitPrice(*req.UnitPrice)
	}
	if req.QuantityUnit != nil {
		u.SetQuantityUnit(*req.QuantityUnit)
	}
	if req.QuantityUnitAmount != nil {
		u.SetQuantityUnitAmount(*req.QuantityUnitAmount)
	}
	if req.DeliveryAt != nil {
		u.SetDeliveryAt(*req.DeliveryAt)
	}
	if req.UnitLockDeposit != nil {
		u.SetUnitLockDeposit(*req.UnitLockDeposit)
	}
	if req.DurationDisplayType != nil {
		u.SetDurationDisplayType(req.DurationDisplayType.String())
	}
	if req.StockMode != nil {
		u.SetStockMode(req.StockMode.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID                *cruder.Cond
	IDs               *cruder.Cond
	EntID             *cruder.Cond
	EntIDs            *cruder.Cond
	GoodID            *cruder.Cond
	GoodIDs           *cruder.Cond
	DeviceTypeID      *cruder.Cond
	DeviceTypeIDs     *cruder.Cond
	VendorLocationID  *cruder.Cond
	VendorLocationIDs *cruder.Cond
	StockMode         *cruder.Cond
}

func SetQueryConds(q *ent.PowerRentalQuery, conds *Conds) (*ent.PowerRentalQuery, error) {
	q.Where(entpowerrental.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entpowerrental.ID(id))
		case cruder.NEQ:
			q.Where(entpowerrental.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entpowerrental.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entpowerrental.EntID(id))
		case cruder.NEQ:
			q.Where(entpowerrental.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entpowerrental.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entpowerrental.GoodID(id))
		case cruder.NEQ:
			q.Where(entpowerrental.GoodIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entpowerrental.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.DeviceTypeID != nil {
		id, ok := conds.DeviceTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid devicetypeid")
		}
		switch conds.DeviceTypeID.Op {
		case cruder.EQ:
			q.Where(entpowerrental.DeviceTypeID(id))
		case cruder.NEQ:
			q.Where(entpowerrental.DeviceTypeIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.DeviceTypeIDs != nil {
		ids, ok := conds.DeviceTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid devicetypeids")
		}
		switch conds.DeviceTypeIDs.Op {
		case cruder.IN:
			q.Where(entpowerrental.DeviceTypeIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.VendorLocationID != nil {
		id, ok := conds.VendorLocationID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid vendorlocationid")
		}
		switch conds.VendorLocationID.Op {
		case cruder.EQ:
			q.Where(entpowerrental.VendorLocationID(id))
		case cruder.NEQ:
			q.Where(entpowerrental.VendorLocationIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.VendorLocationIDs != nil {
		ids, ok := conds.VendorLocationIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid vendorlocationids")
		}
		switch conds.VendorLocationIDs.Op {
		case cruder.IN:
			q.Where(entpowerrental.VendorLocationIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.StockMode != nil {
		_mode, ok := conds.StockMode.Val.(types.GoodStockMode)
		if !ok {
			return nil, wlog.Errorf("invalid stockmode")
		}
		switch conds.StockMode.Op {
		case cruder.EQ:
			q.Where(entpowerrental.StockMode(_mode.String()))
		case cruder.NEQ:
			q.Where(entpowerrental.StockModeNEQ(_mode.String()))
		default:
			return nil, wlog.Errorf("invalid stockmode field")
		}
	}
	return q, nil
}
