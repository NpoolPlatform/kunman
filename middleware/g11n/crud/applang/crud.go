//nolint:nolintlint,dupl
package applang

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
	entapplang "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/applang"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	Main      *bool
	DeletedAt *uint32
}

func CreateSet(c *ent.AppLangCreate, req *Req) *ent.AppLangCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.LangID != nil {
		c.SetLangID(*req.LangID)
	}
	if req.Main != nil {
		c.SetMain(*req.Main)
	}
	return c
}

func UpdateSet(u *ent.AppLangUpdateOne, req *Req) *ent.AppLangUpdateOne {
	if req.Main != nil {
		u.SetMain(*req.Main)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	EntID   *cruder.Cond
	EntIDs  *cruder.Cond
	AppID   *cruder.Cond
	LangID  *cruder.Cond
	AppIDs  *cruder.Cond
	LangIDs *cruder.Cond
	Main    *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppLangQuery, conds *Conds) (*ent.AppLangQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(
				entapplang.ID(id),
				entapplang.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entapplang.IDNEQ(id),
				entapplang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(
				entapplang.EntID(id),
				entapplang.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entapplang.EntIDNEQ(id),
				entapplang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(
				entapplang.EntIDIn(ids...),
				entapplang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(
				entapplang.AppID(id),
				entapplang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(
				entapplang.AppIDIn(ids...),
				entapplang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid appids field")
		}
	}
	if conds.LangID != nil {
		id, ok := conds.LangID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langid")
		}
		switch conds.LangID.Op {
		case cruder.EQ:
			q.Where(
				entapplang.LangID(id),
				entapplang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid langid field")
		}
	}
	if conds.LangIDs != nil {
		ids, ok := conds.LangIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langids")
		}
		switch conds.LangIDs.Op {
		case cruder.IN:
			q.Where(
				entapplang.LangIDIn(ids...),
				entapplang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid langids field")
		}
	}
	if conds.Main != nil {
		main, ok := conds.Main.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid main")
		}
		switch conds.Main.Op {
		case cruder.EQ:
			q.Where(entapplang.Main(main))
		default:
			return nil, fmt.Errorf("invalid main field")
		}
	}
	return q, nil
}
