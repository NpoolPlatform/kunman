package exchange

import (
	"context"

	submwpb "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/credit/exchange"

	"github.com/google/uuid"
)

func (h *Handler) CreateExchange(ctx context.Context) (*submwpb.Exchange, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := submwcli.CreateExchange(ctx, &submwpb.ExchangeReq{
		EntID:             h.EntID,
		AppID:             h.AppID,
		UsageType:         h.UsageType,
		Credit:            h.Credit,
		ExchangeThreshold: h.ExchangeThreshold,
		Path:              h.Path,
	}); err != nil {
		return nil, err
	}
	return h.GetExchange(ctx)
}
