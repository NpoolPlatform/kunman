package default1

import (
	"context"

	defaultmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/default"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/default"
	defaultmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/default"

	"github.com/google/uuid"
)

func (h *Handler) CreateDefault(ctx context.Context) (*npool.Default, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := defaultmwcli.CreateDefault(ctx, &defaultmwpb.DefaultReq{
		EntID:      h.EntID,
		CoinTypeID: h.CoinTypeID,
		AppGoodID:  h.AppGoodID,
	}); err != nil {
		return nil, err
	}
	return h.GetDefault(ctx)
}
