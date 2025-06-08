package appfee

import (
	"context"

	appfeemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateAppFee(ctx context.Context) (*npool.AppFee, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkAppFee(ctx); err != nil {
		return nil, err
	}

	if err := appfeemwcli.UpdateFee(ctx, &appfeemwpb.FeeReq{
		ID:                      h.ID,
		EntID:                   h.EntID,
		AppGoodID:               h.AppGoodID,
		ProductPage:             h.ProductPage,
		Name:                    h.Name,
		Banner:                  h.Banner,
		UnitValue:               h.UnitValue,
		MinOrderDurationSeconds: h.MinOrderDurationSeconds,
		CancelMode:              h.CancelMode,
	}); err != nil {
		return nil, err
	}
	return h.GetAppFee(ctx)
}
