package lang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/lang"
	langcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/lang"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	langIDs map[string]*uuid.UUID
}

func (h *createHandler) createLang(ctx context.Context, tx *ent.Tx, req *langcrud.Req) error {
	_id, ok := h.langIDs[*req.Lang]
	if ok {
		h.EntID = _id
		return nil
	}

	h.Conds = &langcrud.Conds{
		Lang: &cruder.Cond{Op: cruder.EQ, Val: *req.Lang},
	}
	h.Limit = 2
	infos, _, err := h.GetLangsWithClient(ctx, tx.Client())
	if err != nil {
		return err
	}
	if infos != nil {
		id := uuid.MustParse(infos[0].EntID)
		h.EntID = &id
		return nil
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := langcrud.CreateSet(
		tx.Lang.Create(),
		&langcrud.Req{
			EntID: req.EntID,
			Lang:  req.Lang,
			Logo:  req.Logo,
			Name:  req.Name,
			Short: req.Short,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID
	h.langIDs[*req.Lang] = h.EntID

	return nil
}

func (h *Handler) CreateLang(ctx context.Context) (*npool.Lang, error) {
	handler := &createHandler{
		Handler: h,
		langIDs: map[string]*uuid.UUID{},
	}
	h.Conds = &langcrud.Conds{
		Lang: &cruder.Cond{Op: cruder.EQ, Val: *h.Lang},
	}
	exist, err := h.ExistLangConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("lang exist")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		req := &langcrud.Req{
			EntID: h.EntID,
			Lang:  h.Lang,
			Logo:  h.Logo,
			Name:  h.Name,
			Short: h.Short,
		}
		if err := handler.createLang(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetLang(ctx)
}

func (h *Handler) CreateLangs(ctx context.Context) ([]*npool.Lang, error) {
	handler := &createHandler{
		Handler: h,
		langIDs: map[string]*uuid.UUID{},
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.createLang(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &langcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetLangs(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
