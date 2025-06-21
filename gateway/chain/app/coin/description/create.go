package description

import (
	"context"

	descmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin/description"
	descmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin/description"
)

func (h *Handler) CreateCoinDescription(ctx context.Context) (*descmwpb.CoinDescription, error) {
	handler, err := descmw.NewHandler(
		ctx,
		descmw.WithAppID(h.AppID, true),
		descmw.WithCoinTypeID(h.CoinTypeID, true),
		descmw.WithUsedFor(h.UsedFor, true),
		descmw.WithTitle(h.Title, true),
		descmw.WithMessage(h.Message, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.CreateCoinDescription(ctx)
}
