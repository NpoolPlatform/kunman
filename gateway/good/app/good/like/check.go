package like

import (
	"context"
	"fmt"

	likemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/like"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	likemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/like"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkUserLike(ctx context.Context) error {
	exist, err := likemwcli.ExistLikeConds(ctx, &likemwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid like")
	}
	return nil
}
