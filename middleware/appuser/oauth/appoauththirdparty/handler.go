package appoauththirdparty

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/middleware/appuser/const"
	appoauththirdpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/appoauththirdparty"
	app "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/oauth/appoauththirdparty"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID            *uint32
	EntID         *uuid.UUID
	AppID         uuid.UUID
	ThirdPartyID  *uuid.UUID
	ClientID      *string
	ClientSecret  *string
	CallbackURL   *string
	Salt          *string
	ThirdPartyIDs []*uuid.UUID
	Reqs          []*appoauththirdpartycrud.Req
	Conds         *appoauththirdpartycrud.Conds
	Offset        int32
	Limit         int32
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
		h.AppID = _id
		return nil
	}
}

func WithThirdPartyID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid thirdpartyid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ThirdPartyID = &_id
		return nil
	}
}

func WithClientID(clientID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientID == nil {
			if must {
				return fmt.Errorf("invalid clientid")
			}
			return nil
		}
		if *clientID == "" {
			return fmt.Errorf("invalid clientid")
		}
		h.ClientID = clientID
		return nil
	}
}

func WithClientSecret(clientSecret *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientSecret == nil {
			if must {
				return fmt.Errorf("invalid clientsecret")
			}
			return nil
		}
		if *clientSecret == "" {
			return fmt.Errorf("invalid clientsecret")
		}
		h.ClientSecret = clientSecret
		return nil
	}
}

func WithSalt(salt *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if salt == nil {
			if must {
				return fmt.Errorf("invalid salt")
			}
			return nil
		}
		h.Salt = salt
		return nil
	}
}

func WithCallbackURL(callbackURL *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if callbackURL == nil {
			return nil
		}
		if *callbackURL == "" {
			return fmt.Errorf("invalid callbackurl")
		}
		h.CallbackURL = callbackURL
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &appoauththirdpartycrud.Conds{}
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
		if conds.ThirdPartyID != nil {
			id, err := uuid.Parse(conds.GetThirdPartyID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ThirdPartyID = &cruder.Cond{Op: conds.GetThirdPartyID().GetOp(), Val: id}
		}
		if conds.ClientName != nil {
			h.Conds.ClientName = &cruder.Cond{
				Op:  conds.GetClientName().GetOp(),
				Val: basetypes.SignMethod(conds.GetClientName().GetValue()),
			}
		}
		if conds.DecryptSecret != nil {
			h.Conds.DecryptSecret = &cruder.Cond{
				Op:  conds.GetDecryptSecret().GetOp(),
				Val: conds.GetDecryptSecret().GetValue(),
			}
		}
		if len(conds.GetThirdPartyIDs().GetValue()) > 0 {
			_ids := []uuid.UUID{}
			for _, id := range conds.GetThirdPartyIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				_ids = append(_ids, _id)
			}
			h.Conds.ThirdPartyIDs = &cruder.Cond{Op: conds.GetThirdPartyIDs().GetOp(), Val: _ids}
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

func WithReqs(reqs []*npool.OAuthThirdPartyReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			_req := &appoauththirdpartycrud.Req{
				ClientID:     h.ClientID,
				ClientSecret: h.ClientSecret,
				CallbackURL:  h.CallbackURL,
			}
			if req.AppID == nil {
				return fmt.Errorf("invalid appid")
			}
			appID, err := uuid.Parse(*req.AppID)
			if err != nil {
				return err
			}
			_req.AppID = &appID
			if req.ThirdPartyID == nil {
				return fmt.Errorf("invalid thirdpartyid")
			}
			thirdPartyID, err := uuid.Parse(*req.ThirdPartyID)
			if err != nil {
				return err
			}
			_req.ThirdPartyID = &thirdPartyID
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
