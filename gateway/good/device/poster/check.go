package poster

import (
	"context"
	"fmt"

	devicepostermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/poster"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	devicepostermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/poster"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPoster(ctx context.Context) error {
	exist, err := devicepostermwcli.ExistPosterConds(ctx, &devicepostermwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid poster")
	}
	return nil
}
