package message

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/message"
	messagecrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/message"
	lang1 "github.com/NpoolPlatform/kunman/middleware/g11n/lang"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	MessageID *string
	Message   *string
	GetIndex  *uint32
	Disabled  *bool
	Reqs      []*messagecrud.Req
	Conds     *messagecrud.Conds
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
		handler, err := lang1.NewHandler(
			ctx,
			lang1.WithEntID(id, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistLang(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid lang")
		}
		h.LangID = &_id
		return nil
	}
}

func WithMessageID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid messageid")
			}
			return nil
		}
		if *id == "" {
			return fmt.Errorf("invalid messageid")
		}
		h.MessageID = id
		return nil
	}
}

func WithMessage(message *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if message == nil {
			if must {
				return fmt.Errorf("invalid message")
			}
			return nil
		}
		h.Message = message
		return nil
	}
}

func WithGetIndex(getindex *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if getindex == nil {
			if must {
				return fmt.Errorf("invalid getindex")
			}
			return nil
		}
		h.GetIndex = getindex
		return nil
	}
}

func WithDisabled(disabled *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if disabled == nil {
			if must {
				return fmt.Errorf("invalid disabled")
			}
			return nil
		}
		h.Disabled = disabled
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &messagecrud.Conds{}
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
		if conds.MessageID != nil {
			h.Conds.MessageID = &cruder.Cond{Op: conds.GetMessageID().GetOp(), Val: conds.GetMessageID().GetValue()}
		}
		if conds.Disabled != nil {
			h.Conds.Disabled = &cruder.Cond{Op: conds.GetDisabled().GetOp(), Val: conds.GetDisabled().GetValue()}
		}
		if len(conds.GetMessageIDs().GetValue()) > 0 {
			h.Conds.MessageIDs = &cruder.Cond{Op: conds.GetMessageIDs().GetOp(), Val: conds.GetMessageIDs().GetValue()}
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
func WithReqs(reqs []*npool.MessageReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*messagecrud.Req{}
		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.LangID == nil {
					return fmt.Errorf("invalid langid")
				}
				if req.MessageID == nil {
					return fmt.Errorf("invalid messageid")
				}
				if req.Message == nil {
					return fmt.Errorf("invalid message")
				}
			}
			_req := &messagecrud.Req{}
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
			if req.MessageID != nil {
				if *req.MessageID == "" {
					return fmt.Errorf("invalid messageid")
				}
				_req.MessageID = req.MessageID
			}
			if req.Message != nil {
				if *req.Message == "" {
					return fmt.Errorf("invalid message")
				}
				_req.Message = req.Message
			}
			if req.GetIndex != nil {
				_req.GetIndex = req.GetIndex
			}
			if req.Disabled != nil {
				_req.Disabled = req.Disabled
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
