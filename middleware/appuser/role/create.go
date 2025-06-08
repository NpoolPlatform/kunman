package role

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/role"
	rolecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/role"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateRole(ctx context.Context) (*npool.Role, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	// TODO: deduplicate

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := rolecrud.SetQueryConds(cli.AppRole.Query(), &rolecrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			Role:  &cruder.Cond{Op: cruder.EQ, Val: *h.Role},
		})
		if err != nil {
			return err
		}

		exist, err := stm.Exist(_ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("role exist")
		}

		if h.Default != nil && *h.Default {
			stm, err := rolecrud.SetQueryConds(cli.AppRole.Query(), &rolecrud.Conds{
				Default: &cruder.Cond{Op: cruder.EQ, Val: *h.Default},
				AppID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				Role:    &cruder.Cond{Op: cruder.NEQ, Val: *h.Role},
			})
			if err != nil {
				return err
			}

			exist, err := stm.Exist(_ctx)
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("default role exist")
			}
		}

		if _, err := rolecrud.CreateSet(
			cli.AppRole.Create(),
			&rolecrud.Req{
				EntID:       h.EntID,
				AppID:       h.AppID,
				CreatedBy:   h.CreatedBy,
				Role:        h.Role,
				Description: h.Description,
				Default:     h.Default,
				Genesis:     h.Genesis,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRole(ctx)
}

func (h *Handler) CreateRoles(ctx context.Context) ([]*npool.Role, error) {
	ids := []uuid.UUID{}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.EntID == nil {
				req.EntID = &id
			}

			stm, err := rolecrud.SetQueryConds(cli.AppRole.Query(), &rolecrud.Conds{
				AppID: &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
				Role:  &cruder.Cond{Op: cruder.EQ, Val: *req.Role},
			})
			if err != nil {
				return err
			}

			exist, err := stm.Exist(_ctx)
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("role exist")
			}

			if _, err := rolecrud.CreateSet(
				cli.AppRole.Create(),
				req,
			).Save(ctx); err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &rolecrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Limit = int32(len(ids))
	infos, _, err := h.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
