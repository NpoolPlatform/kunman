package achievement

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	orderstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	orderstatement1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/statement/order"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type expropriateHandler struct {
	*Handler
	statements []*orderstatementmwpb.Statement
}

func (h *expropriateHandler) getOrderStatements(ctx context.Context) error {
	orderStatementHandler, err := orderstatement1.NewHandler(
		ctx,
		orderstatement1.WithConds(&orderstatementmwpb.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: h.OrderID.String()},
		}),
		orderstatement1.WithLimit(constant.DefaultRowLimit),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	for {
		statements, _, err := orderStatementHandler.GetStatements(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(statements) == 0 {
			break
		}
		h.statements = append(h.statements, statements...)
		orderStatementHandler.Offset += orderStatementHandler.Limit
	}
	return nil
}

func (h *Handler) ExpropriateAchievement(ctx context.Context) error {
	handler := &expropriateHandler{
		Handler: h,
	}
	if err := handler.getOrderStatements(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if len(handler.statements) == 0 {
		return nil
	}

	multiHandler := &orderstatement1.MultiHandler{}
	for _, statement := range handler.statements {
		h1, err := orderstatement1.NewHandler(
			ctx,
			orderstatement1.WithEntID(&statement.EntID, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		multiHandler.AppendHandler(h1)
	}
	return multiHandler.DeleteStatements(ctx)
}
