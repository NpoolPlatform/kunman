package oauththirdparty

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	thidpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/oauththirdparty"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/oauth/oauththirdparty"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID             *uint32
	EntID          *uuid.UUID
	ClientName     *basetypes.SignMethod
	ClientTag      *string
	ClientLogoURL  *string
	ClientOAuthURL *string
	ResponseType   *string
	Scope          *string
	Conds          *thidpartycrud.Conds
	Offset         int32
	Limit          int32
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

func WithClientName(clientName *basetypes.SignMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientName == nil {
			if must {
				return fmt.Errorf("invalid clientname")
			}
			return nil
		}
		switch *clientName {
		case basetypes.SignMethod_Twitter:
		case basetypes.SignMethod_Github:
		case basetypes.SignMethod_Facebook:
		case basetypes.SignMethod_Linkedin:
		case basetypes.SignMethod_Wechat:
		case basetypes.SignMethod_Google:
		default:
			return fmt.Errorf("invalid clientname")
		}
		h.ClientName = clientName
		return nil
	}
}

func WithClientTag(clientTag *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientTag == nil {
			if must {
				return fmt.Errorf("invalid clienttag")
			}
			return nil
		}
		if *clientTag == "" {
			return fmt.Errorf("invalid clienttag")
		}
		h.ClientTag = clientTag
		return nil
	}
}

func WithClientLogoURL(clientLogoURL *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientLogoURL == nil {
			if must {
				return fmt.Errorf("invalid clientlogourl")
			}
			return nil
		}
		if *clientLogoURL == "" {
			return fmt.Errorf("invalid clientlogourl")
		}
		h.ClientLogoURL = clientLogoURL
		return nil
	}
}

func WithClientOAuthURL(clientOAuthURL *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientOAuthURL == nil {
			if must {
				return fmt.Errorf("invalid clientoauthurl")
			}
			return nil
		}
		if *clientOAuthURL == "" {
			return fmt.Errorf("invalid clientoauthurl")
		}
		h.ClientOAuthURL = clientOAuthURL
		return nil
	}
}

func WithResponseType(responseType *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if responseType == nil {
			if must {
				return fmt.Errorf("invalid responsetype")
			}
			return nil
		}
		if *responseType == "" {
			return fmt.Errorf("invalid responsetype")
		}
		h.ResponseType = responseType
		return nil
	}
}

func WithScope(scope *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if scope == nil {
			if must {
				return fmt.Errorf("invalid scope")
			}
			return nil
		}
		h.Scope = scope
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &thidpartycrud.Conds{}
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

		if conds.ClientName != nil {
			h.Conds.ClientName = &cruder.Cond{
				Op:  conds.GetClientName().GetOp(),
				Val: basetypes.SignMethod(conds.GetClientName().GetValue()),
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
