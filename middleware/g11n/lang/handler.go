package lang

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/lang"
	langcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/lang"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID     *uint32
	EntID  *uuid.UUID
	Lang   *string
	Name   *string
	Logo   *string
	Short  *string
	Reqs   []*langcrud.Req
	Conds  *langcrud.Conds
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

func WithLang(lang *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lang == nil {
			if must {
				return fmt.Errorf("invalid lang")
			}
			return nil
		}
		if *lang == "" {
			return fmt.Errorf("invalid lang")
		}
		h.Lang = lang
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid name")
		}
		h.Name = name
		return nil
	}
}

func WithLogo(logo *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if logo == nil {
			if must {
				return fmt.Errorf("invalid logo")
			}
			return nil
		}
		if *logo == "" {
			return fmt.Errorf("invalid logo")
		}
		h.Logo = logo
		return nil
	}
}

func WithShort(short *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if short == nil {
			if must {
				return fmt.Errorf("invalid short")
			}
			return nil
		}
		if *short == "" {
			return fmt.Errorf("invalid short")
		}
		h.Short = short
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &langcrud.Conds{}
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
		if conds.Lang != nil {
			h.Conds.Lang = &cruder.Cond{Op: conds.GetLang().GetOp(), Val: conds.GetLang().GetValue()}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{Op: conds.GetName().GetOp(), Val: conds.GetName().GetValue()}
		}
		if conds.Short != nil {
			h.Conds.Short = &cruder.Cond{Op: conds.GetShort().GetOp(), Val: conds.GetShort().GetValue()}
		}
		if len(conds.GetLangs().GetValue()) > 0 {
			h.Conds.Langs = &cruder.Cond{Op: conds.GetLangs().GetOp(), Val: conds.GetLangs().GetValue()}
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

//nolint:gocyclo
func WithReqs(reqs []*npool.LangReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*langcrud.Req{}
		for _, req := range reqs {
			if must {
				if req.Lang == nil {
					return fmt.Errorf("invalid lang")
				}
				if req.Name == nil {
					return fmt.Errorf("invalid name")
				}
				if req.Logo == nil {
					return fmt.Errorf("invalid logo")
				}
				if req.Short == nil {
					return fmt.Errorf("invalid short")
				}
			}
			_req := &langcrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.Lang != nil {
				if *req.Lang == "" {
					return fmt.Errorf("invalid lang")
				}
				_req.Lang = req.Lang
			}
			if req.Name != nil {
				if *req.Name == "" {
					return fmt.Errorf("invalid name")
				}
				_req.Name = req.Name
			}
			if req.Logo != nil {
				if *req.Logo == "" {
					return fmt.Errorf("invalid logo")
				}
				_req.Logo = req.Logo
			}
			if req.Short != nil {
				if *req.Short == "" {
					return fmt.Errorf("invalid short")
				}
				_req.Short = req.Short
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
