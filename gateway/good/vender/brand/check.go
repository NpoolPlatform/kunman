package brand

import (
	"context"
	"fmt"

	brandmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	brandmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/brand"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkBrand(ctx context.Context) error {
	exist, err := brandmwcli.ExistBrandConds(ctx, &brandmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid brand")
	}
	return nil
}
