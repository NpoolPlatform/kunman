package quota

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entquota "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/quota"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID            *uint32
	EntID         *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	Quota         *uint32
	ConsumedQuota *uint32
	ExpiredAt     *uint32
	DeletedAt     *uint32
}

func CreateSet(c *ent.QuotaCreate, req *Req) *ent.QuotaCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Quota != nil {
		c.SetQuota(*req.Quota)
	}
	if req.ConsumedQuota != nil {
		c.SetConsumedQuota(*req.ConsumedQuota)
	}
	if req.ExpiredAt != nil {
		c.SetExpiredAt(*req.ExpiredAt)
	}
	return c
}

func UpdateSet(u *ent.QuotaUpdateOne, req *Req) *ent.QuotaUpdateOne {
	if req.ConsumedQuota != nil {
		u.SetConsumedQuota(*req.ConsumedQuota)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	IDs     *cruder.Cond
	EntID   *cruder.Cond
	EntIDs  *cruder.Cond
	AppID   *cruder.Cond
	AppIDs  *cruder.Cond
	UserID  *cruder.Cond
	UserIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.QuotaQuery, conds *Conds) (*ent.QuotaQuery, error) {
	q.Where(entquota.DeletedAt(0))
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
			q.Where(entquota.ID(id))
		default:
			return nil, wlog.Errorf("invalid quota field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entquota.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid quota field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entquota.EntID(id))
		default:
			return nil, wlog.Errorf("invalid quota field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entquota.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid quota field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entquota.AppID(id))
		default:
			return nil, wlog.Errorf("invalid quota field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entquota.AppIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid quota field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entquota.UserID(id))
		default:
			return nil, wlog.Errorf("invalid quota field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entquota.UserIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid quota field")
		}
	}

	return q, nil
}
