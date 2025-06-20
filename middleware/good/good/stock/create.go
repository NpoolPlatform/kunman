package stock

import (
	"context"

	stockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) _createStock(ctx context.Context, cli *ent.Client) error {
	h.SpotQuantity = h.Total
	if _, err := stockcrud.CreateSet(
		cli.Stock.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

// Only for test. Stock should always be created with good
func (h *Handler) CreateStock(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler._createStock(ctx, cli)
	})
}
