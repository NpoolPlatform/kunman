package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	powerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	powerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
	"github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPowerRental(ctx context.Context) error {
	exist, err := powerrentalmwcli.ExistPowerRentalConds(ctx, &powerrentalmwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid powerrental")
	}
	return nil
}

func (h *checkHandler) checkPoolRootUserIDs(ctx context.Context, ids []string) error {
	ruInfos, _, err := rootusermwcli.GetRootUsers(ctx, &rootuser.Conds{
		EntIDs: &basetypes.StringSliceVal{
			Op:    cruder.IN,
			Value: ids,
		},
	}, 0, int32(len(ids)))
	if err != nil {
		return wlog.WrapError(err)
	}

	for _, id := range ids {
		exist := false
		for _, ruInfo := range ruInfos {
			if ruInfo.EntID == id {
				exist = true
				break
			}
		}
		if !exist {
			return wlog.Errorf("invalid pool rootuserid")
		}
	}
	return nil
}
