package persistent

import (
	"context"
	"fmt"

	ledgersvcname "github.com/NpoolPlatform/kunman/middleware/ledger/servicename"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	dtm1 "github.com/NpoolPlatform/kunman/cron/scheduler/dtm"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/cancel/commission/types"
	ordersvcname "github.com/NpoolPlatform/kunman/middleware/order/servicename"

	dtmcli "github.com/NpoolPlatform/dtm-cluster/pkg/dtm"
	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(dispose *dtmcli.SagaDispose, order *types.PersistentPowerRentalOrder) {
	state := ordertypes.OrderState_OrderStateCancelAchievement
	rollback := true
	req := &powerrentalordermwpb.PowerRentalOrderReq{
		ID:         &order.ID,
		OrderState: &state,
		Rollback:   &rollback,
	}
	dispose.Add(
		ordersvcname.ServiceDomain,
		"order.middleware.powerrental.v1.Middleware/UpdatePowerRentalOrder",
		"order.middleware.powerrental.v1.Middleware/UpdatePowerRentalOrder",
		&powerrentalordermwpb.UpdatePowerRentalOrderRequest{
			Info: req,
		},
	)
}

func (p *handler) withDeductLockedCommission(dispose *dtmcli.SagaDispose, order *types.PersistentPowerRentalOrder) {
	for _, revoke := range order.CommissionRevokes {
		dispose.Add(
			ledgersvcname.ServiceDomain,
			"ledger.middleware.ledger.v2.Middleware/SettleBalances",
			"",
			&ledgermwpb.SettleBalancesRequest{
				LockID:       revoke.LockID,
				IOSubType:    ledgertypes.IOSubType_CommissionRevoke,
				IOExtra:      revoke.IOExtra,
				StatementIDs: revoke.StatementIDs,
			},
		)
	}
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentPowerRentalOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	const timeoutSeconds = 10
	sagaDispose := dtmcli.NewSagaDispose(dtmimp.TransOptions{
		WaitResult:     true,
		RequestTimeout: timeoutSeconds,
		TimeoutToFail:  timeoutSeconds,
		RetryInterval:  timeoutSeconds,
	})
	p.withUpdateOrderState(sagaDispose, _order)
	p.withDeductLockedCommission(sagaDispose, _order)
	if err := dtm1.Do(ctx, sagaDispose); err != nil {
		return err
	}

	return nil
}
