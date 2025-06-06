package api

import (
	"context"

	crud "github.com/NpoolPlatform/basal-middleware/pkg/crud/api"
	"github.com/NpoolPlatform/basal-middleware/pkg/db"
	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent"
	entapi "github.com/NpoolPlatform/basal-middleware/pkg/db/ent/api"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	ids []uuid.UUID
}

func (h *createHandler) createAPI(ctx context.Context, tx *ent.Tx, req *crud.Req) error {
	_api, err := tx.
		API.
		Query().
		Where(
			entapi.Protocol(req.Protocol.String()),
			entapi.ServiceName(*req.ServiceName),
			entapi.Method(req.Method.String()),
			entapi.Path(*req.Path),
			entapi.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if _api != nil {
		h.ids = append(h.ids, _api.EntID)
		return nil
	}
	_api, err = crud.CreateSet(tx.API.Create(), req).Save(ctx)
	if err != nil {
		return err
	}
	h.ids = append(h.ids, _api.EntID)
	return nil
}

func (h *Handler) CreateAPIs(ctx context.Context) ([]*npool.API, error) {
	handler := &createHandler{
		Handler: h,
		ids:     []uuid.UUID{},
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.createAPI(ctx, tx, req); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	h.Conds = &crud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: handler.ids},
	}
	h.Offset = 0
	h.Limit = int32(len(handler.ids))
	infos, _, err := h.GetAPIs(ctx)
	return infos, err
}

func (h *Handler) CreateAPI(ctx context.Context) (*npool.API, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createAPI(ctx, tx, &h.Req)
	})
	if err != nil {
		return nil, err
	}

	return h.GetAPI(ctx)
}
