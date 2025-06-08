package description

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin/description"
	descriptioncrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin/description"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
)

func (h *Handler) UpdateCoinDescription(ctx context.Context) (*npool.CoinDescription, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := descriptioncrud.UpdateSet(
			cli.CoinDescription.UpdateOneID(*h.ID),
			&descriptioncrud.Req{
				Title:   h.Title,
				Message: h.Message,
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
