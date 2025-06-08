package fee

import (
	"context"
	"fmt"

	feemwcli "github.com/NpoolPlatform/kunman/middleware/good/fee"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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
