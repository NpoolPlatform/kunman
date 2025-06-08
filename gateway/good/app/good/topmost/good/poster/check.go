package poster

import (
	"context"
	"fmt"

	topmostgoodpostermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good/poster"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostgoodpostermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/poster"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPoster(ctx context.Context) error {
	exist, err := topmostgoodpostermwcli.ExistPosterConds(ctx, &topmostgoodpostermwpb.Conds{
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
