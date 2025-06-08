package displayname

import (
	"context"

	appgooddisplaynamemwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/display/name"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
	appgooddisplaynamemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/display/name"

	"github.com/google/uuid"
)

func (h *Handler) CreateDisplayName(ctx context.Context) (*npool.DisplayName, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := appgooddisplaynamemwcli.CreateDisplayName(ctx, &appgooddisplaynamemwpb.DisplayNameReq{
		EntID:     h.EntID,
		AppGoodID: h.AppGoodID,
		Name:      h.Name,
		Index:     h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetDisplayName(ctx)
}
