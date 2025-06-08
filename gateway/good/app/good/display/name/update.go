package displayname

import (
	"context"

	appgooddisplaynamemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/display/name"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/display/name"
	appgooddisplaynamemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/name"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateDisplayName(ctx context.Context) (*npool.DisplayName, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDisplayName(ctx); err != nil {
		return nil, err
	}
	if err := appgooddisplaynamemwcli.UpdateDisplayName(ctx, &appgooddisplaynamemwpb.DisplayNameReq{
		ID:    h.ID,
		EntID: h.EntID,
		Name:  h.Name,
		Index: h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetDisplayName(ctx)
}
