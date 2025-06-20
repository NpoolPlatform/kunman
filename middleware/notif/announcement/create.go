package announcement

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if *h.StartAt > *h.EndAt {
		return fmt.Errorf("start at less than end at")
	}
	return nil
}

func (h *Handler) CreateAnnouncement(ctx context.Context) (info *npool.Announcement, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := crud.CreateSet(
			cli.Announcement.Create(),
			&crud.Req{
				EntID:   h.EntID,
				AppID:   h.AppID,
				Title:   h.Title,
				Content: h.Content,
				LangID:  h.LangID,
				Channel: h.Channel,
				Type:    h.Type,
				StartAt: h.StartAt,
				EndAt:   h.EndAt,
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

	return h.GetAnnouncement(ctx)
}
