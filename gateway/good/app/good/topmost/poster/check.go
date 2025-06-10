package poster

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	topmostpostermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/poster"
	topmostpostermw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/poster"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPoster(ctx context.Context) error {
	conds := &topmostpostermwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := topmostpostermw.NewHandler(
		ctx,
		topmostpostermw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistPosterConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid topmostposter")
	}
	return nil
}
