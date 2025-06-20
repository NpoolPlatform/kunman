package frontend

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/frontend"
	frontendtemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/frontend"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

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
			return fmt.Errorf("duplicate frontendtemplate")
		}
		countryMap[req.AppID.String()+req.LangID.String()+req.UsedFor.String()] = req.LangID
	}
	return nil
}

func (h *createHandler) createFrontendTemplate(ctx context.Context, tx *ent.Tx, req *frontendtemplatecrud.Req) error {
	if req.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if req.LangID == nil {
		return fmt.Errorf("invalid langid")
	}
	if req.UsedFor == nil {
		return fmt.Errorf("invalid create usedfor")
	}

	h.Conds = &frontendtemplatecrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		LangID:  &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
		UsedFor: &cruder.Cond{Op: cruder.EQ, Val: *req.UsedFor},
	}
	exist, err := h.ExistFrontendTemplateCondsWithClient(ctx, tx.Client())
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("frontendtemplate exist")
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := frontendtemplatecrud.CreateSet(
		tx.FrontendTemplate.Create(),
		&frontendtemplatecrud.Req{
			EntID:   req.EntID,
			AppID:   req.AppID,
			LangID:  req.LangID,
			UsedFor: req.UsedFor,
			Title:   req.Title,
			Content: req.Content,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil
}

func (h *Handler) CreateFrontendTemplate(ctx context.Context) (*npool.FrontendTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}
	req := &frontendtemplatecrud.Req{
		EntID:   handler.EntID,
		AppID:   handler.AppID,
		LangID:  handler.LangID,
		UsedFor: handler.UsedFor,
		Title:   handler.Title,
		Content: handler.Content,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createFrontendTemplate(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFrontendTemplate(ctx)
}

func (h *Handler) CreateFrontendTemplates(ctx context.Context) ([]*npool.FrontendTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.checkRepeat(); err != nil {
			return err
		}
		for _, req := range h.Reqs {
			if err := handler.createFrontendTemplate(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &frontendtemplatecrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetFrontendTemplates(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
