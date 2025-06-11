package manufacturer

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	manufacturermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/manufacturer"
	manufacturermw "github.com/NpoolPlatform/kunman/middleware/good/device/manufacturer"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkManufacturer(ctx context.Context) error {
	conds := &manufacturermwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	}
	handler, err := manufacturermw.NewHandler(
		ctx,
		manufacturermw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistManufacturerConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid manufacturer")
	}
	return nil
}
