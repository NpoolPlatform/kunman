package delegatedstaking

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appdelegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/delegatedstaking"
	appdelegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/app/delegatedstaking"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDelegatedStaking(ctx context.Context) error {
	conds := &appdelegatedstakingmwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	}

	handler, err := appdelegatedstakingmw.NewHandler(
		ctx,
		appdelegatedstakingmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistDelegatedStakingConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid appdelegatedstaking")
	}
	return nil
}
