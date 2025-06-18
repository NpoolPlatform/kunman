package persistent

import (
	"context"
	"fmt"

	withdrawmwcli "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	reviewmwpb "github.com/NpoolPlatform/kunman/message/review/middleware/v2/review"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/reviewing/types"
	reviewmwcli "github.com/NpoolPlatform/review-middleware/pkg/client/review"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, withdraw interface{}, reward, notif, done chan interface{}) error {
	_withdraw, ok := withdraw.(*types.PersistentWithdraw)
	if !ok {
		return fmt.Errorf("invalid withdraw")
	}

	defer asyncfeed.AsyncFeed(ctx, _withdraw, done)

	if _withdraw.NeedUpdateReview {
		review, err := reviewmwcli.GetReview(ctx, _withdraw.ReviewID)
		if err != nil {
			return err
		}
		if review == nil {
			return fmt.Errorf("review not found")
		}
		if _, err := reviewmwcli.UpdateReview(ctx, &reviewmwpb.ReviewReq{
			ID:    &review.ID,
			State: &_withdraw.NewReviewState,
		}); err != nil {
			return err
		}
	}
	if _, err := withdrawmwcli.UpdateWithdraw(ctx, &withdrawmwpb.WithdrawReq{
		ID:    &_withdraw.ID,
		State: &_withdraw.NewWithdrawState,
	}); err != nil {
		return err
	}

	return nil
}
