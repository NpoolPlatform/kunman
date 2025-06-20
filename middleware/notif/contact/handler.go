package contact

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/contact"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/contact"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	UsedFor     *basetypes.UsedFor
	AccountType *basetypes.SignMethod
	Account     *string
	Sender      *string
	Conds       *crud.Conds
	Offset      int32
	Limit       int32
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, handler); err != nil {
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

func WithAccount(account *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if account == nil {
			if must {
				return fmt.Errorf("invalid account")
			}
			return nil
		}
		if *account == "" {
			return fmt.Errorf("account is empty")
		}
		h.Account = account
		return nil
	}
}

func WithSender(sender *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if sender == nil {
			if must {
				return fmt.Errorf("invalid sender")
			}
			return nil
		}
		if *sender == "" {
			return fmt.Errorf("sender is empty")
		}
		h.Sender = sender
		return nil
	}
}

func WithUsedFor(usedFor *basetypes.UsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if usedFor == nil {
			if must {
				return fmt.Errorf("invalid usedfor")
			}
			return nil
		}
		switch *usedFor {
		case basetypes.UsedFor_Contact:
		default:
			return fmt.Errorf("used for %v invalid", *usedFor)
		}
		h.UsedFor = usedFor
		return nil
	}
}

func WithAccountType(_type *basetypes.SignMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid accounttype")
			}
			return nil
		}
		switch *_type {
		case basetypes.SignMethod_Email:
		default:
			return fmt.Errorf("type %v invalid", *_type)
		}
		h.AccountType = _type
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op: conds.GetEntID().GetOp(), Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.AccountType != nil {
			accountType := conds.GetAccountType().GetValue()
			h.Conds.AccountType = &cruder.Cond{Op: conds.GetAccountType().GetOp(), Val: basetypes.SignMethod(accountType)}
		}
		if conds.UsedFor != nil {
			usedFor := conds.GetUsedFor().GetValue()
			h.Conds.UsedFor = &cruder.Cond{Op: conds.GetUsedFor().GetOp(), Val: basetypes.UsedFor(usedFor)}
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
