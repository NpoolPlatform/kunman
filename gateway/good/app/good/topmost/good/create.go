package topmostgood

import (
	"context"

	topmostcommon "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost/common"
	goodgwcommon "github.com/NpoolPlatform/good-gateway/pkg/common"
	topmostgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good"
	topmostgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	*topmostcommon.CheckHandler
}

func (h *Handler) CreateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	handler := &createHandler{
		Handler: h,
		CheckHandler: &topmostcommon.CheckHandler{
			AppUserCheckHandler: goodgwcommon.AppUserCheckHandler{
				AppID: h.AppID,
			},
			TopMostID: h.TopMostID,
		},
	}
	if err := handler.CheckTopMost(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := topmostgoodmwcli.CreateTopMostGood(ctx, &topmostgoodmwpb.TopMostGoodReq{
		EntID:        h.EntID,
		AppGoodID:    h.AppGoodID,
		TopMostID:    h.TopMostID,
		UnitPrice:    h.UnitPrice,
		DisplayIndex: h.DisplayIndex,
	}); err != nil {
		return nil, err
	}

	return h.GetTopMostGood(ctx)
}
