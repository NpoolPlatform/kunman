package poster

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constant "github.com/NpoolPlatform/kunman/middleware/good/const"
	devicepostercrud "github.com/NpoolPlatform/kunman/middleware/good/crud/device/poster"
	device1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/device"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	devicepostercrud.Req
	PosterConds *devicepostercrud.Conds
	Offset      int32
	Limit       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		PosterConds: &devicepostercrud.Conds{},
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

func WithDeviceTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := device1.NewHandler(
			ctx,
			device1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistDeviceType(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid devicetype")
		}
		h.DeviceTypeID = handler.EntID
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

func (h *Handler) withAppPosterGoodConds(conds *npool.Conds) error {
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
	if conds.DeviceTypeID != nil {
		id, err := uuid.Parse(conds.GetDeviceTypeID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PosterConds.DeviceTypeID = &cruder.Cond{
			Op:  conds.GetDeviceTypeID().GetOp(),
			Val: id,
		}
	}
	if conds.DeviceTypeIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetDeviceTypeIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.PosterConds.DeviceTypeIDs = &cruder.Cond{
			Op:  conds.GetDeviceTypeIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		return h.withAppPosterGoodConds(conds)
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
