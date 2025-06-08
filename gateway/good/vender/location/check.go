package location

import (
	"context"
	"fmt"

	locationmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	locationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
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
