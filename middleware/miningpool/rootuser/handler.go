package rootuser

import (
	"context"
	"net/mail"
	"regexp"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/rootuser"
	rootusercrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/rootuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pool"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	rootusercrud.Req
	AuthTokenPlain *string
	AuthTokenSalt  *string
	Reqs           []*rootusercrud.Req
	Conds          *rootusercrud.Conds
	Offset         int32
	Limit          int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid id")
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
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

func WithPoolID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid poolid")
			}
			return nil
		}
		poolH, err := pool.NewHandler(ctx, pool.WithEntID(id, true))
		if err != nil {
			return wlog.WrapError(err)
		}

		exist, err := poolH.ExistPool(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid poolid")
		}
		h.PoolID = poolH.EntID
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return wlog.Errorf("invalid name")
			}
			return nil
		}
		re := regexp.MustCompile("^[a-zA-Z0-9\u3040-\u31ff][[a-zA-Z0-9_\\-\\.\u3040-\u31ff]{3,32}$") //nolint
		if !re.MatchString(*name) {
			return wlog.Errorf("invalid name")
		}
		h.Name = name
		return nil
	}
}

func WithEmail(email *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if email == nil {
			if must {
				return wlog.Errorf("invalid email")
			}
			return nil
		}
		if _, err := mail.ParseAddress(*email); err != nil {
			return wlog.WrapError(err)
		}
		h.Email = email
		return nil
	}
}

// authtoken that passed in will be authtokenplain
// then h.authtokenplain is encrypted into h.authtoken with h.authtokensalt
// and will store h.authtoken and authtokensalt in db
func WithAuthToken(authtoken *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if authtoken == nil {
			if must {
				return wlog.Errorf("invalid authtoken")
			}
			return nil
		}
		h.AuthTokenPlain = authtoken
		return h.withAuthTokenEncrypt()
	}
}

func WithAuthed(authed *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if authed == nil {
			if must {
				return wlog.Errorf("invalid authed")
			}
			return nil
		}
		if !*authed {
			return wlog.Errorf("invalid authed")
		}
		h.Authed = authed
		return nil
	}
}

func WithRemark(remark *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if remark == nil {
			if must {
				return wlog.Errorf("invalid remark")
			}
			return nil
		}
		h.Remark = remark
		return nil
	}
}

//nolint:gocognit
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &rootusercrud.Conds{}
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
				return wlog.WrapError(err)
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.PoolID != nil {
			id, err := uuid.Parse(conds.GetPoolID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.PoolID = &cruder.Cond{
				Op:  conds.GetPoolID().GetOp(),
				Val: id,
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return wlog.WrapError(err)
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
			}
		}
		if conds.Email != nil {
			h.Conds.Email = &cruder.Cond{
				Op:  conds.GetEmail().GetOp(),
				Val: conds.GetEmail().GetValue(),
			}
		}
		if conds.Authed != nil {
			h.Conds.Authed = &cruder.Cond{
				Op:  conds.GetAuthed().GetOp(),
				Val: conds.GetAuthed().GetValue(),
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
