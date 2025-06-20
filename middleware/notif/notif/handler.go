package notif

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	templatemwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template"
	notifcrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID             *uint32
	EntID          *uuid.UUID
	AppID          *uuid.UUID
	UserID         *uuid.UUID
	LangID         *uuid.UUID
	EventID        *uuid.UUID
	Notified       *bool
	EventType      *basetypes.UsedFor
	UseTemplate    *bool
	Title          *string
	Content        *string
	Channel        *basetypes.NotifChannel
	Extra          *string
	NotifType      *basetypes.NotifType
	Vars           *templatemwpb.TemplateVars
	IDs            *[]uuid.UUID
	Reqs           []*notifcrud.Req
	MultiNotifReqs []*MultiNotifReq
	Conds          *notifcrud.Conds
	Offset         int32
	Limit          int32
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

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
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

func WithLangID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid langid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.LangID = &_id
		return nil
	}
}

func WithEventID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid eventid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EventID = &_id
		return nil
	}
}

func WithNotified(notified *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if notified == nil {
			if must {
				return fmt.Errorf("invalid notified")
			}
			return nil
		}
		h.Notified = notified
		return nil
	}
}

func WithUseTemplate(usetemplate *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if usetemplate == nil {
			if must {
				return fmt.Errorf("invalid usetemplate")
			}
			return nil
		}
		h.UseTemplate = usetemplate
		return nil
	}
}

func WithTitle(title *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if title == nil {
			if must {
				return fmt.Errorf("invalid title")
			}
			return nil
		}
		if *title == "" {
			return fmt.Errorf("invalid title")
		}
		h.Title = title
		return nil
	}
}

func WithContent(content *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if content == nil {
			if must {
				return fmt.Errorf("invalid content")
			}
			return nil
		}
		if *content == "" {
			return fmt.Errorf("invalid content")
		}
		h.Content = content
		return nil
	}
}

func WithChannel(_channel *basetypes.NotifChannel, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _channel == nil {
			if must {
				return fmt.Errorf("invalid channel")
			}
			return nil
		}
		switch *_channel {
		case basetypes.NotifChannel_ChannelFrontend:
		case basetypes.NotifChannel_ChannelEmail:
		case basetypes.NotifChannel_ChannelSMS:
		default:
			return fmt.Errorf("invalid channel")
		}
		h.Channel = _channel
		return nil
	}
}

func WithExtra(extra *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if extra == nil {
			if must {
				return fmt.Errorf("invalid extra")
			}
			return nil
		}
		if *extra == "" {
			return fmt.Errorf("invalid extra")
		}
		h.Extra = extra
		return nil
	}
}

//nolint:gocyclo
func WithEventType(eventtype *basetypes.UsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if eventtype == nil {
			if must {
				return fmt.Errorf("invalid eventtype")
			}
			return nil
		}
		switch *eventtype {
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
			return fmt.Errorf("invalid eventtype")
		}
		h.EventType = eventtype
		return nil
	}
}

func WithNotifType(_type *basetypes.NotifType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid notiftype")
			}
			return nil
		}
		switch *_type {
		case basetypes.NotifType_NotifMulticast:
		case basetypes.NotifType_NotifUnicast:
		default:
			return fmt.Errorf("invalid type")
		}
		h.NotifType = _type
		return nil
	}
}

func WithVars(vars *templatemwpb.TemplateVars, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if vars == nil {
			if must {
				return fmt.Errorf("invalid vars")
			}
			return nil
		}
		h.Vars = vars
		return nil
	}
}

func WithEntIDs(ids *[]string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if ids == nil {
			if must {
				return fmt.Errorf("invalid entids")
			}
			return nil
		}
		if len(*ids) == 0 {
			return fmt.Errorf("invalid entids")
		}
		_reqs := []*notifcrud.Req{}
		for _, id := range *ids {
			_req := &notifcrud.Req{}
			if id != "" {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				_req.EntID = &_id
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

//nolint:funlen,gocyclo
func WithReqs(reqs []*npool.NotifReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*notifcrud.Req{}
		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.LangID == nil {
					return fmt.Errorf("invalid langid")
				}
				if req.EventID == nil {
					return fmt.Errorf("invalid eventid")
				}
				if req.EventType == nil {
					return fmt.Errorf("invalid eventtype")
				}
				if req.Title == nil {
					return fmt.Errorf("invalid title")
				}
				if req.Content == nil {
					return fmt.Errorf("invalid content")
				}
				if req.Channel == nil {
					return fmt.Errorf("invalid channel")
				}
				if req.NotifType == nil {
					return fmt.Errorf("invalid notiftype")
				}
			}
			_req := &notifcrud.Req{}
			if req.ID != nil {
				id := req.GetID()
				_req.ID = &id
			}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.AppID != nil {
				id, err := uuid.Parse(req.GetAppID())
				if err != nil {
					return err
				}
				_req.AppID = &id
			}
			if req.UserID != nil {
				id, err := uuid.Parse(req.GetUserID())
				if err != nil {
					return err
				}
				_req.UserID = &id
			}
			if req.LangID != nil {
				id, err := uuid.Parse(req.GetLangID())
				if err != nil {
					return err
				}
				_req.LangID = &id
			}
			if req.EventID != nil {
				id, err := uuid.Parse(req.GetEventID())
				if err != nil {
					return err
				}
				_req.EventID = &id
			}
			if req.NotifType != nil {
				switch req.GetNotifType() {
				case basetypes.NotifType_NotifMulticast:
				case basetypes.NotifType_NotifUnicast:
				default:
					return fmt.Errorf("invalid Type")
				}
				_req.NotifType = req.NotifType
			}
			if req.EventType != nil {
				switch req.GetEventType() {
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
					return fmt.Errorf("invalid EventType")
				}
				_req.EventType = req.EventType
			}
			if req.Channel != nil {
				switch req.GetChannel() {
				case basetypes.NotifChannel_ChannelEmail:
				case basetypes.NotifChannel_ChannelSMS:
				case basetypes.NotifChannel_ChannelFrontend:
				default:
					return fmt.Errorf("invalid Channel")
				}
				_req.Channel = req.Channel
			}
			if req.UseTemplate != nil {
				_req.UseTemplate = req.UseTemplate
			}
			if req.Title != nil {
				_req.Title = req.Title
			}
			if req.Content != nil {
				_req.Content = req.Content
			}
			if req.Extra != nil {
				_req.Extra = req.Extra
			}
			if req.Notified != nil {
				_req.Notified = req.Notified
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

//nolint:gocyclo
func WithMultiNotifReqs(reqs []*npool.XNotifReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			_req := &MultiNotifReq{
				Vars:  req.Vars,
				Extra: req.Extra,
			}
			if req.UserID != nil {
				id, err := uuid.Parse(req.GetUserID())
				if err != nil {
					return err
				}
				_req.UserID = &id
			}
			switch req.GetEventType() {
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
				return fmt.Errorf("invalid EventType")
			}
			_req.EventType = req.EventType
			switch req.GetNotifType() {
			case basetypes.NotifType_NotifMulticast:
			case basetypes.NotifType_NotifUnicast:
			default:
				return fmt.Errorf("invalid NotifType")
			}
			_req.NotifType = req.NotifType
			h.MultiNotifReqs = append(h.MultiNotifReqs, _req)
		}
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &notifcrud.Conds{}
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
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.LangID != nil {
			id, err := uuid.Parse(conds.GetLangID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.LangID = &cruder.Cond{Op: conds.GetLangID().GetOp(), Val: id}
		}
		if conds.EventID != nil {
			id, err := uuid.Parse(conds.GetEventID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EventID = &cruder.Cond{Op: conds.GetEventID().GetOp(), Val: id}
		}
		if conds.NotifType != nil {
			switch conds.GetNotifType().GetValue() {
			case uint32(basetypes.NotifType_NotifMulticast):
			case uint32(basetypes.NotifType_NotifUnicast):
			default:
				return fmt.Errorf("invalid Type")
			}
			_type := conds.GetNotifType().GetValue()
			h.Conds.Type = &cruder.Cond{Op: conds.GetNotifType().GetOp(), Val: basetypes.NotifType(_type)}
		}
		if conds.EventType != nil {
			switch conds.GetEventType().GetValue() {
			case uint32(basetypes.UsedFor_WithdrawalRequest):
			case uint32(basetypes.UsedFor_WithdrawalCompleted):
			case uint32(basetypes.UsedFor_DepositReceived):
			case uint32(basetypes.UsedFor_KYCApproved):
			case uint32(basetypes.UsedFor_KYCRejected):
			case uint32(basetypes.UsedFor_Announcement):
			case uint32(basetypes.UsedFor_GoodBenefit1):
			case uint32(basetypes.UsedFor_UpdateEmail):
			case uint32(basetypes.UsedFor_UpdateMobile):
			case uint32(basetypes.UsedFor_UpdatePassword):
			case uint32(basetypes.UsedFor_UpdateGoogleAuth):
			case uint32(basetypes.UsedFor_NewLogin):
			case uint32(basetypes.UsedFor_OrderCompleted):
			case uint32(basetypes.UsedFor_OrderChildsRenewNotify):
			case uint32(basetypes.UsedFor_OrderChildsRenew):
			case uint32(basetypes.UsedFor_WithdrawReviewNotify):
			default:
				return fmt.Errorf("invalid EventType")
			}
			_type := conds.GetEventType().GetValue()
			h.Conds.EventType = &cruder.Cond{Op: conds.GetEventType().GetOp(), Val: basetypes.UsedFor(_type)}
		}
		if conds.Channel != nil {
			switch conds.GetChannel().GetValue() {
			case uint32(basetypes.NotifChannel_ChannelFrontend):
			case uint32(basetypes.NotifChannel_ChannelEmail):
			case uint32(basetypes.NotifChannel_ChannelSMS):
			default:
				return fmt.Errorf("invalid channel")
			}
			channel := conds.GetChannel().GetValue()
			h.Conds.Channel = &cruder.Cond{Op: conds.GetChannel().GetOp(), Val: basetypes.NotifChannel(channel)}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{Op: conds.GetEntIDs().GetOp(), Val: ids}
		}
		if conds.IDs != nil {
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: conds.GetIDs().GetValue()}
		}
		if conds.EventIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEventIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EventIDs = &cruder.Cond{Op: conds.GetEventIDs().GetOp(), Val: ids}
		}
		if conds.Notified != nil {
			h.Conds.Notified = &cruder.Cond{Op: conds.GetNotified().GetOp(), Val: conds.GetNotified().GetValue()}
		}
		if conds.Extra != nil {
			h.Conds.Extra = &cruder.Cond{Op: conds.GetExtra().GetOp(), Val: conds.GetExtra().GetValue()}
		}
		if conds.UseTemplate != nil {
			h.Conds.UseTemplate = &cruder.Cond{Op: conds.GetUseTemplate().GetOp(), Val: conds.GetUseTemplate().GetValue()}
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
