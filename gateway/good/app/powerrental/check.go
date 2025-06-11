package powerrental

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPowerRental(ctx context.Context) error {
	conds := &apppowerrentalmwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	}
	handler, err := apppowerrentalmw.NewHandler(
		ctx,
		apppowerrentalmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistPowerRentalConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid apppowerrental")
	}
	return nil
}
