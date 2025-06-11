package brand

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	brandmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/brand"
	brandmw "github.com/NpoolPlatform/kunman/middleware/good/vender/brand"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkBrand(ctx context.Context) error {
	conds := &brandmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	}
	handler, err := brandmw.NewHandler(
		ctx,
		brandmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistBrandConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid brand")
	}
	return nil
}
