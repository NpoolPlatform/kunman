package poster

import (
	"context"
	"fmt"

	appgoodpostermwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/poster"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodpostermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/poster"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPoster(ctx context.Context) error {
	exist, err := appgoodpostermwcli.ExistPosterConds(ctx, &appgoodpostermwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid poster")
	}
	return nil
}
