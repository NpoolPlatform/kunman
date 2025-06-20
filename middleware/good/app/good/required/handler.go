//nolint:dupl
package required

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/goodbase"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	requiredcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/required"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	requiredcrud.Req
	RequiredConds         *requiredcrud.Conds
	AppGoodBaseConds      *appgoodbasecrud.Conds
	RequiredGoodBaseConds *goodbasecrud.Conds
	Offset                int32
	Limit                 int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		RequiredConds:         &requiredcrud.Conds{},
		AppGoodBaseConds:      &appgoodbasecrud.Conds{},
		RequiredGoodBaseConds: &goodbasecrud.Conds{},
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

func WithMainAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
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
		h.MainAppGoodID = handler.EntID
		return nil
	}
}

func WithRequiredAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
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
		h.RequiredAppGoodID = handler.EntID
		return nil
	}
}

func WithMust(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Must = b
		return nil
	}
}

//nolint:gocyclo
func (h *Handler) withRequiredConds(conds *npool.Conds) error {
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.RequiredConds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
	}
	if conds.ID != nil {
		h.RequiredConds.ID = &cruder.Cond{
			Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
		}
	}
	if conds.MainAppGoodID != nil {
		id, err := uuid.Parse(conds.GetMainAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.RequiredConds.MainAppGoodID = &cruder.Cond{
			Op:  conds.GetMainAppGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.MainAppGoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetMainAppGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.RequiredConds.MainAppGoodIDs = &cruder.Cond{
			Op:  conds.GetMainAppGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.RequiredAppGoodID != nil {
		id, err := uuid.Parse(conds.GetRequiredAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.RequiredConds.RequiredAppGoodID = &cruder.Cond{
			Op:  conds.GetRequiredAppGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.RequiredConds.AppGoodID = &cruder.Cond{
			Op:  conds.GetAppGoodID().GetOp(),
			Val: id,
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
		h.RequiredConds.AppGoodIDs = &cruder.Cond{
			Op:  conds.GetAppGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.Must != nil {
		h.RequiredConds.Must = &cruder.Cond{
			Op:  conds.GetMust().GetOp(),
			Val: conds.GetMust().GetValue(),
		}
	}
	return nil
}

func (h *Handler) withAppGoodBaseConds(conds *npool.Conds) error {
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.EntID = &cruder.Cond{
			Op:  conds.GetAppGoodID().GetOp(),
			Val: id,
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
			Op:  conds.GetAppGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func (h *Handler) withRequiredGoodBaseConds(conds *npool.Conds) {
	if conds.RequiredGoodType != nil {
		h.RequiredGoodBaseConds.GoodType = &cruder.Cond{
			Op:  conds.GetRequiredGoodType().GetOp(),
			Val: types.GoodType(conds.GetRequiredGoodType().GetValue()),
		}
	}
	if conds.RequiredGoodTypes != nil {
		_types := []types.GoodType{}
		for _, _type := range conds.GetRequiredGoodTypes().GetValue() {
			_types = append(_types, types.GoodType(_type))
		}
		h.RequiredGoodBaseConds.GoodTypes = &cruder.Cond{
			Op:  conds.GetRequiredGoodTypes().GetOp(),
			Val: _types,
		}
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withAppGoodBaseConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		h.withRequiredGoodBaseConds(conds)
		return h.withRequiredConds(conds)
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
