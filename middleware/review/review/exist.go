package review

import (
	"context"

	reviewcrud "github.com/NpoolPlatform/kunman/middleware/review/crud/review"
	"github.com/NpoolPlatform/kunman/middleware/review/db"
	ent "github.com/NpoolPlatform/kunman/middleware/review/db/ent/generated"
)

func (h *Handler) ExistReviewConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := reviewcrud.SetQueryConds(cli.Review.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
