package simulate

import (
	"context"
	"fmt"

	apppowerrentalsimulatemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental/simulate"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	apppowerrentalsimulatemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkSimulate(ctx context.Context) error {
	exist, err := apppowerrentalsimulatemwcli.ExistSimulateConds(ctx, &apppowerrentalsimulatemwpb.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid apppowerrentalsimulate")
	}
	return nil
}
