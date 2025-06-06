package api

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/basal-middleware/pkg/const"
	crud "github.com/NpoolPlatform/basal-middleware/pkg/crud/api"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	"github.com/google/uuid"
)

type Handler struct {
	crud.Req
	Reqs   []*crud.Req
	Conds  *crud.Conds
	Offset int32
	Limit  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithProtocol(protocol *npool.Protocol, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if protocol == nil {
			if must {
				return fmt.Errorf("invalid protocol")
			}
			return nil
		}
		switch *protocol {
		case npool.Protocol_HTTP:
		case npool.Protocol_GRPC:
		default:
			return fmt.Errorf("invalid protocol %v: ", *protocol)
		}

		h.Protocol = protocol
		return nil
	}
}

func WithServiceName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("service name is empty")
			}
			return nil
		}
		const leastNameLen = 2
		if len(*name) < leastNameLen {
			return fmt.Errorf("service name %v too short", *name)
		}

		h.ServiceName = name
		return nil
	}
}

func WithMethod(method *npool.Method, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if method == nil {
			if must {
				return fmt.Errorf("invalid method")
			}
			return nil
		}
		switch *method {
		case npool.Method_GET:
		case npool.Method_POST:
		case npool.Method_STREAM:
		default:
			return fmt.Errorf("invalid method %v: ", *method)
		}

		h.Method = method
		return nil
	}
}

func WithMethodName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid method name")
			}
			return nil
		}
		h.MethodName = name
		return nil
	}
}

func WithPath(path *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if path == nil {
			if must {
				return fmt.Errorf("invalid path")
			}
			return nil
		}
		h.Path = path
		return nil
	}
}

func WithPathPrefix(prefix *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if prefix == nil {
			if must {
				return fmt.Errorf("invalid pathprefix")
			}
			return nil
		}
		h.PathPrefix = prefix
		return nil
	}
}

func WithExported(exported *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if exported == nil {
			if must {
				return fmt.Errorf("invalid exported")
			}
			return nil
		}
		h.Exported = exported
		return nil
	}
}

func WithDeprecated(deprecated *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if deprecated == nil {
			if must {
				return fmt.Errorf("invalid depracated")
			}
			return nil
		}
		h.Deprecated = deprecated
		return nil
	}
}

func WithDomains(domains *[]string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if domains == nil {
			if must {
				return fmt.Errorf("invalid domains")
			}
			return nil
		}
		h.Domains = domains
		return nil
	}
}

func WithReqs(reqs []*npool.APIReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*crud.Req{}
		for _, req := range reqs {
			_req := &crud.Req{
				Method:     req.Method,
				MethodName: req.MethodName,
				Path:       req.Path,
				PathPrefix: req.PathPrefix,
				Exported:   req.Exported,
				Deprecated: req.Deprecated,
			}
			if req.EntID != nil {
				_id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &_id
			}
			if req.Protocol != nil {
				switch *req.Protocol {
				case npool.Protocol_HTTP:
				case npool.Protocol_GRPC:
				default:
					return fmt.Errorf("invalid protocol")
				}
				_req.Protocol = req.Protocol
			}
			if req.ServiceName != nil {
				const leastNameLen = 2
				if len(*req.ServiceName) < leastNameLen {
					return fmt.Errorf("service name too short")
				}
				_req.ServiceName = req.ServiceName
			}
			if len(req.GetDomains()) > 0 {
				_req.Domains = &req.Domains
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.Protocol != nil {
			h.Conds.Protocol = &cruder.Cond{
				Op:  conds.GetProtocol().GetOp(),
				Val: npool.Protocol(conds.GetProtocol().GetValue()),
			}
		}
		if conds.ServiceName != nil {
			h.Conds.ServiceName = &cruder.Cond{
				Op:  conds.GetServiceName().GetOp(),
				Val: conds.GetServiceName().GetValue(),
			}
		}
		if conds.Method != nil {
			h.Conds.Method = &cruder.Cond{
				Op:  conds.GetMethod().GetOp(),
				Val: npool.Method(conds.GetMethod().GetValue()),
			}
		}
		if conds.Path != nil {
			h.Conds.Path = &cruder.Cond{
				Op:  conds.GetPath().GetOp(),
				Val: conds.GetPath().GetValue(),
			}
		}
		if conds.Exported != nil {
			h.Conds.Exported = &cruder.Cond{
				Op:  conds.GetExported().GetOp(),
				Val: conds.GetExported().GetValue(),
			}
		}
		if conds.Deprecated != nil {
			h.Conds.Deprecated = &cruder.Cond{
				Op:  conds.GetDeprecated().GetOp(),
				Val: conds.GetDeprecated().GetValue(),
			}
		}
		if len(conds.GetEntIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
