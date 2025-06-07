package required

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	requiredcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/required"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

func (h *Handler) UpdateRequired(ctx context.Context) error {
	info, err := h.GetRequired(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid required")
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := requiredcrud.UpdateSet(
			cli.RequiredGood.UpdateOneID(*h.ID),
			&requiredcrud.Req{
				Must: h.Must,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
