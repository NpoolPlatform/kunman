package topmost

import (
	"context"

	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateTopMost(ctx context.Context) (*npool.TopMost, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkTopMost(ctx); err != nil {
		return nil, err
	}

	if err := topmostmwcli.UpdateTopMost(ctx, &topmostmwpb.TopMostReq{
		ID:        h.ID,
		EntID:     h.EntID,
		Title:     h.Title,
		Message:   h.Message,
		TargetUrl: h.TargetURL,
		StartAt:   h.StartAt,
		EndAt:     h.EndAt,
	}); err != nil {
		return nil, err
	}
	return h.GetTopMost(ctx)
}
