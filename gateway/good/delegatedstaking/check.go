package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	delegatedstakingmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/delegatedstaking"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	delegatedstakingmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/delegatedstaking"
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
