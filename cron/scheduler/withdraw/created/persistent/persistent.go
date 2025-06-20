package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/created/types"
	ledgergwname "github.com/NpoolPlatform/kunman/gateway/ledger/servicename"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	reviewtypes "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	withdrawmw "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
	reviewmw "github.com/NpoolPlatform/kunman/middleware/review/review"

	"github.com/google/uuid"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateWithdrawState(ctx context.Context, withdraw *types.PersistentWithdraw, reviewID string) error {
	state := ledgertypes.WithdrawState_Reviewing

	handler, err := withdrawmw.NewHandler(
		ctx,
		withdrawmw.WithID(&withdraw.ID, true),
		withdrawmw.WithState(&state, true),
		withdrawmw.WithReviewID(&reviewID, true),
	)
	if err != nil {
		return err
	}

	_, err = handler.UpdateWithdraw(ctx)
	return err
}

func (p *handler) withCreateReview(ctx context.Context, withdraw *types.PersistentWithdraw, reviewID string) error {
	serviceName := ledgergwname.ServiceDomain
	objType := reviewtypes.ReviewObjectType_ObjectWithdrawal

	handler, err := reviewmw.NewHandler(
		ctx,
		reviewmw.WithEntID(&reviewID, true),
		reviewmw.WithAppID(&withdraw.AppID, true),
		reviewmw.WithDomain(&serviceName, true),
		reviewmw.WithObjectType(&objType, true),
		reviewmw.WithObjectID(&withdraw.EntID, true),
		reviewmw.WithTrigger(&withdraw.ReviewTrigger, true),
	)
	if err != nil {
		return err
	}

	_, err = handler.CreateReview(ctx)
	return err
}

func (p *handler) Update(ctx context.Context, withdraw interface{}, reward, notif, done chan interface{}) error {
	_withdraw, ok := withdraw.(*types.PersistentWithdraw)
	if !ok {
		return fmt.Errorf("invalid withdraw")
	}

	defer asyncfeed.AsyncFeed(ctx, _withdraw, done)

	reviewID := uuid.NewString()
	if err := p.withCreateReview(ctx, _withdraw, reviewID); err != nil {
		return err
	}
	if err := p.withUpdateWithdrawState(ctx, _withdraw, reviewID); err != nil {
		return err
	}

	asyncfeed.AsyncFeed(ctx, _withdraw, notif)

	return nil
}
