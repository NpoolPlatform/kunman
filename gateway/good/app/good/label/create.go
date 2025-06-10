package label

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/label"
	appgoodlabelmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/label"

	"github.com/google/uuid"
)

func (h *Handler) CreateLabel(ctx context.Context) (*npool.Label, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := appgoodlabelmw.NewHandler(
		ctx,
		appgoodlabelmw.WithEntID(h.EntID, true),
		appgoodlabelmw.WithAppGoodID(h.AppGoodID, true),
		appgoodlabelmw.WithIcon(h.Icon, true),
		appgoodlabelmw.WithIconBgColor(h.IconBgColor, true),
		appgoodlabelmw.WithLabel(h.Label, true),
		appgoodlabelmw.WithLabelBgColor(h.LabelBgColor, true),
		appgoodlabelmw.WithIndex(func() *uint8 {
			if h.Index == nil {
				return nil
			}
			index := uint8(*h.Index)
			return &index
		}(), true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateLabel(ctx); err != nil {
		return nil, err
	}
	return h.GetLabel(ctx)
}
