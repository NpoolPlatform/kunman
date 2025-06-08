package appfee

import (
	"context"

	appfeemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"

	"github.com/google/uuid"
)

func (h *Handler) CreateAppFee(ctx context.Context) (*npool.AppFee, error) {
	if h.AppGoodID == nil {
		h.AppGoodID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := appfeemwcli.CreateFee(ctx, &appfeemwpb.FeeReq{
		AppID:                   h.AppID,
		GoodID:                  h.GoodID,
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
