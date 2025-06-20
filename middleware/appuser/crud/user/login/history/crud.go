package history

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entloginhistory "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/loginhistory"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	ClientIP  *string
	UserAgent *string
	Location  *string
	LoginType *basetypes.LoginType
}

func CreateSet(c *ent.LoginHistoryCreate, req *Req) *ent.LoginHistoryCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.ClientIP != nil {
		c.SetClientIP(*req.ClientIP)
	}
	if req.UserAgent != nil {
		c.SetUserAgent(*req.UserAgent)
	}
	if req.Location != nil {
		c.SetLocation(*req.Location)
	}
	if req.LoginType != nil {
		c.SetLoginType(req.LoginType.String())
	}
	return c
}

func UpdateSet(u *ent.LoginHistoryUpdateOne, req *Req) *ent.LoginHistoryUpdateOne {
	if req.Location != nil {
		u.SetLocation(*req.Location)
	}
	return u
}

type Conds struct {
	EntID     *cruder.Cond
	AppID     *cruder.Cond
	UserID    *cruder.Cond
	LoginType *cruder.Cond
	ClientIP  *cruder.Cond
	Location  *cruder.Cond
	UserAgent *cruder.Cond
}

func SetQueryConds(q *ent.LoginHistoryQuery, conds *Conds) (*ent.LoginHistoryQuery, error) { //nolint
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entloginhistory.EntID(id))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entloginhistory.AppID(id))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entloginhistory.UserID(id))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.LoginType != nil {
		loginType, ok := conds.LoginType.Val.(basetypes.LoginType)
		if !ok {
			return nil, fmt.Errorf("invalid login type")
		}
		switch conds.LoginType.Op {
		case cruder.EQ:
			q.Where(entloginhistory.LoginType(loginType.String()))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.ClientIP != nil {
		ip, ok := conds.ClientIP.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid client ip")
		}
		switch conds.ClientIP.Op {
		case cruder.EQ:
			q.Where(entloginhistory.ClientIP(ip))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.Location != nil {
		loc, ok := conds.Location.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid location")
		}
		switch conds.Location.Op {
		case cruder.EQ:
			q.Where(entloginhistory.Location(loc))
		case cruder.NEQ:
			q.Where(entloginhistory.LocationNEQ(loc))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.UserAgent != nil {
		agent, ok := conds.UserAgent.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid user agent")
		}
		switch conds.UserAgent.Op {
		case cruder.EQ:
			q.Where(entloginhistory.UserAgent(agent))
		default:
			return nil, fmt.Errorf("invalid user agent op field")
		}
	}
	return q, nil
}
