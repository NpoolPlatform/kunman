//nolint:nolintlint,dupl
package notif

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	notifcrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createNotif(ctx context.Context, tx *ent.Tx, req *notifcrud.Req) error {
	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := notifcrud.CreateSet(
		tx.Notif.Create(),
		&notifcrud.Req{
			EntID:       req.EntID,
			AppID:       req.AppID,
			LangID:      req.LangID,
			UserID:      req.UserID,
			EventID:     req.EventID,
			Notified:    req.Notified,
			EventType:   req.EventType,
			UseTemplate: req.UseTemplate,
			Title:       req.Title,
			Content:     req.Content,
			Channel:     req.Channel,
			Extra:       req.Extra,
			NotifType:   req.NotifType,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID
	return nil
}

func (h *Handler) CreateNotif(ctx context.Context) (*npool.Notif, error) {
	handler := &createHandler{
		Handler: h,
	}
	req := &notifcrud.Req{
		EntID:       handler.EntID,
		AppID:       handler.AppID,
		LangID:      handler.LangID,
		UserID:      handler.UserID,
		EventID:     handler.EventID,
		Notified:    handler.Notified,
		EventType:   handler.EventType,
		UseTemplate: handler.UseTemplate,
		Title:       handler.Title,
		Content:     handler.Content,
		Channel:     handler.Channel,
		Extra:       handler.Extra,
		NotifType:   handler.NotifType,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createNotif(_ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetNotif(ctx)
}

func (h *Handler) CreateNotifs(ctx context.Context) ([]*npool.Notif, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.EntID != nil {
				handler.EntID = req.EntID
				exist, err := handler.ExistNotif(ctx)
				if err != nil {
					return err
				}
				if exist {
					return fmt.Errorf("notif id is exist")
				}
			}
			if err := handler.createNotif(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &notifcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetNotifs(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
