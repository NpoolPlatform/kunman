package lang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/lang"
	langcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/lang"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) UpdateLang(ctx context.Context) (*npool.Lang, error) {
	info, err := h.GetLang(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("lang not exist")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if h.Lang != nil {
			h.Conds = &langcrud.Conds{
				ID:   &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
				Lang: &cruder.Cond{Op: cruder.EQ, Val: *h.Lang},
			}
			exist, err := h.ExistLangCondsWithClient(ctx, tx.Client())
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("lang is exist")
			}
		}

		if _, err := langcrud.UpdateSet(
			tx.Lang.UpdateOneID(*h.ID),
			&langcrud.Req{
				Lang:  h.Lang,
				Logo:  h.Logo,
				Name:  h.Name,
				Short: h.Short,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetLang(ctx)
}
