package sms

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/sms"
	smstemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/sms"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkRepeat() error {
	countryMap := map[string]*uuid.UUID{}
	for _, req := range h.Reqs {
		_, ok := countryMap[req.AppID.String()+req.LangID.String()+req.UsedFor.String()]
		if ok {
			return fmt.Errorf("duplicate smstemplate")
		}
		countryMap[req.AppID.String()+req.LangID.String()+req.UsedFor.String()] = req.LangID
	}
	return nil
}

func (h *createHandler) createSMSTemplate(ctx context.Context, tx *ent.Tx, req *smstemplatecrud.Req) error {
	if req.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if req.LangID == nil {
		return fmt.Errorf("invalid langid")
	}
	if req.UsedFor == nil {
		return fmt.Errorf("invalid usedfor")
	}

	h.Conds = &smstemplatecrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		LangID:  &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
		UsedFor: &cruder.Cond{Op: cruder.EQ, Val: *req.UsedFor},
	}
	exist, err := h.ExistSMSTemplateCondsWithClient(ctx, tx.Client())
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("smstemplate exist")
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := smstemplatecrud.CreateSet(
		tx.SMSTemplate.Create(),
		&smstemplatecrud.Req{
			EntID:   req.EntID,
			AppID:   req.AppID,
			LangID:  req.LangID,
			UsedFor: req.UsedFor,
			Subject: req.Subject,
			Message: req.Message,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil
}

func (h *Handler) CreateSMSTemplate(ctx context.Context) (*npool.SMSTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}
	req := &smstemplatecrud.Req{
		EntID:   handler.EntID,
		AppID:   handler.AppID,
		LangID:  handler.LangID,
		UsedFor: handler.UsedFor,
		Subject: handler.Subject,
		Message: handler.Message,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createSMSTemplate(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSMSTemplate(ctx)
}

func (h *Handler) CreateSMSTemplates(ctx context.Context) ([]*npool.SMSTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.checkRepeat(); err != nil {
			return err
		}
		for _, req := range h.Reqs {
			if err := handler.createSMSTemplate(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &smstemplatecrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetSMSTemplates(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
