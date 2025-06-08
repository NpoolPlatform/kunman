package auth

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/authing/auth"
	handler "github.com/NpoolPlatform/kunman/middleware/appuser/authing/handler"
	authcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/authing/auth"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	*handler.Handler
	Conds *authcrud.Conds
	Reqs  []*authcrud.Req
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	_handler, err := handler.NewHandler(ctx, options...)
	if err != nil {
		return nil, err
	}

	h := &Handler{
		Handler: _handler,
	}
	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, h); err != nil {
			return nil, err
		}
	}
	return h, nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &authcrud.Conds{}
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
		if conds.RoleID != nil {
			id, err := uuid.Parse(conds.GetRoleID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.RoleID = &cruder.Cond{Op: conds.GetRoleID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.Resource != nil {
			h.Conds.Resource = &cruder.Cond{
				Op:  conds.GetResource().GetOp(),
				Val: conds.GetResource().GetValue(),
			}
		}
		if conds.Method != nil {
			h.Conds.Method = &cruder.Cond{
				Op:  conds.GetMethod().GetOp(),
				Val: conds.GetMethod().GetValue(),
			}
		}
		return nil
	}
}

func WithReqs(reqs []*npool.AuthReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			_req := &authcrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.AppID == nil {
				return fmt.Errorf("invalid appid")
			}
			appID, err := uuid.Parse(*req.AppID)
			if err != nil {
				return err
			}
			_req.AppID = &appID
			if req.UserID != nil {
				userID, err := uuid.Parse(*req.UserID)
				if err != nil {
					return err
				}
				_req.UserID = &userID
			}
			if req.RoleID != nil {
				roleID, err := uuid.Parse(*req.RoleID)
				if err != nil {
					return err
				}
				_req.RoleID = &roleID
			}
			if req.Resource == nil {
				return fmt.Errorf("invalid resource")
			}
			const leastResourceLen = 3
			if len(*req.Resource) < leastResourceLen {
				return fmt.Errorf("resource %v invalid", *req.Resource)
			}
			_req.Resource = req.Resource
			if req.Method == nil {
				return fmt.Errorf("invalid method")
			}
			switch *req.Method {
			case "POST":
			case "GET":
			default:
				return fmt.Errorf("method %v invalid", *req.Method)
			}
			_req.Method = req.Method
			h.Reqs = append(h.Reqs, _req)
		}
		return nil
	}
}
