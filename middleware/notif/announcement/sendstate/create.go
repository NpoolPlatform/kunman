package sendstate

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/sendstate"
	amt "github.com/NpoolPlatform/kunman/middleware/notif/announcement"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement/sendstate"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *Handler) CreateSendStateWithClient(ctx context.Context, cli *ent.Client) (info *npool.SendState, err error) {
	handler := &createHandler{
		Handler: h,
	}

	// get announcement first to get channel attr
	amtID := handler.AnnouncementID.String()
	amtHandler, err := amt.NewHandler(ctx, amt.WithEntID(&amtID, true))
	if err != nil {
		return nil, err
	}

	announcement, err := amtHandler.GetAnnouncementWithClient(ctx, cli)
	if err != nil {
		return nil, err
	}
	if announcement == nil {
		return nil, fmt.Errorf("invalid announcement id")
	}

	h.Conds = &crud.Conds{
		AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:         &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		AnnouncementID: &cruder.Cond{Op: cruder.EQ, Val: *h.AnnouncementID},
	}

	exist, err := h.ExistSendStateCondsWithClient(ctx, cli)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("send state exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	_info, err := crud.CreateSet(
		cli.SendAnnouncement.Create(),
		&crud.Req{
			EntID:          h.EntID,
			AppID:          h.AppID,
			UserID:         h.UserID,
			AnnouncementID: h.AnnouncementID,
			Channel:        &announcement.Channel,
		},
	).Save(ctx)
	if err != nil {
		return nil, err
	}

	h.ID = &_info.ID

	return h.GetSendStateWithClient(ctx, cli)
}

func (h *Handler) CreateSendState(ctx context.Context) (info *npool.SendState, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = h.CreateSendStateWithClient(_ctx, cli)
		return err
	})
	return
}

func (h *Handler) CreateSendStates(ctx context.Context) (infos []*npool.SendState, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			h.AppID = req.AppID
			h.UserID = req.UserID
			h.AnnouncementID = req.AnnouncementID
			h.Channel = req.Channel

			h.Conds = &crud.Conds{
				AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				UserID:         &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
				AnnouncementID: &cruder.Cond{Op: cruder.EQ, Val: *h.AnnouncementID},
			}

			exist, err := h.ExistSendStateCondsWithClient(ctx, cli)
			if err != nil {
				return err
			}
			if exist {
				continue
			}

			info, err := h.CreateSendStateWithClient(ctx, cli)
			if err != nil {
				return err
			}
			infos = append(infos, info)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}
