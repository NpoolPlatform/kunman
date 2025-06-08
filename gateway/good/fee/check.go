package fee

import (
	"context"
	"fmt"

	feemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/fee"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	feemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkFee(ctx context.Context) error {
	exist, err := feemwcli.ExistFeeConds(ctx, &feemwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid fee")
	}
	return nil
}
