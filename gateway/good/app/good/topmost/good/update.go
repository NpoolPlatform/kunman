package topmostgood

import (
	"context"

	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkTopMostGood(ctx); err != nil {
		return nil, err
	}

	if err := topmostgoodmwcli.UpdateTopMostGood(ctx, &topmostgoodmwpb.TopMostGoodReq{
		ID:           h.ID,
		EntID:        h.EntID,
		UnitPrice:    h.UnitPrice,
		DisplayIndex: h.DisplayIndex,
	}); err != nil {
		return nil, err
	}
	return h.GetTopMostGood(ctx)
}
