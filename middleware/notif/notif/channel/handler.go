package channel

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/channel"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/channel"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Handler struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	Channel   *basetypes.NotifChannel
	EventType *basetypes.UsedFor
	Reqs      []*crud.Req
	Conds     *crud.Conds
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

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}

		h.AppID = &_id
		return nil
	}
}

func WithChannel(channel *basetypes.NotifChannel, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if channel == nil {
			if must {
				return fmt.Errorf("invalid channel")
			}
			return nil
		}
		switch *channel {
		case basetypes.NotifChannel_ChannelEmail:
		case basetypes.NotifChannel_ChannelSMS:
		case basetypes.NotifChannel_ChannelFrontend:
		default:
			return fmt.Errorf("channel %v invalid", *channel)
		}
		h.Channel = channel
		return nil
	}
}

//nolint:gocyclo
func WithEventType(_type *basetypes.UsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid eventtype")
			}
			return nil
		}
		switch *_type {
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Announcement:
		case basetypes.UsedFor_GoodBenefit1:
		case basetypes.UsedFor_UpdateEmail:
		case basetypes.UsedFor_UpdateMobile:
		case basetypes.UsedFor_UpdatePassword:
		case basetypes.UsedFor_UpdateGoogleAuth:
		case basetypes.UsedFor_NewLogin:
		case basetypes.UsedFor_OrderCompleted:
		case basetypes.UsedFor_OrderChildsRenewNotify:
		case basetypes.UsedFor_OrderChildsRenew:
		case basetypes.UsedFor_WithdrawReviewNotify:
		default:
			return fmt.Errorf("EventType is invalid")
		}

		h.EventType = _type
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			appID, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: appID}
		}
		if conds.Channel != nil {
			channel := conds.GetChannel().GetValue()
			h.Conds.Channel = &cruder.Cond{Op: conds.GetChannel().GetOp(), Val: basetypes.NotifChannel(channel)}
		}
		if conds.EventType != nil {
			_type := conds.GetEventType().GetValue()
			h.Conds.EventType = &cruder.Cond{Op: conds.GetEventType().GetOp(), Val: basetypes.UsedFor(_type)}
		}
		return nil
	}
}

//nolint:gocyclo
func WithReqs(reqs []*npool.ChannelReq, must bool) func(context.Context, *Handler) error {
	return func(_ctx context.Context, h *Handler) error {
		_reqs := []*crud.Req{}
		for _, req := range _reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.EventType == nil {
					return fmt.Errorf("invalid eventtype")
				}
				if req.Channel == nil {
					return fmt.Errorf("invalid channel")
				}
			}
			if req.AppID == nil || req.Channel == nil || req.EventType == nil {
				continue
			}
			// AppID
			_req := &crud.Req{}
			appID, err := uuid.Parse(req.AppID.String())
			if err != nil {
				return err
			}
			_req.AppID = &appID

			// EventType
			switch *req.EventType {
			case basetypes.UsedFor_WithdrawalRequest:
			case basetypes.UsedFor_WithdrawalCompleted:
			case basetypes.UsedFor_DepositReceived:
			case basetypes.UsedFor_KYCApproved:
			case basetypes.UsedFor_KYCRejected:
			case basetypes.UsedFor_Announcement:
			default:
				return fmt.Errorf("EventType is invalid %v", *req.EventType)
			}
			_req.EventType = req.EventType

			switch *req.Channel {
			case basetypes.NotifChannel_ChannelEmail:
			case basetypes.NotifChannel_ChannelSMS:
			case basetypes.NotifChannel_ChannelFrontend:
			default:
				return fmt.Errorf("channel %v invalid", *req.Channel)
			}
			_req.Channel = req.Channel

			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
