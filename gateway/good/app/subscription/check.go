package subscription

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appsubscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription"
	appsubscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkSubscription(ctx context.Context) error {
	conds := &appsubscriptionmwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	}
	handler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistSubscriptionConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid appsubscription")
	}
	return nil
}
