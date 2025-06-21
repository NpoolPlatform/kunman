package appfiat

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/fiat"
	appfiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/fiat"
	fiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat"
	fiatmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createAppFiat(ctx context.Context, tx *ent.Tx) error {
	if _, err := appfiatcrud.CreateSet(
		tx.AppFiat.Create(),
		&appfiatcrud.Req{
			ID:           h.ID,
			EntID:        h.EntID,
			AppID:        h.AppID,
			FiatID:       h.FiatID,
			Name:         h.Name,
			DisplayNames: h.DisplayNames,
			Logo:         h.Logo,
			Display:      h.Display,
			DisplayIndex: h.DisplayIndex,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateFiat(ctx context.Context) (*npool.Fiat, error) {
	// TODO: deduplicate

	h.Conds = &appfiatcrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		FiatID: &cruder.Cond{Op: cruder.EQ, Val: *h.FiatID},
	}
	exist, err := h.ExistFiatConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("appfiat exist")
	}

	fiatHandler, err := fiatmw.NewHandler(
		ctx,
	)
	if err != nil {
		return nil, err
	}
	fiatHandler.Conds = &fiatcrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *h.FiatID},
	}
	fiat, err := fiatHandler.GetFiatOnly(ctx)
	if err != nil {
		return nil, err
	}
	if fiat == nil {
		return nil, fmt.Errorf("fiat not exist")
	}
	if h.Name == nil {
		h.Name = &fiat.Name
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppFiat(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFiat(ctx)
}
