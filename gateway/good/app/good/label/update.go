package label

import (
	"context"

	appgoodlabelmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/label"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/label"
	appgoodlabelmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/label"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateLabel(ctx context.Context) (*npool.Label, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkLabel(ctx); err != nil {
		return nil, err
	}

	if err := appgoodlabelmwcli.UpdateLabel(ctx, &appgoodlabelmwpb.LabelReq{
		ID:           h.ID,
		EntID:        h.EntID,
		Icon:         h.Icon,
		IconBgColor:  h.IconBgColor,
		Label:        h.Label,
		LabelBgColor: h.LabelBgColor,
		Index:        h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetLabel(ctx)
}
