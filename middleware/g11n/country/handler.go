package country

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	countrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/country"
	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/country"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID      *uint32
	EntID   *uuid.UUID
	Country *string
	Flag    *string
	Code    *string
	Short   *string
	Reqs    []*countrycrud.Req
	Conds   *countrycrud.Conds
	Offset  int32
	Limit   int32
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

func WithCountry(country *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if country == nil {
			if must {
				return fmt.Errorf("invalid country")
			}
			return nil
		}
		if *country == "" {
			return fmt.Errorf("invalid country")
		}
		h.Country = country
		return nil
	}
}

func WithFlag(flag *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if flag == nil {
			if must {
				return fmt.Errorf("invalid flag")
			}
			return nil
		}
		if *flag == "" {
			return fmt.Errorf("invalid flag")
		}
		h.Flag = flag
		return nil
	}
}

func WithCode(code *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if code == nil {
			if must {
				return fmt.Errorf("invalid code")
			}
			return nil
		}
		if *code == "" {
			return fmt.Errorf("invalid code")
		}
		h.Code = code
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
		h.Conds = &countrycrud.Conds{}
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
		if conds.Country != nil {
			h.Conds.Country = &cruder.Cond{Op: conds.GetCountry().GetOp(), Val: conds.GetCountry().GetValue()}
		}
		if conds.Code != nil {
			h.Conds.Code = &cruder.Cond{Op: conds.GetCode().GetOp(), Val: conds.GetCode().GetValue()}
		}
		if conds.Short != nil {
			h.Conds.Short = &cruder.Cond{Op: conds.GetShort().GetOp(), Val: conds.GetShort().GetValue()}
		}
		if len(conds.GetCountries().GetValue()) > 0 {
			h.Conds.Countries = &cruder.Cond{Op: conds.GetCountries().GetOp(), Val: conds.GetCountries().GetValue()}
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
func WithReqs(reqs []*npool.CountryReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*countrycrud.Req{}
		for _, req := range reqs {
			if must {
				if req.Country == nil {
					return fmt.Errorf("invalid country")
				}
				if req.Flag == nil {
					return fmt.Errorf("invalid flag")
				}
				if req.Code == nil {
					return fmt.Errorf("invalid code")
				}
				if req.Short == nil {
					return fmt.Errorf("invalid short")
				}
			}
			_req := &countrycrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.Country != nil {
				if *req.Country == "" {
					return fmt.Errorf("invalid country")
				}
				_req.Country = req.Country
			}
			if req.Flag != nil {
				if *req.Flag == "" {
					return fmt.Errorf("invalid flag")
				}
				_req.Flag = req.Flag
			}
			if req.Code != nil {
				if *req.Code == "" {
					return fmt.Errorf("invalid code")
				}
				_req.Code = req.Code
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
