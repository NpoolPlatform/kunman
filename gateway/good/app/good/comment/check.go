package comment

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	commentmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/comment"
	ordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order"
	commentmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/comment"
	ordermw "github.com/NpoolPlatform/kunman/middleware/order/order"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkOrder(ctx context.Context) error {
	conds := &ordermwpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
		UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: *h.CommentUserID},
	}
	if h.OrderID != nil {
		conds.EntID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderID}
	}

	handler, err := ordermw.NewHandler(
		ctx,
		ordermw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistOrderConds(ctx)
	if err != nil {
		return err
	}
	if !exist && h.OrderID != nil {
		return fmt.Errorf("order not matched")
	}
	return nil
}

func (h *checkHandler) checkUserComment(ctx context.Context) error {
	conds := &commentmwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.CommentUserID},
	}

	handler, err := commentmw.NewHandler(
		ctx,
		commentmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistCommentConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid comment")
	}
	return nil
}
