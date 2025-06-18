package user

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/user"
	usercrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/user"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createNotifUser(ctx context.Context, tx *ent.Tx, req *usercrud.Req) error {
	if req.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if req.UserID == nil {
		return fmt.Errorf("invalid langid")
	}
	if req.EventType == nil {
		return fmt.Errorf("invalid eventtype")
	}

	h.Conds = &usercrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:    &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: *req.EventType},
	}
	exist, err := h.ExistNotifUserCondsWithClient(ctx, tx.Client())
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("user notif exist")
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := usercrud.CreateSet(
		tx.NotifUser.Create(),
		&usercrud.Req{
			EntID:     req.EntID,
			AppID:     req.AppID,
			UserID:    req.UserID,
			EventType: req.EventType,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *Handler) CreateNotifUser(ctx context.Context) (*npool.NotifUser, error) {
	handler := &createHandler{
		Handler: h,
	}
	req := &usercrud.Req{
		EntID:     handler.EntID,
		AppID:     handler.AppID,
		UserID:    handler.UserID,
		EventType: handler.EventType,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createNotifUser(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetNotifUser(ctx)
}
