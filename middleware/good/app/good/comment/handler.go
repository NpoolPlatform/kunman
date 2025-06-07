package comment

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constant "github.com/NpoolPlatform/kunman/middleware/good/const"
	commentcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/comment"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	scorecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/score"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/app/good/goodbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/comment"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	commentcrud.Req
	ScoreReq         *scorecrud.Req
	CommentConds     *commentcrud.Conds
	AppGoodBaseConds *appgoodbasecrud.Conds
	GoodBaseConds    *goodbasecrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		ScoreReq:         &scorecrud.Req{},
		CommentConds:     &commentcrud.Conds{},
		AppGoodBaseConds: &appgoodbasecrud.Conds{},
		GoodBaseConds:    &goodbasecrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
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
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		h.ScoreReq.CommentID = &_id
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.UserID = &_id
		h.ScoreReq.UserID = &_id
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := appgoodbase1.NewHandler(
			ctx,
			appgoodbase1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistGoodBase(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid appgood")
		}
		h.AppGoodID = handler.EntID
		h.ScoreReq.AppGoodID = handler.EntID
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = &_id
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
		handler, err := NewHandler(
			ctx,
			WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistComment(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid replytoid")
		}
		h.ReplyToID = handler.EntID
		return nil
	}
}

func WithAnonymous(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Anonymous = b
		return nil
	}
}

func WithPurchasedUser(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.PurchasedUser = b
		return nil
	}
}

func WithTrialUser(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.TrialUser = b
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
		h.ScoreReq.Score = &score
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

//nolint:gocyclo
func (h *Handler) withCommentConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.CommentConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.CommentConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.UserID != nil {
		id, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.CommentConds.UserID = &cruder.Cond{
			Op: conds.GetUserID().GetOp(), Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.CommentConds.AppGoodID = &cruder.Cond{
			Op: conds.GetAppGoodID().GetOp(), Val: id,
		}
	}
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.CommentConds.OrderID = &cruder.Cond{
			Op: conds.GetOrderID().GetOp(), Val: id,
		}
	}
	if conds.AppGoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetAppGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.CommentConds.AppGoodIDs = &cruder.Cond{
			Op: conds.GetAppGoodIDs().GetOp(), Val: ids,
		}
	}
	if conds.OrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetOrderIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.CommentConds.OrderIDs = &cruder.Cond{
			Op: conds.GetOrderIDs().GetOp(), Val: ids,
		}
	}
	return nil
}

func (h *Handler) withAppGoodBaseConds(conds *npool.Conds) error {
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.AppID = &cruder.Cond{
			Op: conds.GetAppID().GetOp(), Val: id,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.GoodID = &cruder.Cond{
			Op: conds.GetGoodID().GetOp(), Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.EntID = &cruder.Cond{
			Op: conds.GetAppGoodID().GetOp(), Val: id,
		}
	}
	if conds.AppGoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetAppGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.AppGoodBaseConds.EntIDs = &cruder.Cond{
			Op: conds.GetAppGoodIDs().GetOp(), Val: ids,
		}
	}
	return nil
}

func (h *Handler) withGoodBaseConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodBaseConds.EntID = &cruder.Cond{
			Op: conds.GetGoodID().GetOp(), Val: id,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withCommentConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withAppGoodBaseConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withGoodBaseConds(conds)
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
