package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/reviewing/types"
	withdrawmw "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
	reviewmw "github.com/NpoolPlatform/kunman/middleware/review/review"
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
		reviewHandler, err := reviewmw.NewHandler(
			ctx,
			reviewmw.WithEntID(&_withdraw.ReviewID, true),
		)
		if err != nil {
			return err
		}

		review, err := reviewHandler.GetReview(ctx)
		if err != nil {
			return err
		}
		if review == nil {
			return fmt.Errorf("review not found")
		}

		reviewHandler, err = reviewmw.NewHandler(
			ctx,
			reviewmw.WithID(&review.ID, true),
			reviewmw.WithState(&_withdraw.NewReviewState, true),
		)
		if err != nil {
			return err
		}

		if _, err := reviewHandler.UpdateReview(ctx); err != nil {
			return err
		}
	}

	handler, err := withdrawmw.NewHandler(
		ctx,
		withdrawmw.WithID(&_withdraw.ID, true),
		withdrawmw.WithState(&_withdraw.NewWithdrawState, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateWithdraw(ctx); err != nil {
		return err
	}

	return nil
}
