package recommend

import (
	"context"
	"fmt"

	recommendmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/recommend"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	recommendmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/recommend"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkUserRecommend(ctx context.Context) error {
	exist, err := recommendmwcli.ExistRecommendConds(ctx, &recommendmwpb.Conds{
		ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.RecommenderID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid recommend")
	}
	return nil
}
