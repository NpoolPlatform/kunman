package api

import (
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/basal/middleware/v1/api"
	ent "github.com/NpoolPlatform/kunman/middleware/basal/db/ent/generated"
	entapi "github.com/NpoolPlatform/kunman/middleware/basal/db/ent/generated/api"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	Protocol    *npool.Protocol
	ServiceName *string
	Method      *npool.Method
	MethodName  *string
	Path        *string
	Exported    *bool
	PathPrefix  *string
	Domains     *[]string
	Deprecated  *bool
	DeletedAt   *uint32
}

func CreateSet(c *ent.APICreate, req *Req) *ent.APICreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Protocol != nil {
		c.SetProtocol(req.Protocol.String())
	}
	if req.ServiceName != nil {
		c.SetServiceName(*req.ServiceName)
	}
	if req.Method != nil {
		c.SetMethod(req.Method.String())
	}
	if req.MethodName != nil {
		c.SetMethodName(*req.MethodName)
	}
	if req.Path != nil {
		c.SetPath(*req.Path)
	}
	if req.PathPrefix != nil {
		c.SetPathPrefix(*req.PathPrefix)
	}
	if req.Domains != nil {
		c.SetDomains(*req.Domains)
	}
	if req.Exported != nil {
		c.SetExported(*req.Exported)
	}
	if req.Deprecated != nil {
		c.SetDeprecated(*req.Deprecated)
	}
	return c
}

func UpdateSet(u *ent.APIUpdateOne, req *Req) *ent.APIUpdateOne {
	if req.Exported != nil {
		u.SetExported(*req.Exported)
	}
	if req.Deprecated != nil {
		u.SetDeprecated(*req.Deprecated)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	Protocol    *cruder.Cond
	ServiceName *cruder.Cond
	Method      *cruder.Cond
	Path        *cruder.Cond
	Exported    *cruder.Cond
	Deprecated  *cruder.Cond
	EntIDs      *cruder.Cond
}

func SetQueryConds(q *ent.APIQuery, conds *Conds) (*ent.APIQuery, error) { //nolint
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entapi.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.Protocol != nil {
		protocol, ok := conds.Protocol.Val.(npool.Protocol)
		if !ok {
			return nil, fmt.Errorf("invalid protocol")
		}
		switch conds.Protocol.Op {
		case cruder.EQ:
			q.Where(entapi.Protocol(protocol.String()))
		default:
			return nil, fmt.Errorf("invalid protocol field")
		}
	}
	if conds.ServiceName != nil {
		name, ok := conds.ServiceName.Val.(string)
		if !ok {
			return nil, fmt.Errorf("fail transfer service name: %v to string", name)
		}
		switch conds.ServiceName.Op {
		case cruder.EQ:
			q.Where(entapi.ServiceName(name))
		default:
			return nil, fmt.Errorf("invalid service name field")
		}
	}
	if conds.Method != nil {
		method, ok := conds.Method.Val.(npool.Method)
		if !ok {
			return nil, fmt.Errorf("invalid method")
		}
		switch conds.Method.Op {
		case cruder.EQ:
			q.Where(entapi.Method(method.String()))
		default:
			return nil, fmt.Errorf("invalid method field")
		}
	}
	if conds.Path != nil {
		path, ok := conds.Path.Val.(string)
		if !ok {
			return nil, fmt.Errorf("fail transfer path: %v to string", path)
		}
		switch conds.Path.Op {
		case cruder.EQ:
			q.Where(entapi.Path(path))
		default:
			return nil, fmt.Errorf("invalid path field")
		}
	}
	if conds.Exported != nil {
		exported, ok := conds.Exported.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid exported")
		}
		switch conds.Exported.Op {
		case cruder.EQ:
			q.Where(entapi.Exported(exported))
		default:
			return nil, fmt.Errorf("invalid exported field")
		}
	}
	if conds.Deprecated != nil {
		deprecated, ok := conds.Deprecated.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid deprecated")
		}
		switch conds.Deprecated.Op {
		case cruder.EQ:
			q.Where(entapi.Deprecated(deprecated))
		default:
			return nil, fmt.Errorf("invalid deprecated field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entapi.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid api entids filed")
		}
	}
	return q, nil
}
