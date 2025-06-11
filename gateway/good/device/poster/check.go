package poster

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	devicepostermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"
	devicepostermw "github.com/NpoolPlatform/kunman/middleware/good/device/poster"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPoster(ctx context.Context) error {
	conds := &devicepostermwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	}
	handler, err := devicepostermw.NewHandler(
		ctx,
		devicepostermw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistPosterConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid poster")
	}
	return nil
}
