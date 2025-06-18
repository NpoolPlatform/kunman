package message

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/message"
	messagecrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/message"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
	entmessage "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/message"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

func (h *Handler) UpdateMessage(ctx context.Context) (*npool.Message, error) {
	info, err := h.GetMessage(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("message not exist")
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if h.MessageID != nil {
			appid := uuid.MustParse(info.AppID)
			h.AppID = &appid
			langid := uuid.MustParse(info.LangID)
			h.LangID = &langid
			h.Conds = &messagecrud.Conds{
				AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				LangID:    &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
				ID:        &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
				MessageID: &cruder.Cond{Op: cruder.EQ, Val: *h.MessageID},
			}
			exist, err := h.ExistMessageCondsWithClient(ctx, tx.Client())
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("messageid exist")
			}
		}

		if _, err := tx.
			Message.
			Query().
			Where(
				entmessage.ID(*h.ID),
				entmessage.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx); err != nil {
			return err
		}

		if _, err := messagecrud.UpdateSet(
			tx.Message.UpdateOneID(*h.ID),
			&messagecrud.Req{
				MessageID: h.MessageID,
				Message:   h.Message,
				GetIndex:  h.GetIndex,
				Disabled:  h.Disabled,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetMessage(ctx)
}
