package history

import (
	"context"
	"fmt"
	"net"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	historycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user/login/history"
	app "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user/login/history"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	ClientIP  *string
	UserAgent *string
	Location  *string
	LoginType *basetypes.LoginType
	Conds     *historycrud.Conds
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
		h.AppID = &_id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithClientIP(ip *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if ip == nil {
			if must {
				return fmt.Errorf("invalid clientip")
			}
			return nil
		}
		if ip := net.ParseIP(*ip); ip == nil {
			return fmt.Errorf("invalid client ip")
		}
		h.ClientIP = ip
		return nil
	}
}

func WithUserAgent(agent *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.UserAgent = agent
		return nil
	}
}

func WithLocation(location *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Location = location
		return nil
	}
}

func WithLoginType(loginType *basetypes.LoginType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if loginType == nil {
			if must {
				return fmt.Errorf("invalid logintype")
			}
			return nil
		}
		switch *loginType {
		case basetypes.LoginType_FreshLogin:
		case basetypes.LoginType_RefreshLogin:
		default:
			return fmt.Errorf("invalid login type")
		}
		h.LoginType = loginType
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &historycrud.Conds{}
		if conds == nil {
			return nil
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
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.LoginType != nil {
			h.Conds.LoginType = &cruder.Cond{
				Op:  conds.GetLoginType().GetOp(),
				Val: basetypes.LoginType(conds.GetLoginType().GetValue()),
			}
		}
		if conds.ClientIP != nil {
			h.Conds.ClientIP = &cruder.Cond{
				Op:  conds.GetClientIP().GetOp(),
				Val: conds.GetClientIP().GetValue(),
			}
		}
		if conds.Location != nil {
			h.Conds.Location = &cruder.Cond{
				Op:  conds.GetLocation().GetOp(),
				Val: conds.GetLocation().GetValue(),
			}
		}
		if conds.UserAgent != nil {
			h.Conds.UserAgent = &cruder.Cond{
				Op:  conds.GetUserAgent().GetOp(),
				Val: conds.GetUserAgent().GetValue(),
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
