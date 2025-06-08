package label

import (
	"context"

	appgoodlabelmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/label"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/label"
	appgoodlabelmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/label"

	"github.com/google/uuid"
)

func (h *Handler) CreateLabel(ctx context.Context) (*npool.Label, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := appgoodlabelmwcli.CreateLabel(ctx, &appgoodlabelmwpb.LabelReq{
		EntID:        h.EntID,
		AppGoodID:    h.AppGoodID,
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
