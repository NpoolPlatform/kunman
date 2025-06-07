package like

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constant "github.com/NpoolPlatform/kunman/middleware/good/const"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	likecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/like"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/app/good/goodbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/like"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	likecrud.Req
	LikeConds        *likecrud.Conds
	AppGoodBaseConds *appgoodbasecrud.Conds
	GoodBaseConds    *goodbasecrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		LikeConds:        &likecrud.Conds{},
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
			return err
		}
		h.EntID = &_id
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
			return err
		}
		h.UserID = &_id
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
		handler, err := appgoodbase1.NewHandler(
			ctx,
			appgoodbase1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.Errorf("invalid appgoodid")
		}
		exist, err := handler.ExistGoodBase(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return wlog.Errorf("invalid app good")
		}
		h.AppGoodID = handler.EntID
		return nil
	}
}

func WithLike(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if b == nil {
			if must {
				return wlog.Errorf("invalid like")
			}
			return nil
		}
		h.Like = b
		return nil
	}
}

func (h *Handler) withLikeConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.LikeConds.ID = &cruder.Cond{
			Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.LikeConds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
	}
	if conds.UserID != nil {
		id, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return err
		}
		h.LikeConds.UserID = &cruder.Cond{
			Op: conds.GetUserID().GetOp(), Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return err
		}
		h.LikeConds.AppGoodID = &cruder.Cond{
			Op: conds.GetAppGoodID().GetOp(), Val: id,
		}
	}
	if conds.AppGoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetAppGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			ids = append(ids, _id)
		}
		h.LikeConds.AppGoodIDs = &cruder.Cond{
			Op: conds.GetAppGoodIDs().GetOp(), Val: ids,
		}
	}
	return nil
}

func (h *Handler) withAppGoodBaseConds(conds *npool.Conds) error {
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return err
		}
		h.AppGoodBaseConds.AppID = &cruder.Cond{
			Op: conds.GetAppID().GetOp(), Val: id,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return err
		}
		h.AppGoodBaseConds.GoodID = &cruder.Cond{
			Op: conds.GetGoodID().GetOp(), Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return err
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
				return err
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
			return err
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
		if err := h.withLikeConds(conds); err != nil {
			return err
		}
		if err := h.withAppGoodBaseConds(conds); err != nil {
			return err
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
