package scope

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/app/scope"
	appgoodscopecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/app/scope"
)

func (h *Handler) DeleteAppGoodScope(ctx context.Context) (*npool.Scope, error) {
	info, err := h.GetAppGoodScope(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := appgoodscopecrud.UpdateSet(
			cli.AppGoodScope.UpdateOneID(*h.ID),
			&appgoodscopecrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return info, nil
}
