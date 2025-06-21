package description

import (
	"context"

	descmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin/description"
	descmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin/description"
)

func (h *Handler) UpdateCoinDescription(ctx context.Context) (*descmwpb.CoinDescription, error) {
	handler, err := descmw.NewHandler(
		ctx,
		descmw.WithID(h.ID, true),
		descmw.WithTitle(h.Title, true),
		descmw.WithMessage(h.Message, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.UpdateCoinDescription(ctx)
}
