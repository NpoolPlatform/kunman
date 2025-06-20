package announcement

import (
	"context"
	"fmt"
	"time"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Handler struct {
	ID      *uint32
	EntID   *uuid.UUID
	AppID   *uuid.UUID
	LangID  *uuid.UUID
	Title   *string
	Content *string
	Channel *basetypes.NotifChannel
	Type    *basetypes.NotifType
	EndAt   *uint32
	StartAt *uint32
	Conds   *crud.Conds
	Offset  int32
	Limit   int32
	UserID  *string
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

func WithLangID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid langid")
			}
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}

		h.LangID = &_id
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
		const leastTitleLen = 4
		if len(*title) < leastTitleLen {
			return fmt.Errorf("name %v too short", *title)
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
		const leastContentLen = 4
		if len(*content) < leastContentLen {
			return fmt.Errorf("content %v too short", *content)
		}
		h.Content = content
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

func WithAnnouncementType(_type *basetypes.NotifType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid announcementtype")
			}
			return nil
		}
		switch *_type {
		case basetypes.NotifType_NotifBroadcast:
		case basetypes.NotifType_NotifMulticast:
		default:
			return fmt.Errorf("type %v invalid", *_type)
		}
		h.Type = _type
		return nil
	}
}

func WithStartAt(startAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if startAt == nil {
			if must {
				return fmt.Errorf("invalid startat")
			}
			return nil
		}
		if *startAt < uint32(time.Now().Unix()) {
			return fmt.Errorf("invalid start at")
		}
		h.StartAt = startAt
		return nil
	}
}

func WithEndAt(endAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if endAt == nil {
			if must {
				return fmt.Errorf("invalid endat")
			}
			return nil
		}
		if *endAt < uint32(time.Now().Unix()) {
			return fmt.Errorf("invalid end at")
		}
		h.EndAt = endAt
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

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op: conds.GetEntID().GetOp(), Val: id,
			}
		}
		if conds.AppID != nil {
			appID, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op: conds.GetAppID().GetOp(), Val: appID,
			}
		}
		if conds.LangID != nil {
			langID, err := uuid.Parse(conds.GetLangID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.LangID = &cruder.Cond{
				Op: conds.GetLangID().GetOp(), Val: langID,
			}
		}
		if conds.StartAt != nil {
			h.Conds.StartAt = &cruder.Cond{
				Op: conds.GetStartAt().GetOp(), Val: conds.GetStartAt().GetValue(),
			}
		}
		if conds.EndAt != nil {
			h.Conds.EndAt = &cruder.Cond{
				Op: conds.GetEndAt().GetOp(), Val: conds.GetEndAt().GetValue(),
			}
		}
		if conds.Channel != nil {
			channel := conds.GetChannel().GetValue()
			h.Conds.Channel = &cruder.Cond{
				Op: conds.GetChannel().GetOp(), Val: basetypes.NotifChannel(channel),
			}
		}
		if conds.UserID != nil {
			userID, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			_userID := userID.String()
			h.UserID = &_userID
			h.Conds.UserID = &cruder.Cond{
				Op: conds.GetUserID().GetOp(), Val: userID,
			}
		}
		if conds.AnnouncementType != nil {
			switch conds.GetAnnouncementType().GetValue() {
			case uint32(basetypes.NotifType_NotifBroadcast):
			case uint32(basetypes.NotifType_NotifMulticast):
			default:
				return fmt.Errorf("invalid announcementtype")
			}
			_type := conds.GetAnnouncementType().GetValue()
			h.Conds.AnnouncementType = &cruder.Cond{
				Op: conds.GetAnnouncementType().GetOp(), Val: basetypes.NotifType(_type),
			}
		}
		return nil
	}
}
