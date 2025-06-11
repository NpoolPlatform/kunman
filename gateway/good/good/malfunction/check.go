package malfunction

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	malfunctionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/malfunction"
	malfunctionmw "github.com/NpoolPlatform/kunman/middleware/good/good/malfunction"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkMalfunction(ctx context.Context) error {
	conds := &malfunctionmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	}
	handler, err := malfunctionmw.NewHandler(
		ctx,
		malfunctionmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistMalfunctionConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid malfunction")
	}
	return nil
}
