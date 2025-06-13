package statement

import (
	"context"
	"encoding/json"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/simulate/ledger/statement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/simulate/ledger/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	crud.Req
	Rollback *bool
	Reqs     []*crud.Req
	StartAt  uint32
	EndAT    uint32
	Conds    *crud.Conds
	Offset   int32
	Limit    int32
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
				return fmt.Errorf("invalid app id")
			}
			return nil
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
				return fmt.Errorf("invalid user id")
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

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid coin type id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

//nolint
func WithIOType(_type *basetypes.IOType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid io type")
			}
			return nil
		}
		flag := false
		for ioType := range basetypes.IOType_value {
			if ioType == _type.String() && ioType != basetypes.IOType_DefaultType.String() {
				flag = true
			}
		}
		if !flag {
			return fmt.Errorf("invalid io type %v", *_type)
		}
		h.IOType = _type
		return nil
	}
}

//nolint
func WithIOSubType(_type *basetypes.IOSubType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid io sub type")
			}
			return nil
		}
		flag := false
		for ioSubType := range basetypes.IOSubType_value {
			if ioSubType == _type.String() && ioSubType != basetypes.IOSubType_DefaultSubType.String() {
				flag = true
			}
		}
		if !flag {
			return fmt.Errorf("invalid io sub type %v", *_type)
		}
		h.IOSubType = _type
		return nil
	}
}

func WithAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid amount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("amount is less than 0 %v", *amount)
		}
		h.Amount = &_amount
		return nil
	}
}

func WithIOExtra(extra *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if extra == nil {
			if must {
				return fmt.Errorf("invalid io extra")
			}
			return nil
		}
		if !json.Valid([]byte(*extra)) {
			return fmt.Errorf("io extra is invalid json str %v", *extra)
		}

		h.IOExtra = extra
		return nil
	}
}

func WithCreatedAt(createdAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if createdAt == nil {
			if must {
				return fmt.Errorf("invalid created at")
			}
			return nil
		}
		if *createdAt == 0 {
			return fmt.Errorf("invalid created at 0")
		}
		h.CreatedAt = createdAt
		return nil
	}
}

func WithStartAt(startAt uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.StartAt = startAt
		return nil
	}
}

func WithEndAt(endAt uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EndAT = endAt
		return nil
	}
}

func WithRollback(rollback *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if rollback == nil {
			if must {
				return fmt.Errorf("invalid rollback")
			}
			return nil
		}
		h.Rollback = rollback
		return nil
	}
}

func WithSendCoupon(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid sendcoupon")
			}
			return nil
		}
		h.SendCoupon = value
		return nil
	}
}

func WithCashable(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid cashable")
			}
			return nil
		}
		h.Cashable = value
		return nil
	}
}

//nolint
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
				Val: id,
			}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{
				Op:  conds.GetCoinTypeID().GetOp(),
				Val: id,
			}
		}
		if conds.IOType != nil {
			ioType := conds.GetIOType().GetValue()
			h.Conds.IOType = &cruder.Cond{
				Op:  conds.GetIOType().GetOp(),
				Val: basetypes.IOType(ioType),
			}
		}
		if conds.IOSubType != nil {
			ioSubType := conds.GetIOSubType().GetValue()
			h.Conds.IOSubType = &cruder.Cond{
				Op:  conds.GetIOSubType().GetOp(),
				Val: basetypes.IOSubType(ioSubType),
			}
		}
		if conds.IOExtra != nil {
			h.Conds.IOExtra = &cruder.Cond{
				Op:  conds.GetIOExtra().GetOp(),
				Val: conds.GetIOExtra().GetValue(),
			}
		}
		if conds.StartAt != nil {
			h.Conds.StartAt = &cruder.Cond{
				Op:  conds.GetStartAt().GetOp(),
				Val: conds.GetStartAt().GetValue(),
			}
		}
		if conds.EndAt != nil {
			h.Conds.EndAt = &cruder.Cond{
				Op:  conds.GetEndAt().GetOp(),
				Val: conds.GetEndAt().GetValue(),
			}
		}
		if len(conds.GetIOSubTypes().GetValue()) > 0 {
			ioSubTypes := []string{}
			for _, val := range conds.GetIOSubTypes().GetValue() {
				ioSubTypes = append(ioSubTypes, basetypes.IOSubType_name[int32(val)])
			}
			h.Conds.IOSubTypes = &cruder.Cond{Op: conds.GetIOSubTypes().GetOp(), Val: ioSubTypes}
		}
		if len(conds.GetCoinTypeIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, val := range conds.GetCoinTypeIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return err
				}
				ids = append(ids, id)
			}
			h.Conds.CoinTypeIDs = &cruder.Cond{
				Op:  conds.GetCoinTypeIDs().GetOp(),
				Val: ids,
			}
		}
		if len(conds.GetUserIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, val := range conds.GetUserIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return err
				}
				ids = append(ids, id)
			}
			h.Conds.UserIDs = &cruder.Cond{
				Op:  conds.GetUserIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.SendCoupon != nil {
			h.Conds.SendCoupon = &cruder.Cond{
				Op:  conds.GetSendCoupon().GetOp(),
				Val: conds.GetSendCoupon().GetValue(),
			}
		}
		if conds.Cashable != nil {
			h.Conds.Cashable = &cruder.Cond{
				Op:  conds.GetCashable().GetOp(),
				Val: conds.GetCashable().GetValue(),
			}
		}
		return nil
	}
}

//nolint
func WithReqs(reqs []*npool.StatementReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*crud.Req{}
		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid app id")
				}
				if req.UserID == nil {
					return fmt.Errorf("invalid user id")
				}
				if req.CoinTypeID == nil {
					return fmt.Errorf("invalid coin type id")
				}
				if req.Amount == nil {
					return fmt.Errorf("invalid amount")
				}
				if req.IOExtra == nil {
					return fmt.Errorf("invalid io extra")
				}
				if req.IOType == nil {
					return fmt.Errorf("invalid io type")
				}
				if req.IOSubType == nil {
					return fmt.Errorf("invalid io sub type")
				}
				if *req.IOType != basetypes.IOType_Incoming {
					return fmt.Errorf("invalid io type %v", *req.IOType)
				}
				if *req.IOSubType != basetypes.IOSubType_MiningBenefit {
					return fmt.Errorf("io subtype not match io type")
				}
			}

			_req := &crud.Req{}
			if req.ID != nil {
				_req.ID = req.ID
			}
			if req.EntID != nil {
				_id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &_id
			}
			if req.AppID != nil {
				_id, err := uuid.Parse(*req.AppID)
				if err != nil {
					return err
				}
				_req.AppID = &_id
			}
			if req.UserID != nil {
				_id, err := uuid.Parse(*req.UserID)
				if err != nil {
					return err
				}
				_req.UserID = &_id
			}
			if req.CoinTypeID != nil {
				_id, err := uuid.Parse(*req.CoinTypeID)
				if err != nil {
					return err
				}
				_req.CoinTypeID = &_id
			}
			if req.Amount != nil {
				amount, err := decimal.NewFromString(*req.Amount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("amount is less than 0 %v", *req.Amount)
				}
				_req.Amount = &amount
			}
			if req.IOExtra != nil {
				if !json.Valid([]byte(*req.IOExtra)) {
					return fmt.Errorf("io extra is invalid json str %v", *req.IOExtra)
				}
				_req.IOExtra = req.IOExtra
			}
			if req.CreatedAt != nil {
				if *req.CreatedAt == 0 {
					return fmt.Errorf("invalid created at %v", *req.CreatedAt)
				}
				_req.CreatedAt = req.CreatedAt
			}
			if req.IOType != nil {
				_req.IOType = req.IOType
			}
			if req.IOSubType != nil {
				_req.IOSubType = req.IOSubType
			}
			if req.SendCoupon != nil {
				_req.SendCoupon = req.SendCoupon
			}
			if req.Cashable != nil {
				_req.Cashable = req.Cashable
			}
			if req.Rollback != nil {
				h.Rollback = req.Rollback
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
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
