package appfiat

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/fiat"
	appfiatmw "github.com/NpoolPlatform/kunman/middleware/chain/app/fiat"
)

func (h *Handler) UpdateFiat(ctx context.Context) (*npool.Fiat, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	// TODO: check appid / fiatid / id

	handler, err := appfiatmw.NewHandler(
		ctx,
		appfiatmw.WithID(h.ID, true),
		appfiatmw.WithName(h.Name, true),
		appfiatmw.WithDisplayNames(h.DisplayNames, true),
		appfiatmw.WithLogo(h.Logo, true),
		appfiatmw.WithDisabled(h.Disabled, true),
		appfiatmw.WithDisplay(h.Display, true),
		appfiatmw.WithDisplayIndex(h.DisplayIndex, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := handler.UpdateFiat(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid fiat")
	}

	h.EntID = &info.EntID

	return h.GetFiat(ctx)
}
