package constraint

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/constraint"
	topmostconstraintmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good/constraint"
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

	constraintHandler, err := topmostconstraintmw.NewHandler(
		ctx,
		topmostconstraintmw.WithID(h.ID, true),
		topmostconstraintmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := constraintHandler.DeleteConstraint(ctx); err != nil {
		return nil, err
	}

	return info, nil
}
