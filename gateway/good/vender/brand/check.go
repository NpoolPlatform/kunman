package brand

import (
	"context"
	"fmt"

	brandmwcli "github.com/NpoolPlatform/kunman/middleware/good/vender/brand"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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
