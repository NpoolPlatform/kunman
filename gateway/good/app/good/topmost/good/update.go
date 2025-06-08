package topmostgood

import (
	"context"

	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
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
