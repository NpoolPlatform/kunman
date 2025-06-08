package comment

import (
	"context"
	"fmt"

	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/comment"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	commentmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/comment"
	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordermwcli "github.com/NpoolPlatform/order-middleware/pkg/client/order"
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
	exist, err := ordermwcli.ExistOrderConds(ctx, conds)
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
	exist, err := commentmwcli.ExistCommentConds(ctx, conds)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid comment")
	}
	return nil
}
