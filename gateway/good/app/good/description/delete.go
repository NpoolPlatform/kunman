package description

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgooddescriptionmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/description"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/description"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDescription(ctx context.Context) (*npool.Description, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDescription(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetDescription(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid description")
	}

	if err := appgooddescriptionmwcli.DeleteDescription(ctx, h.ID, h.EntID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
