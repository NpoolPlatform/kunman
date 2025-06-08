package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	delegatedstakingmwcli "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	delegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/delegatedstaking"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkDelegatedStaking(ctx context.Context) error {
	exist, err := delegatedstakingmwcli.ExistDelegatedStakingConds(ctx, &delegatedstakingmwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid delegatedstaking")
	}
	return nil
}
