package comment

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodcommon "github.com/NpoolPlatform/good-gateway/pkg/app/good/common"
	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	ordermwcli "github.com/NpoolPlatform/order-middleware/pkg/client/order"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID    *uint32
	EntID *string
	appgoodcommon.AppGoodCheckHandler
	OrderID       *string
	Content       *string
	ReplyToID     *string
	Anonymous     *bool
	Score         *string
	Hide          *bool
	HideReason    *types.GoodCommentHideReason
	CommentUserID *string
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
				return wlog.Errorf("invalid id")
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
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		exist, err := appmwcli.ExistApp(ctx, *id)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid app")
		}
		h.AppID = id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid userid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return wlog.WrapError(err)
		}
		h.UserID = id
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appgoodid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodID = id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderid")
			}
			return nil
		}
		exist, err := ordermwcli.ExistOrder(ctx, *id)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid order")
		}
		h.OrderID = id
		return nil
	}
}

func WithContent(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		const leastContentLen = 10
		if s == nil {
			if must {
				return wlog.Errorf("invalid content")
			}
			return nil
		}
		if len(*s) < leastContentLen {
			return wlog.Errorf("invalid content")
		}
		h.Content = s
		return nil
	}
}

func WithReplyToID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid replytoid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return wlog.WrapError(err)
		}
		h.ReplyToID = id
		return nil
	}
}

func WithAnonymous(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Anonymous = b
		return nil
	}
}

func WithScore(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid score")
			}
			return nil
		}
		score, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if score.Cmp(decimal.NewFromInt(0)) < 0 ||
			score.Cmp(decimal.NewFromInt(5)) > 0 { //nolint
			return wlog.Errorf("invalid score")
		}
		h.Score = s
		return nil
	}
}

func WithHide(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Hide = b
		return nil
	}
}

func WithHideReason(e *types.GoodCommentHideReason, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid hidereason")
			}
			return nil
		}
		switch *e {
		case types.GoodCommentHideReason_GoodCommentHideBySpam:
		case types.GoodCommentHideReason_GoodCommentHideByNotThisGood:
		case types.GoodCommentHideReason_GoodCommentHideByFalseDescription:
		default:
			return wlog.Errorf("invalid hidereason")
		}
		h.HideReason = e
		return nil
	}
}

func WithCommentUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid targetuserid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return wlog.WrapError(err)
		}
		h.CommentUserID = id
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
