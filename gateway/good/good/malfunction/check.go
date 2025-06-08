package malfunction

import (
	"context"
	"fmt"

	malfunctionmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/malfunction"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	malfunctionmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/malfunction"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkMalfunction(ctx context.Context) error {
	exist, err := malfunctionmwcli.ExistMalfunctionConds(ctx, &malfunctionmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid malfunction")
	}
	return nil
}
