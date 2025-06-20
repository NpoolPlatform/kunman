package readstate

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/readstate"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement/readstate"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateReadState(ctx context.Context) (info *npool.ReadState, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	h.Conds = &crud.Conds{
		AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:         &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		AnnouncementID: &cruder.Cond{Op: cruder.EQ, Val: *h.AnnouncementID},
	}

	exist, err := h.ExistReadStateConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("read state exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := crud.CreateSet(
			cli.ReadAnnouncement.Create(),
			&crud.Req{
				EntID:          h.EntID,
				AppID:          h.AppID,
				UserID:         h.UserID,
				AnnouncementID: h.AnnouncementID,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetReadState(ctx)
}
