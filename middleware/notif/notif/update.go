//nolint:nolintlint,dupl
package notif

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	notifcrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) UpdateNotif(ctx context.Context) (*npool.Notif, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.Notified == nil {
		return nil, fmt.Errorf("invalid notified")
	}

	info, err := h.GetNotif(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("notif not exist")
	}
	if info.Notified && *h.Notified != info.Notified {
		return nil, fmt.Errorf("invalid notified")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := notifcrud.UpdateSet(
			tx.Notif.UpdateOneID(*h.ID),
			&notifcrud.Req{
				Notified:    h.Notified,
				UseTemplate: h.UseTemplate,
				Title:       h.Title,
				Content:     h.Content,
				Channel:     h.Channel,
				Extra:       h.Extra,
				NotifType:   h.NotifType,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetNotif(ctx)
}

// nolint:gocyclo
func (h *Handler) UpdateNotifs(ctx context.Context) ([]*npool.Notif, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.ID == nil {
				return fmt.Errorf("invalid id")
			}
			if req.Notified == nil {
				return fmt.Errorf("invalid notified")
			}

			h.ID = req.ID
			info, err := h.GetNotif(ctx)
			if err != nil {
				return err
			}
			if info == nil {
				return fmt.Errorf("notif not exist")
			}
			if info.Notified {
				if *req.Notified != info.Notified {
					return fmt.Errorf("invalid notified")
				}
			}
			if _, err := notifcrud.UpdateSet(
				tx.Notif.UpdateOneID(*req.ID),
				&notifcrud.Req{
					Notified: req.Notified,
				},
			).Save(ctx); err != nil {
				return err
			}
			entID := uuid.MustParse(info.EntID)
			ids = append(ids, entID)
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
