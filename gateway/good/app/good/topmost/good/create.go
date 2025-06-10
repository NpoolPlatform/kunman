package topmostgood

import (
	"context"

	topmostcommon "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/common"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
	topmostgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"

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

	goodHandler, err := topmostgoodmw.NewHandler(
		ctx,
		topmostgoodmw.WithEntID(h.EntID, true),
		topmostgoodmw.WithAppGoodID(h.AppGoodID, true),
		topmostgoodmw.WithTopMostID(h.TopMostID, true),
		topmostgoodmw.WithUnitPrice(h.UnitPrice, true),
		topmostgoodmw.WithDisplayIndex(h.DisplayIndex, true),
	)
	if err != nil {
		return nil, err
	}

	if err := goodHandler.CreateTopMostGood(ctx); err != nil {
		return nil, err
	}

	return h.GetTopMostGood(ctx)
}
