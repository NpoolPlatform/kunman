package unsold

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/unsold"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entunsoldstatement "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/unsoldstatement"
)

func (h *Handler) DeleteUnsoldStatement(ctx context.Context) (*npool.UnsoldStatement, error) {
	info, err := h.GetUnsoldStatement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("id not exist %v", *h.ID)
	}
	if h.ID == nil {
		h.ID = &info.ID
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := cli.UnsoldStatement.
			Update().
			Where(
				entunsoldstatement.ID(*h.ID),
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
