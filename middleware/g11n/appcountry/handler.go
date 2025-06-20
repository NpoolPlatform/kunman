package appcountry

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/appcountry"
	appcountrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/appcountry"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	CountryID *uuid.UUID
	Reqs      []*appcountrycrud.Req
	Conds     *appcountrycrud.Conds
	Offset    int32
	Limit     int32
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

func WithCountryID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid countryid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CountryID = &_id
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &appcountrycrud.Conds{}
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
		if conds.CountryID != nil {
			id, err := uuid.Parse(conds.GetCountryID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CountryID = &cruder.Cond{Op: conds.GetCountryID().GetOp(), Val: id}
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

func WithReqs(reqs []*npool.CountryReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*appcountrycrud.Req{}
		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.CountryID == nil {
					return fmt.Errorf("invalid countryid")
				}
			}
			_req := &appcountrycrud.Req{}
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
			if req.CountryID != nil {
				id, err := uuid.Parse(*req.CountryID)
				if err != nil {
					return err
				}
				_req.CountryID = &id
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
