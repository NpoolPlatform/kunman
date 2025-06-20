package poster

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/poster"
	topmost1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost"
	topmostcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost"
	appgoodpostercrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost/poster"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	appgoodpostercrud.Req
	PosterConds  *appgoodpostercrud.Conds
	TopMostConds *topmostcrud.Conds
	Offset       int32
	Limit        int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		PosterConds:  &appgoodpostercrud.Conds{},
		TopMostConds: &topmostcrud.Conds{},
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
		return nil
	}
}

func WithTopMostID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := topmost1.NewHandler(
			ctx,
			topmost1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistTopMost(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid topmost")
		}
		h.TopMostID = handler.EntID
		return nil
	}
}

func WithPoster(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid poster")
			}
			return nil
		}
		h.Poster = s
		return nil
	}
}

func WithIndex(n *uint8, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Index = n
		return nil
	}
}

func (h *Handler) withTopMostPosterConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.PosterConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PosterConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.TopMostID != nil {
		id, err := uuid.Parse(conds.GetTopMostID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PosterConds.TopMostID = &cruder.Cond{
			Op:  conds.GetTopMostID().GetOp(),
			Val: id,
		}
	}
	if conds.TopMostIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetTopMostIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.PosterConds.TopMostIDs = &cruder.Cond{
			Op:  conds.GetTopMostIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withTopMostConds(conds *npool.Conds) error {
	if conds.TopMostID != nil {
		id, err := uuid.Parse(conds.GetTopMostID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.TopMostConds.EntID = &cruder.Cond{
			Op:  conds.GetTopMostID().GetOp(),
			Val: id,
		}
	}
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.TopMostConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withTopMostPosterConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withTopMostConds(conds)
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}
