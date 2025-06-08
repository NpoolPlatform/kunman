package description

import (
	"context"

	appgooddescriptionmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/description"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/description"
	appgooddescriptionmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"

	"github.com/google/uuid"
)

func (h *Handler) CreateDescription(ctx context.Context) (*npool.Description, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := appgooddescriptionmwcli.CreateDescription(ctx, &appgooddescriptionmwpb.DescriptionReq{
		EntID:       h.EntID,
		AppGoodID:   h.AppGoodID,
		Description: h.Description,
		Index:       h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetDescription(ctx)
}
