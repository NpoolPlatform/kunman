package score

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	scoremwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/score"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	scoremwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/score"

	"github.com/shopspring/decimal"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkScore(ctx context.Context) error {
	exist, err := scoremwcli.ExistScoreConds(ctx, &scoremwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid score")
	}
	return nil
}

func (h *checkHandler) validateScore() error {
	maxScore := decimal.RequireFromString("10.0")
	score, err := decimal.NewFromString(*h.Score)
	if err != nil {
		return wlog.WrapError(err)
	}
	if score.GreaterThan(maxScore) || score.LessThan(decimal.NewFromInt(0)) {
		return wlog.Errorf("invalid score")
	}
	return nil
}
