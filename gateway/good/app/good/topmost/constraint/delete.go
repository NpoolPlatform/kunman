package constraint

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	topmostconstraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/constraint"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/constraint"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteConstraint(ctx context.Context) (*npool.TopMostConstraint, error) {
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

	if err := topmostconstraintmwcli.DeleteTopMostConstraint(ctx, h.ID, h.EntID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
