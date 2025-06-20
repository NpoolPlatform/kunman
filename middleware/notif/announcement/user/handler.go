package user

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/user"
	announcement1 "github.com/NpoolPlatform/kunman/middleware/notif/announcement"
	"github.com/NpoolPlatform/kunman/middleware/notif/announcement/handler"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement/user"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Handler struct {
	*handler.Handler
	Conds *crud.Conds
	Reqs  []*crud.Req
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
		return nil
	}
}

func WithReqs(reqs []*npool.AnnouncementUserReq, must bool) func(context.Context, *Handler) error {
	return func(_ctx context.Context, h *Handler) error {
		_reqs := []*crud.Req{}
		for _, req := range _reqs {
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
			appID, err := uuid.Parse(req.AppID.String())
			if err != nil {
				return err
			}
			_req.AppID = &appID

			// UserID
			userID, err := uuid.Parse(req.UserID.String())
			if err != nil {
				return err
			}
			_req.UserID = &userID

			// AnnouncementID
			amtID, err := uuid.Parse(req.AnnouncementID.String())
			if err != nil {
				return err
			}
			_amtID := req.AnnouncementID.String()
			amtHandler, err := announcement1.NewHandler(_ctx, announcement1.WithEntID(&_amtID, true))
			if err != nil {
				return err
			}

			exist, err := amtHandler.ExistAnnouncement(_ctx)
			if err != nil {
				return err
			}
			if !exist {
				return fmt.Errorf("invalid user")
			}
			_req.AnnouncementID = &amtID

			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
