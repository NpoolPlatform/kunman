package description

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin/description"
	descriptioncrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin/description"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateCoinDescription(ctx context.Context) (*npool.CoinDescription, error) {
	// TODO: deduplicate

	h.Conds = &descriptioncrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		UsedFor:    &cruder.Cond{Op: cruder.EQ, Val: *h.UsedFor},
	}
	exist, err := h.ExistCoinDescriptionConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("description exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := descriptioncrud.CreateSet(
			cli.CoinDescription.Create(),
			&descriptioncrud.Req{
				ID:         h.ID,
				EntID:      h.EntID,
				AppID:      h.AppID,
				CoinTypeID: h.CoinTypeID,
				UsedFor:    h.UsedFor,
				Title:      h.Title,
				Message:    h.Message,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCoinDescription(ctx)
}
