package location

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	locationmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/location"
	locationmw "github.com/NpoolPlatform/kunman/middleware/good/vender/location"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkLocation(ctx context.Context) error {
	conds := &locationmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	}
	handler, err := locationmw.NewHandler(
		ctx,
		locationmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistLocationConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid location")
	}
	return nil
}
