package event

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"
	eventcoinmw "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event/coin"
	eventcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	eventcrud.Req
	Coins           []*eventcoinmw.EventCoinReq
	RemoveCouponIDs *bool
	RemoveCoins     *bool
	Consecutive     *uint32
	Amount          *decimal.Decimal
	UserID          *uuid.UUID
	Conds           *eventcrud.Conds
	Offset          int32
	Limit           int32
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppID = &_id
		return nil
	}
}

//nolint:dupl
func WithCredits(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid credits")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid credits")
		}
		h.Credits = &_amount
		return nil
	}
}

//nolint:dupl
func WithCreditsPerUSD(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid creditsperusd")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid creditsperusd")
		}
		h.CreditsPerUSD = &_amount
		return nil
	}
}

func WithMaxConsecutive(value *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MaxConsecutive = value
		return nil
	}
}

func WithInviterLayers(value *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.InviterLayers = value
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
		return nil
	}
}

func WithEventType(eventType *basetypes.UsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if eventType == nil {
			if must {
				return wlog.Errorf("invalid eventtype")
			}
			return nil
		}

		switch *eventType {
		// Already implemented
		case basetypes.UsedFor_Signup:
			fallthrough //nolint
		case basetypes.UsedFor_AffiliateSignup:
			fallthrough //nolint
		case basetypes.UsedFor_Purchase:
			fallthrough //nolint
		case basetypes.UsedFor_AffiliatePurchase:
			fallthrough //nolint
		case basetypes.UsedFor_SimulateOrderProfit:
			fallthrough //nolint
		case basetypes.UsedFor_SetWithdrawAddress:
			fallthrough //nolint
		case basetypes.UsedFor_ConsecutiveLogin:
			fallthrough //nolint
		case basetypes.UsedFor_GoodSocialSharing:
			fallthrough //nolint
		case basetypes.UsedFor_FirstOrderCompleted:
			fallthrough //nolint
		case basetypes.UsedFor_SetAddress:
			fallthrough //nolint
		case basetypes.UsedFor_Set2FA:
			fallthrough //nolint
		case basetypes.UsedFor_FirstBenefit:
			fallthrough //nolint
		case basetypes.UsedFor_WriteComment:
			fallthrough //nolint
		case basetypes.UsedFor_WriteRecommend:
			fallthrough //nolint
		case basetypes.UsedFor_GoodScoring:
			fallthrough //nolint
		case basetypes.UsedFor_SubmitTicket:
			fallthrough //nolint
		case basetypes.UsedFor_IntallApp:
			fallthrough //nolint
		case basetypes.UsedFor_SetNFTAvatar:
			fallthrough //nolint
		case basetypes.UsedFor_SetPersonalImage:
			fallthrough //nolint
		case basetypes.UsedFor_Signin:
			fallthrough //nolint
		case basetypes.UsedFor_KYCApproved:
			fallthrough //nolint
		case basetypes.UsedFor_OrderCompleted:
			fallthrough //nolint
		case basetypes.UsedFor_WithdrawalCompleted:
			fallthrough //nolint
		case basetypes.UsedFor_DepositReceived:
			fallthrough //nolint
		case basetypes.UsedFor_UpdatePassword:
			fallthrough //nolint
		case basetypes.UsedFor_ResetPassword:
			fallthrough //nolint
		case basetypes.UsedFor_InternalTransfer:
		// Not implemented
		case basetypes.UsedFor_Update:
			fallthrough //nolint
		case basetypes.UsedFor_Contact:
			fallthrough //nolint
		case basetypes.UsedFor_Withdraw:
			fallthrough //nolint
		case basetypes.UsedFor_CreateInvitationCode:
			fallthrough //nolint
		case basetypes.UsedFor_SetCommission:
			fallthrough //nolint
		case basetypes.UsedFor_SetTransferTargetUser:
			fallthrough //nolint
		case basetypes.UsedFor_WithdrawalRequest:
			fallthrough //nolint
		case basetypes.UsedFor_KYCRejected:
			return wlog.Errorf("not implemented")
		default:
			return wlog.Errorf("invalid eventtype")
		}

		h.EventType = eventType
		return nil
	}
}

func WithGoodID(goodID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if goodID == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*goodID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = &_id
		return nil
	}
}

func WithAppGoodID(goodID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if goodID == nil {
			if must {
				return wlog.Errorf("invalid appgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*goodID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodID = &_id
		return nil
	}
}

func WithConsecutive(consecutive *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Consecutive = consecutive
		return nil
	}
}

func WithAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid amount")
			}
			_amount := decimal.NewFromInt(0)
			h.Amount = &_amount
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.Amount = &_amount
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &eventcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.EventType != nil {
			h.Conds.EventType = &cruder.Cond{
				Op:  conds.GetEventType().GetOp(),
				Val: basetypes.UsedFor(conds.GetEventType().GetValue()),
			}
		}
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.GoodID = &cruder.Cond{
				Op:  conds.GetGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.AppGoodID != nil {
			id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.AppGoodID = &cruder.Cond{
				Op:  conds.GetAppGoodID().GetOp(),
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
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: conds.GetID().GetValue(),
			}
		}
		return nil
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
