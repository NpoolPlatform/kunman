package message

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/message"
	applangmw "github.com/NpoolPlatform/kunman/middleware/g11n/applang"
	applangcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/applang"
	messagecrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/message"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkAppLang(ctx context.Context, tx *ent.Tx) error {
	for _, req := range h.Reqs {
		handler, err := applangmw.NewHandler(
			ctx,
		)
		if err != nil {
			return err
		}
		handler.Conds = &applangcrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			LangID: &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
		}
		exist, err := handler.ExistAppLangCondsWithClient(ctx, tx.Client())
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid applang")
		}
	}
	return nil
}

func (h *createHandler) createMessage(ctx context.Context, tx *ent.Tx, req *messagecrud.Req) (*npool.Message, error) {
	h.Conds = &messagecrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		LangID:    &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
		MessageID: &cruder.Cond{Op: cruder.EQ, Val: *req.MessageID},
	}
	h.Limit = 2
	infos, _, err := h.GetMessagesWithClient(ctx, tx.Client())
	if err != nil {
		return nil, err
	}
	if infos != nil {
		return infos[0], nil
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := messagecrud.CreateSet(
		tx.Message.Create(),
		&messagecrud.Req{
			EntID:     req.EntID,
			AppID:     req.AppID,
			LangID:    req.LangID,
			MessageID: req.MessageID,
			Message:   req.Message,
			GetIndex:  req.GetIndex,
			Disabled:  req.Disabled,
		},
	).Save(ctx)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil, nil
}

func (h *Handler) CreateMessage(ctx context.Context) (*npool.Message, error) {
	handler := &createHandler{
		Handler: h,
	}
	h.Conds = &messagecrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		LangID:    &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
		MessageID: &cruder.Cond{Op: cruder.EQ, Val: *h.MessageID},
	}
	exist, err := h.ExistMessageConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("message exist")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		req := &messagecrud.Req{
			EntID:     h.EntID,
			AppID:     h.AppID,
			LangID:    h.LangID,
			MessageID: h.MessageID,
			Message:   h.Message,
			GetIndex:  h.GetIndex,
			Disabled:  h.Disabled,
		}
		if err := handler.checkAppLang(ctx, tx); err != nil {
			return err
		}
		info, err := handler.createMessage(ctx, tx, req)
		if err != nil {
			return err
		}
		if info != nil {
			id, err := uuid.Parse(info.GetEntID())
			if err != nil {
				return err
			}
			h.EntID = &id
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetMessage(ctx)
}

func (h *Handler) CreateMessages(ctx context.Context) ([]*npool.Message, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.checkAppLang(ctx, tx); err != nil {
			return err
		}
		for _, req := range h.Reqs {
			info, err := handler.createMessage(ctx, tx, req)
			if err != nil {
				return err
			}
			if info != nil {
				id, err := uuid.Parse(info.GetEntID())
				if err != nil {
					return err
				}
				h.EntID = &id
			}
			ids = append(ids, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &messagecrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetMessages(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
