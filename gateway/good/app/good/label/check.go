package label

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodlabelmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/label"
	appgoodlabelmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/label"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkLabel(ctx context.Context) error {
	conds := &appgoodlabelmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := appgoodlabelmw.NewHandler(
		ctx,
		appgoodlabelmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistLabelConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid label")
	}
	return nil
}
