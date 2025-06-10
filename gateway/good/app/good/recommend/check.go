package recommend

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	recommendmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/recommend"
	recommendmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/recommend"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkUserRecommend(ctx context.Context) error {
	conds := &recommendmwpb.Conds{
		ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.RecommenderID},
	}
	handler, err := recommendmw.NewHandler(
		ctx,
		recommendmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistRecommendConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid recommend")
	}
	return nil
}
