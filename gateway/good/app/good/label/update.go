package label

import (
	"context"

	appgoodlabelmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/label"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/label"
	appgoodlabelmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/label"
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
