package persistent

import (
	"context"

	dtmcli "github.com/NpoolPlatform/dtm-cluster/pkg/dtm"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	orderusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	orderusersvcname "github.com/NpoolPlatform/kunman/middleware/miningpool/servicename"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/miningpool/createorderuser/types"
	ordersvcname "github.com/NpoolPlatform/kunman/middleware/order/servicename"
	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withCreateOrderUser(dispose *dtmcli.SagaDispose, req *orderusermwpb.OrderUserReq) {
	dispose.Add(
		orderusersvcname.ServiceDomain,
		"miningpool.middleware.orderuser.v1.Middleware/CreateOrderUser",
		"miningpool.middleware.orderuser.v1.Middleware/DeleteOrderUser",
		&orderusermwpb.CreateOrderUserRequest{
			Info: req,
		},
	)
}

func (p *handler) withUpdateOrder(dispose *dtmcli.SagaDispose, order *powerrentalordermwpb.PowerRentalOrderReq) {
	rollback := true
	order.Rollback = &rollback
	dispose.Add(
		ordersvcname.ServiceDomain,
		"order.middleware.powerrental.v1.Middleware/UpdatePowerRentalOrder",
		"order.middleware.powerrental.v1.Middleware/UpdatePowerRentalOrder",
		&powerrentalordermwpb.UpdatePowerRentalOrderRequest{
			Info: order,
		},
	)
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return wlog.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	const timeoutSeconds = 10
	sagaDispose := dtmcli.NewSagaDispose(dtmimp.TransOptions{
		WaitResult:     true,
		RequestTimeout: timeoutSeconds,
	})

	if _order.OrderUserReq != nil {
		p.withCreateOrderUser(sagaDispose, _order.OrderUserReq)
	}
	if _order.PowerRentalOrderReq != nil {
		p.withUpdateOrder(sagaDispose, _order.PowerRentalOrderReq)
	}

	if err := dtmcli.WithSaga(ctx, sagaDispose); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}
