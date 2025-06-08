package poster

import (
	"context"
	"fmt"

	topmostpostermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/poster"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostpostermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/poster"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPoster(ctx context.Context) error {
	exist, err := topmostpostermwcli.ExistPosterConds(ctx, &topmostpostermwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmostposter")
	}
	return nil
}
