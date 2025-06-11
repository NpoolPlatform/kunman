package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	"github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/rootuser"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	rootusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/rootuser"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPowerRental(ctx context.Context) error {
	conds := &powerrentalmwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID},
	}
	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistPowerRentalConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid powerrental")
	}
	return nil
}

func (h *checkHandler) checkPoolRootUserIDs(ctx context.Context, ids []string) error {
	conds := &rootuser.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: ids},
	}
	handler, err := rootusermw.NewHandler(
		ctx,
		rootusermw.WithConds(conds),
		rootusermw.WithOffset(0),
		rootusermw.WithLimit(int32(len(ids))),
	)
	if err != nil {
		return err
	}

	rootUsers, _, err := handler.GetRootUsers(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	for _, id := range ids {
		exist := false
		for _, rootUser := range rootUsers {
			if rootUser.EntID == id {
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
