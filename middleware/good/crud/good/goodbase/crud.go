package good

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID                *uuid.UUID
	GoodType             *types.GoodType
	BenefitType          *types.BenefitType
	Name                 *string
	ServiceStartAt       *uint32
	StartMode            *types.GoodStartMode
	TestOnly             *bool
	BenefitIntervalHours *uint32
	Purchasable          *bool
	Online               *bool
	State                *types.GoodState
	DeletedAt            *uint32
}

func CreateSet(c *ent.GoodBaseCreate, req *Req) *ent.GoodBaseCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodType != nil {
		c.SetGoodType(req.GoodType.String())
	}
	if req.BenefitType != nil {
		c.SetBenefitType(req.BenefitType.String())
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.ServiceStartAt != nil {
		c.SetServiceStartAt(*req.ServiceStartAt)
	}
	if req.StartMode != nil {
		c.SetStartMode(req.StartMode.String())
	}
	if req.TestOnly != nil {
		c.SetTestOnly(*req.TestOnly)
	}
	if req.BenefitIntervalHours != nil {
		c.SetBenefitIntervalHours(*req.BenefitIntervalHours)
	}
	if req.Purchasable != nil {
		c.SetPurchasable(*req.Purchasable)
	}
	if req.Online != nil {
		c.SetOnline(*req.Online)
	}
	if req.State != nil {
		c.SetState(req.State.String())
	}
	return c
}

func UpdateSet(u *ent.GoodBaseUpdateOne, req *Req) *ent.GoodBaseUpdateOne {
	if req.GoodType != nil {
		u.SetGoodType(req.GoodType.String())
	}
	if req.BenefitType != nil {
		u.SetBenefitType(req.BenefitType.String())
	}
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.ServiceStartAt != nil {
		u.SetServiceStartAt(*req.ServiceStartAt)
	}
	if req.StartMode != nil {
		u.SetStartMode(req.StartMode.String())
	}
	if req.TestOnly != nil {
		u.SetTestOnly(*req.TestOnly)
	}
	if req.BenefitIntervalHours != nil {
		u.SetBenefitIntervalHours(*req.BenefitIntervalHours)
	}
	if req.Purchasable != nil {
		u.SetPurchasable(*req.Purchasable)
	}
	if req.Online != nil {
		u.SetOnline(*req.Online)
	}
	if req.State != nil {
		u.SetState(req.State.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	IDs         *cruder.Cond
	EntID       *cruder.Cond
	EntIDs      *cruder.Cond
	GoodType    *cruder.Cond
	GoodTypes   *cruder.Cond
	BenefitType *cruder.Cond
	TestOnly    *cruder.Cond
	Purchasable *cruder.Cond
	Online      *cruder.Cond
	State       *cruder.Cond
}

func SetQueryConds(q *ent.GoodBaseQuery, conds *Conds) (*ent.GoodBaseQuery, error) {
	q.Where(entgoodbase.DeletedAt(0))
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
			q.Where(entgoodbase.ID(id))
		case cruder.NEQ:
			q.Where(entgoodbase.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid goodbase field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entgoodbase.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid goodbase field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodbase.EntID(id))
		case cruder.NEQ:
			q.Where(entgoodbase.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid goodbase field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entgoodbase.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid goodbase field")
		}
	}
	if conds.GoodType != nil {
		_type, ok := conds.GoodType.Val.(types.GoodType)
		if !ok {
			return nil, wlog.Errorf("invalid goodtype")
		}
		switch conds.GoodType.Op {
		case cruder.EQ:
			q.Where(entgoodbase.GoodType(_type.String()))
		default:
			return nil, wlog.Errorf("invalid goodbase field")
		}
	}
	if conds.GoodTypes != nil {
		_types, ok := conds.GoodTypes.Val.([]types.GoodType)
		if !ok {
			return nil, wlog.Errorf("invalid goodtypes")
		}
		es := []string{}
		for _, _type := range _types {
			es = append(es, _type.String())
		}
		switch conds.GoodTypes.Op {
		case cruder.IN:
			q.Where(entgoodbase.GoodTypeIn(es...))
		default:
			return nil, wlog.Errorf("invalid goodbase field")
		}
	}
	if conds.BenefitType != nil {
		_type, ok := conds.BenefitType.Val.(types.BenefitType)
		if !ok {
			return nil, wlog.Errorf("invalid benefittype")
		}
		switch conds.BenefitType.Op {
		case cruder.EQ:
			q.Where(entgoodbase.BenefitType(_type.String()))
		case cruder.NEQ:
			q.Where(entgoodbase.BenefitTypeNEQ(_type.String()))
		default:
			return nil, wlog.Errorf("invalid goodbase field")
		}
	}
	if conds.TestOnly != nil {
		b, ok := conds.TestOnly.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid testonly")
		}
		switch conds.TestOnly.Op {
		case cruder.EQ:
			q.Where(entgoodbase.TestOnly(b))
		case cruder.NEQ:
			q.Where(entgoodbase.TestOnlyNEQ(b))
		}
	}
	if conds.Purchasable != nil {
		b, ok := conds.Purchasable.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid purchasable")
		}
		switch conds.Purchasable.Op {
		case cruder.EQ:
			q.Where(entgoodbase.Purchasable(b))
		case cruder.NEQ:
			q.Where(entgoodbase.PurchasableNEQ(b))
		}
	}
	if conds.Online != nil {
		b, ok := conds.Online.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid online")
		}
		switch conds.Online.Op {
		case cruder.EQ:
			q.Where(entgoodbase.Online(b))
		case cruder.NEQ:
			q.Where(entgoodbase.OnlineNEQ(b))
		}
	}
	if conds.State != nil {
		_state, ok := conds.State.Val.(types.GoodState)
		if !ok {
			return nil, wlog.Errorf("invalid goodstate")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(entgoodbase.State(_state.String()))
		case cruder.NEQ:
			q.Where(entgoodbase.StateNEQ(_state.String()))
		default:
			return nil, wlog.Errorf("invalid goodbase field")
		}
	}
	return q, nil
}
