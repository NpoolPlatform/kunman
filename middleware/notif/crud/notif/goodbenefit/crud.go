package goodbenefit

import (
	"fmt"

	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entgoodbenefit "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/goodbenefit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID       *uuid.UUID
	GoodID      *uuid.UUID
	GoodType    *goodtypes.GoodType
	GoodName    *string
	CoinTypeID  *uuid.UUID
	Amount      *string
	State       *basetypes.Result
	Message     *string
	BenefitDate *uint32
	TxID        *uuid.UUID
	Generated   *bool
	DeletedAt   *uint32
}

func CreateSet(c *ent.GoodBenefitCreate, req *Req) *ent.GoodBenefitCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.GoodType != nil {
		c.SetGoodType(req.GoodType.String())
	}
	if req.GoodName != nil {
		c.SetGoodName(*req.GoodName)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.State != nil {
		c.SetState(req.State.String())
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.BenefitDate != nil {
		c.SetBenefitDate(*req.BenefitDate)
	}
	if req.TxID != nil {
		c.SetTxID(*req.TxID)
	}
	c.SetGenerated(false)
	return c
}

func UpdateSet(u *ent.GoodBenefitUpdateOne, req *Req) *ent.GoodBenefitUpdateOne {
	if req.Generated != nil {
		u.SetGenerated(*req.Generated)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID               *cruder.Cond
	EntID            *cruder.Cond
	GoodID           *cruder.Cond
	GoodType         *cruder.Cond
	GoodTypes        *cruder.Cond
	CoinTypeID       *cruder.Cond
	Generated        *cruder.Cond
	BenefitDateStart *cruder.Cond
	BenefitDateEnd   *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.GoodBenefitQuery, conds *Conds) (*ent.GoodBenefitQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.ID(id))
		default:
			return nil, fmt.Errorf("invalid good benefit id op field %s", conds.ID.Op)
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.EntID(id))
		default:
			return nil, fmt.Errorf("invalid good benefit op field %s", conds.EntID.Op)
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit good id")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid good benefit good id op field %s", conds.GoodID.Op)
		}
	}
	if conds.GoodType != nil {
		_type, ok := conds.GoodType.Val.(goodtypes.GoodType)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit good type")
		}
		switch conds.GoodType.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.GoodType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid good benefit good type op field %s", conds.GoodType.Op)
		}
	}
	if conds.GoodTypes != nil {
		_types, ok := conds.GoodTypes.Val.([]goodtypes.GoodType)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit good types")
		}
		switch conds.GoodTypes.Op {
		case cruder.IN:
			q.Where(entgoodbenefit.GoodTypeIn(func() (__types []string) {
				for _, _type := range _types {
					__types = append(__types, _type.String())
				}
				return
			}()...))
		default:
			return nil, fmt.Errorf("invalid good benefit good type op field %s", conds.GoodType.Op)
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit coin type id")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid good benefit good id op field %s", conds.CoinTypeID.Op)
		}
	}
	if conds.Generated != nil {
		notified, ok := conds.Generated.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit notified")
		}
		switch conds.Generated.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.Generated(notified))
		default:
			return nil, fmt.Errorf("invalid good benefit notified op field %s", conds.Generated.Op)
		}
	}
	if conds.BenefitDateStart != nil {
		_date, ok := conds.BenefitDateStart.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit benefit date %s", conds.BenefitDateStart.Op)
		}
		switch conds.BenefitDateStart.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.BenefitDateGTE(_date))
		case cruder.LTE:
			q.Where(entgoodbenefit.BenefitDateLTE(_date))
		case cruder.LT:
			q.Where(entgoodbenefit.BenefitDateLT(_date))
		case cruder.GT:
			q.Where(entgoodbenefit.BenefitDateGT(_date))
		case cruder.GTE:
			q.Where(entgoodbenefit.BenefitDateGTE(_date))
		default:
			return nil, fmt.Errorf("invalid good benefit benefit date op field %s", conds.BenefitDateStart.Op)
		}
	}
	if conds.BenefitDateEnd != nil {
		_date, ok := conds.BenefitDateEnd.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit benefit date %s", conds.BenefitDateEnd.Op)
		}
		switch conds.BenefitDateEnd.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.BenefitDateLTE(_date))
		case cruder.GTE:
			q.Where(entgoodbenefit.BenefitDateGTE(_date))
		case cruder.GT:
			q.Where(entgoodbenefit.BenefitDateGT(_date))
		case cruder.LT:
			q.Where(entgoodbenefit.BenefitDateLT(_date))
		case cruder.LTE:
			q.Where(entgoodbenefit.BenefitDateLTE(_date))
		default:
			return nil, fmt.Errorf("invalid good benefit benefit date op field %s", conds.BenefitDateEnd.Op)
		}
	}
	return q, nil
}
