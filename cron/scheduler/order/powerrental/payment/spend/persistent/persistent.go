package persistent

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ledgersvcname "github.com/NpoolPlatform/kunman/middleware/ledger/servicename"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/payment/spend/types"
	powerrentalordermwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	ordersvcname "github.com/NpoolPlatform/kunman/middleware/order/servicename"

	dtmcli "github.com/NpoolPlatform/dtm-cluster/pkg/dtm"
	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"

	"github.com/google/uuid"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) updateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	return powerrentalordermwcli.UpdatePowerRentalOrder(ctx, &powerrentalordermwpb.PowerRentalOrderReq{
		ID:         &order.ID,
		OrderState: ordertypes.OrderState_OrderStateTransferGoodStockLocked.Enum(),
	})
}

func (p *handler) withUpdateOrderState(dispose *dtmcli.SagaDispose, order *types.PersistentOrder) {
	state := ordertypes.OrderState_OrderStateTransferGoodStockLocked
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

func (p *handler) withSpendLockedBalance(dispose *dtmcli.SagaDispose, order *types.PersistentOrder) {
	dispose.Add(
		ledgersvcname.ServiceDomain,
		"ledger.middleware.ledger.v2.Middleware/SettleBalances",
		"",
		&ledgermwpb.SettleBalancesRequest{
			LockID: order.LedgerLockID,
			StatementIDs: func() (statementIDs []string) {
				for range order.PaymentBalances {
					statementIDs = append(statementIDs, uuid.NewString())
				}
				return
			}(),
			IOExtra:   order.BalanceOutcomingExtra,
			IOSubType: ledgertypes.IOSubType_Payment,
		},
	)
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return wlog.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if len(_order.PaymentBalances) == 0 {
		return wlog.WrapError(p.updateOrderState(ctx, _order))
	}

	const timeoutSeconds = 10
	sagaDispose := dtmcli.NewSagaDispose(dtmimp.TransOptions{
		WaitResult:     true,
		RequestTimeout: timeoutSeconds,
	})
	p.withUpdateOrderState(sagaDispose, _order)
	p.withSpendLockedBalance(sagaDispose, _order)
	if err := dtmcli.WithSaga(ctx, sagaDispose); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}
