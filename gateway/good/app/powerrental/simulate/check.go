package simulate

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	apppowerrentalsimulatemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental/simulate"
	apppowerrentalsimulatemw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental/simulate"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkSimulate(ctx context.Context) error {
	conds := &apppowerrentalsimulatemwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	}
	handler, err := apppowerrentalsimulatemw.NewHandler(
		ctx,
		apppowerrentalsimulatemw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistSimulateConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid apppowerrentalsimulate")
	}
	return nil
}
