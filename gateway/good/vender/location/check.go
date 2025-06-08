package location

import (
	"context"
	"fmt"

	locationmwcli "github.com/NpoolPlatform/kunman/middleware/good/vender/location"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	locationmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/location"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkLocation(ctx context.Context) error {
	exist, err := locationmwcli.ExistLocationConds(ctx, &locationmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid location")
	}
	return nil
}
