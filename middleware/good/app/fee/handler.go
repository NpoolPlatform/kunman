package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constant "github.com/NpoolPlatform/kunman/middleware/good/const"
	appfeecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/fee"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	feecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/fee"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/good/goodbase"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	appfeecrud.Req
	AppGoodBaseReq   *appgoodbasecrud.Req
	AppFeeConds      *appfeecrud.Conds
	FeeConds         *feecrud.Conds
	GoodBaseConds    *goodbasecrud.Conds
	AppGoodBaseConds *appgoodbasecrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	valTrue := true
	handler := &Handler{
		AppGoodBaseReq: &appgoodbasecrud.Req{
			Purchasable:       &valTrue,
			EnableProductPage: &valTrue,
			Online:            &valTrue,
			Visible:           &valTrue,
		},
		AppFeeConds:      &appfeecrud.Conds{},
		FeeConds:         &feecrud.Conds{},
		GoodBaseConds:    &goodbasecrud.Conds{},
		AppGoodBaseConds: &appgoodbasecrud.Conds{},
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

func WithEntID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &id
		return nil
	}
}

func WithAppID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseReq.AppID = &id
		return nil
	}
}

func WithGoodID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		handler, _ := goodbase1.NewHandler(ctx)
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.GoodBaseConds = &goodbasecrud.Conds{
			EntID: &cruder.Cond{Op: cruder.EQ, Val: id},
			GoodTypes: &cruder.Cond{Op: cruder.IN, Val: []types.GoodType{
				types.GoodType_TechniqueServiceFee,
				types.GoodType_ElectricityFee,
			}},
		}
		exist, err := handler.ExistGoodBaseConds(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid goodid")
		}
		h.AppGoodBaseReq.GoodID = &id
		return nil
	}
}

func WithAppGoodID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid appgoodid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodID = &id
		h.AppGoodBaseReq.EntID = &id
		return nil
	}
}

func WithProductPage(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodBaseReq.ProductPage = s
		return nil
	}
}

func WithName(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid name")
			}
			return nil
		}
		const leastNameLen = 3
		if len(*s) < leastNameLen {
			return wlog.Errorf("invalid name")
		}
		h.AppGoodBaseReq.Name = s
		return nil
	}
}

func WithBanner(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodBaseReq.Banner = s
		return nil
	}
}

func WithUnitValue(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid unitvalue")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
			return wlog.Errorf("invalid unitvalue")
		}
		h.UnitValue = &amount
		return nil
	}
}

func WithCancelMode(e *types.CancelMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid cancelmode")
			}
			return nil
		}
		switch *e {
		case types.CancelMode_Uncancellable:
		case types.CancelMode_CancellableBeforeUsed:
		default:
			return wlog.Errorf("invalid cancelmode")
		}
		h.CancelMode = e
		return nil
	}
}

func WithMinOrderDurationSeconds(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return wlog.Errorf("invalid minorderduration")
			}
			return nil
		}
		if *n == 0 {
			return wlog.Errorf("invalid minorderduration")
		}
		h.MinOrderDurationSeconds = n
		return nil
	}
}

func (h *Handler) withFeeConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.FeeConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.FeeConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.SettlementType != nil {
		h.FeeConds.SettlementType = &cruder.Cond{
			Val: types.GoodSettlementType(conds.GetSettlementType().GetValue()),
			Op:  conds.GetSettlementType().GetOp(),
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
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.GoodBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.GoodType != nil {
		h.GoodBaseConds.GoodType = &cruder.Cond{
			Op:  conds.GetGoodType().GetOp(),
			Val: types.GoodType(conds.GetGoodType().GetValue()),
		}
	}
	return nil
}

func (h *Handler) withAppFeeConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.AppFeeConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.IDs != nil {
		h.AppFeeConds.IDs = &cruder.Cond{
			Op:  conds.GetIDs().GetOp(),
			Val: conds.GetIDs().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppFeeConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.EntIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetEntIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.AppFeeConds.EntIDs = &cruder.Cond{
			Op:  conds.GetEntIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppFeeConds.AppGoodID = &cruder.Cond{
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
		h.AppFeeConds.AppGoodIDs = &cruder.Cond{
			Op:  conds.GetAppGoodIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

//nolint:gocyclo
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
	if conds.AppIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetAppIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.AppGoodBaseConds.AppIDs = &cruder.Cond{
			Op:  conds.GetAppIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.AppGoodBaseConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
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
		if err := h.withFeeConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withGoodBaseConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withAppFeeConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withAppGoodBaseConds(conds); err != nil {
			return wlog.WrapError(err)
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
