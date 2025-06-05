package subscription

import (
	"context"
	"fmt"

	subscriptionmwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/user/subscription"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	subscriptionmwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/user/subscription"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkSubscription(ctx context.Context) error {
	exist, err := subscriptionmwcli.ExistSubscriptionConds(ctx, &subscriptionmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid subscription")
	}
	return nil
}
