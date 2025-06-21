package appfiat

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/fiat"
	appfiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/fiat"
	appfiatmw "github.com/NpoolPlatform/kunman/middleware/chain/app/fiat"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type queryHandler struct {
	*Handler
	fiats []*appfiatmwpb.Fiat
	infos []*npool.Fiat
	total uint32
}

func (h *queryHandler) formalize() {
	for _, info := range h.fiats {
		_info := &npool.Fiat{
			ID:           info.ID,
			EntID:        info.EntID,
			AppID:        info.AppID,
			FiatID:       info.FiatID,
			Name:         info.Name,
			DisplayNames: info.DisplayNames,
			Logo:         info.Logo,
			Unit:         info.Unit,
			Disabled:     info.Disabled,
			CreatedAt:    info.CreatedAt,
			UpdatedAt:    info.UpdatedAt,
			Display:      info.Display,
			DisplayIndex: info.DisplayIndex,
		}
		h.infos = append(h.infos, _info)
	}
}

func (h *Handler) GetFiat(ctx context.Context) (*npool.Fiat, error) {
	if h.EntID == nil {
		return nil, fmt.Errorf("invalid entid")
	}

	appFiatHandler, err := appfiatmw.NewHandler(
		ctx,
		appfiatmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := appFiatHandler.GetFiat(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	h.AppID = &info.AppID

	handler := &queryHandler{
		Handler: h,
		fiats:   []*appfiatmwpb.Fiat{info},
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetFiats(ctx context.Context) ([]*npool.Fiat, uint32, error) {
	conds := &appfiatmwpb.Conds{}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}

	appFiatHandler, err := appfiatmw.NewHandler(
		ctx,
		appfiatmw.WithConds(conds),
		appfiatmw.WithOffset(h.Offset),
		appfiatmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	fiats, total, err := appFiatHandler.GetFiats(ctx)
	if err != nil {
		return nil, 0, err
	}

	handler := &queryHandler{
		Handler: h,
		fiats:   fiats,
		total:   total,
	}

	handler.formalize()

	return handler.infos, total, nil
}

func (h *Handler) GetFiatExt(ctx context.Context, info *appfiatmwpb.Fiat) (*npool.Fiat, error) {
	h.AppID = &info.AppID

	handler := &queryHandler{
		Handler: h,
		fiats:   []*appfiatmwpb.Fiat{info},
	}

	handler.formalize()

	return handler.infos[0], nil
}
