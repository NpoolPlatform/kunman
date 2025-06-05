package subscription

import (
	"context"
	"fmt"

	submwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/user/subscription"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	submwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/user/subscription"
)

func (h *Handler) GetSubscriptionsCount(ctx context.Context) (uint32, error) {
	conds := &submwpb.Conds{}

	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	return submwcli.GetSubscriptionsCount(ctx, conds)
}
