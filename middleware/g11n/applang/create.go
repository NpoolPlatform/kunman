package applang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/applang"
	applangcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/applang"
	langcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/lang"
	langmw "github.com/NpoolPlatform/kunman/middleware/g11n/lang"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkReqMainCount() error {
	mainMap := map[uuid.UUID]*uuid.UUID{}
	for _, req := range h.Reqs {
		if req.Main != nil && *req.Main {
			main := mainMap[*req.AppID]
			if main != nil {
				return fmt.Errorf("too many applang main")
			}
			mainMap[*req.AppID] = req.AppID
		}
	}
	return nil
}

func (h *createHandler) checkRepeat() error {
	countryMap := map[string]*uuid.UUID{}
	for _, req := range h.Reqs {
		_, ok := countryMap[req.AppID.String()+req.LangID.String()]
		if ok {
			return fmt.Errorf("duplicate langid")
		}
		countryMap[req.AppID.String()+req.LangID.String()] = req.LangID
	}
	return nil
}

func (h *createHandler) createLang(ctx context.Context, tx *ent.Tx, req *applangcrud.Req) error {
	handler, err := langmw.NewHandler(
		ctx,
	)
	if err != nil {
		return err
	}
	handler.Conds = &langcrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
	}
	existLang, err := handler.ExistLangCondsWithClient(ctx, tx.Client())
	if err != nil {
		return err
	}
	if !existLang {
		return fmt.Errorf("lang not exist")
	}

	h.Conds = &applangcrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		LangID: &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
	}
	exist, err := h.ExistAppLangCondsWithClient(ctx, tx.Client())
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("applang exist")
	}
	if req.Main != nil {
		if *req.Main {
			h.Conds = &applangcrud.Conds{
				AppID: &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
				Main:  &cruder.Cond{Op: cruder.EQ, Val: true},
			}
			exist, err := h.ExistAppLangCondsWithClient(ctx, tx.Client())
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("applang main exist")
			}
		}
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := applangcrud.CreateSet(
		tx.AppLang.Create(),
		&applangcrud.Req{
			EntID:  req.EntID,
			AppID:  req.AppID,
			LangID: req.LangID,
			Main:   req.Main,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil
}

func (h *Handler) CreateLang(ctx context.Context) (*npool.Lang, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		req := &applangcrud.Req{
			EntID:  h.EntID,
			AppID:  h.AppID,
			LangID: h.LangID,
			Main:   h.Main,
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
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.checkReqMainCount(); err != nil {
			return err
		}
		if err := handler.checkRepeat(); err != nil {
			return err
		}
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

	h.Conds = &applangcrud.Conds{
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
