package label

import (
	"context"
	"fmt"

	appgoodlabelmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/label"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgoodlabelmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/label"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkLabel(ctx context.Context) error {
	exist, err := appgoodlabelmwcli.ExistLabelConds(ctx, &appgoodlabelmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid label")
	}
	return nil
}
