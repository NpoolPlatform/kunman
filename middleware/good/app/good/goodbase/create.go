package goodbase

import (
	"context"

	extrainfocrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/extrainfo"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createGoodExtraInfo(ctx context.Context, cli *ent.Client) error {
	if _, err := extrainfocrud.CreateSet(
		cli.ExtraInfo.Create(),
		&extrainfocrud.Req{
			AppGoodID: h.EntID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createGoodBase(ctx context.Context, cli *ent.Client) error {
	if _, err := appgoodbasecrud.CreateSet(
		cli.AppGoodBase.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateGoodBase(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createGoodExtraInfo(_ctx, cli); err != nil {
			return err
		}
		return handler.createGoodBase(_ctx, cli)
	})
}
