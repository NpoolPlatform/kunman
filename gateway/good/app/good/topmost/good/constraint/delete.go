package constraint

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	topmostconstraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good/constraint"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good/constraint"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteConstraint(ctx context.Context) (*npool.TopMostGoodConstraint, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkConstraint(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetConstraint(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid constraint")
	}

	if err := topmostconstraintmwcli.DeleteTopMostGoodConstraint(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}

	return info, nil
}
