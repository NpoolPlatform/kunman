package sendstate

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/sendstate"
	announcement1 "github.com/NpoolPlatform/kunman/middleware/notif/announcement"
	"github.com/NpoolPlatform/kunman/middleware/notif/announcement/handler"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement/sendstate"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Handler struct {
	*handler.Handler
	Channel *basetypes.NotifChannel
	Reqs    []*crud.Req
	Conds   *crud.Conds
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	_handler, err := handler.NewHandler(ctx, options...)
	if err != nil {
		return nil, err
	}
	h := &Handler{
		Handler: _handler,
	}

	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, h); err != nil {
			return nil, err
		}
	}
	return h, nil
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
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op: conds.GetAppID().GetOp(), Val: id,
			}
		}
		if conds.UserID != nil {
			userID, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op: conds.GetUserID().GetOp(), Val: userID,
			}
		}

		if conds.AnnouncementID != nil {
			amtID, err := uuid.Parse(conds.GetAnnouncementID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AnnouncementID = &cruder.Cond{
				Op: conds.GetAnnouncementID().GetOp(), Val: amtID,
			}
		}
		if conds.Channel != nil {
			channel := conds.GetChannel().GetValue()
			h.Conds.Channel = &cruder.Cond{
				Op: conds.GetChannel().GetOp(), Val: basetypes.NotifChannel(channel),
			}
		}
		if conds.UserIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetUserIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.UserIDs = &cruder.Cond{
				Op: conds.GetUserIDs().GetOp(), Val: ids,
			}
		}
		return nil
	}
}

func WithReqs(reqs []*npool.SendStateReq, must bool) func(context.Context, *Handler) error {
	return func(_ctx context.Context, h *Handler) error {
		if len(reqs) == 0 {
			return fmt.Errorf("invalid reqs")
		}
		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.UserID == nil {
					return fmt.Errorf("invalid userid")
				}
				if req.AnnouncementID == nil {
					return fmt.Errorf("invalid announcementid")
				}
			}
			if req.AppID == nil || req.UserID == nil || req.AnnouncementID == nil {
				continue
			}

			// AppID
			_req := &crud.Req{}
			appID, err := uuid.Parse(*req.AppID)
			if err != nil {
				return err
			}
			_req.AppID = &appID

			// UserID
			userID, err := uuid.Parse(*req.UserID)
			if err != nil {
				return err
			}
			_req.UserID = &userID

			// AnnouncementID
			amtID, err := uuid.Parse(*req.AnnouncementID)
			if err != nil {
				return err
			}
			amtHandler, err := announcement1.NewHandler(_ctx, announcement1.WithEntID(req.AnnouncementID, true))
			if err != nil {
				return err
			}
			_req.AnnouncementID = &amtID

			// Channel
			info, err := amtHandler.GetAnnouncement(_ctx)
			if err != nil {
				return err
			}
			_req.Channel = &info.Channel

			h.Reqs = append(h.Reqs, _req)
		}
		return nil
	}
}
