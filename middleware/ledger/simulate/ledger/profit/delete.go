package profit

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/simulate/ledger/profit"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entprofit "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/simulateprofit"
)

func (h *Handler) DeleteProfit(ctx context.Context) (*npool.Profit, error) {
	info, err := h.GetProfit(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid id %v", *h.ID)
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := cli.SimulateProfit.
			Update().
			Where(
				entprofit.ID(*h.ID),
			).
			SetDeletedAt(now).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
