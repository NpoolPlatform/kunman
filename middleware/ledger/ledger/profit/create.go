package profit

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/profit"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/profit"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createProfit(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.Profit.Create(),
			&h.Req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *Handler) CreateProfit(ctx context.Context) (*npool.Profit, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.createProfit(ctx); err != nil {
		return nil, err
	}

	return h.GetProfit(ctx)
}
