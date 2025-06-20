package applang

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/applang"
	applangcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/applang"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID     *uint32
	EntID  *uuid.UUID
	AppID  *uuid.UUID
	LangID *uuid.UUID
	Main   *bool
	Reqs   []*applangcrud.Req
	Conds  *applangcrud.Conds
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithLangID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid langid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.LangID = &_id
		return nil
	}
}

func WithMain(main *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if main == nil {
			if must {
				return fmt.Errorf("invalid main")
			}
			return nil
		}
		h.Main = main
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &applangcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
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
		if conds.LangID != nil {
			id, err := uuid.Parse(conds.GetLangID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.LangID = &cruder.Cond{Op: conds.GetLangID().GetOp(), Val: id}
		}
		if conds.Main != nil {
			h.Conds.Main = &cruder.Cond{Op: conds.GetMain().GetOp(), Val: conds.GetMain().GetValue()}
		}
		if len(conds.GetLangIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetLangIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.LangIDs = &cruder.Cond{Op: conds.GetLangIDs().GetOp(), Val: ids}
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

func WithReqs(reqs []*npool.LangReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*applangcrud.Req{}
		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.LangID == nil {
					return fmt.Errorf("invalid langid")
				}
			}
			_req := &applangcrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.AppID != nil {
				id, err := uuid.Parse(*req.AppID)
				if err != nil {
					return err
				}
				_req.AppID = &id
			}
			if req.LangID != nil {
				id, err := uuid.Parse(*req.LangID)
				if err != nil {
					return err
				}
				_req.LangID = &id
			}
			if req.Main != nil {
				_req.Main = req.Main
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
