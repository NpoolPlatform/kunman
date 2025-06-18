package goodbenefit

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/goodbenefit"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

func (h *Handler) UpdateGoodBenefit(ctx context.Context) (info *npool.GoodBenefit, err error) {
	info, err = h.GetGoodBenefit(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("good benefit not found")
	}
	if info.Generated && !*h.Generated {
		return nil, fmt.Errorf("can not be update")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.GoodBenefit.UpdateOneID(*h.ID),
			&crud.Req{
				Generated: h.Generated,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGoodBenefit(ctx)
}
