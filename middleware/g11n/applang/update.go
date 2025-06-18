package applang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/applang"
	applangcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/applang"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) UpdateLang(ctx context.Context) (*npool.Lang, error) {
	info, err := h.GetLang(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("applang not exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if h.Main != nil {
			if *h.Main {
				id := uuid.MustParse(info.AppID)
				h.AppID = &id
				h.Conds = &applangcrud.Conds{
					ID:    &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
					AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
					Main:  &cruder.Cond{Op: cruder.EQ, Val: true},
				}
				exist, err := h.ExistAppLangConds(ctx)
				if err != nil {
					return err
				}
				if exist {
					return fmt.Errorf("applang main exist")
				}
			}
		}
		if _, err := applangcrud.UpdateSet(
			cli.AppLang.UpdateOneID(*h.ID),
			&applangcrud.Req{
				Main: h.Main,
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
