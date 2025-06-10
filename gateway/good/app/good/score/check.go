package score

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	scoremwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/score"
	scoremw "github.com/NpoolPlatform/kunman/middleware/good/app/good/score"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkScore(ctx context.Context) error {
	conds := &scoremwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	}
	handler, err := scoremw.NewHandler(
		ctx,
		scoremw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistScoreConds(ctx)
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
