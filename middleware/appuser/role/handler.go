package role

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/role"
	rolecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/role"
	app "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	CreatedBy   *uuid.UUID
	Role        *string
	Description *string
	Default     *bool
	Genesis     *bool
	Reqs        []*rolecrud.Req
	Conds       *rolecrud.Conds
	Offset      int32
	Limit       int32
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

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		handler, err := app.NewHandler(
			ctx,
			app.WithEntID(id, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistApp(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid app")
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithCreatedBy(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid createdby")
			}
			return nil
		}
		// TODO: check user exist
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CreatedBy = &_id
		return nil
	}
}

func WithRole(role *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if role == nil {
			if must {
				return fmt.Errorf("invalid role")
			}
			return nil
		}
		if *role == "" {
			return fmt.Errorf("invalid role")
		}
		h.Role = role
		return nil
	}
}

func WithDescription(description *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Description = description
		return nil
	}
}

func WithDefault(defautl *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Default = defautl
		return nil
	}
}

func WithGenesis(genesis *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Genesis = genesis
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &rolecrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.CreatedBy != nil {
			id, err := uuid.Parse(conds.GetCreatedBy().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CreatedBy = &cruder.Cond{Op: conds.GetCreatedBy().GetOp(), Val: id}
		}
		if conds.Role != nil {
			h.Conds.Role = &cruder.Cond{
				Op:  conds.GetRole().GetOp(),
				Val: conds.GetRole().GetValue(),
			}
		}
		if conds.Default != nil {
			h.Conds.Default = &cruder.Cond{
				Op:  conds.GetDefault().GetOp(),
				Val: conds.GetDefault().GetValue(),
			}
		}
		if conds.Roles != nil {
			h.Conds.Roles = &cruder.Cond{
				Op:  conds.GetRoles().GetOp(),
				Val: conds.GetRoles().GetValue(),
			}
		}
		if conds.Genesis != nil {
			h.Conds.Genesis = &cruder.Cond{
				Op:  conds.GetGenesis().GetOp(),
				Val: conds.GetGenesis().GetValue(),
			}
		}
		if len(conds.GetEntIDs().GetValue()) > 0 {
			_ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				_ids = append(_ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{Op: conds.GetEntIDs().GetOp(), Val: _ids}
		}
		if len(conds.GetAppIDs().GetValue()) > 0 {
			_ids := []uuid.UUID{}
			for _, id := range conds.GetAppIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				_ids = append(_ids, _id)
			}
			h.Conds.AppIDs = &cruder.Cond{Op: conds.GetAppIDs().GetOp(), Val: _ids}
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

func WithReqs(reqs []*npool.RoleReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			_req := &rolecrud.Req{
				Role:        req.Role,
				Description: req.Description,
				Default:     req.Default,
				Genesis:     req.Genesis,
			}
			if req.AppID == nil {
				return fmt.Errorf("invalid appid")
			}
			appID, err := uuid.Parse(*req.AppID)
			if err != nil {
				return err
			}
			_req.AppID = &appID
			if req.CreatedBy == nil {
				return fmt.Errorf("invalid createdby")
			}
			createdBy, err := uuid.Parse(*req.CreatedBy)
			if err != nil {
				return err
			}
			_req.CreatedBy = &createdBy
			if req.EntID != nil {
				id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			h.Reqs = append(h.Reqs, _req)
		}
		return nil
	}
}
