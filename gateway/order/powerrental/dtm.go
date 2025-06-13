package powerrental

import (
	"context"
	"time"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	dtmcli "github.com/NpoolPlatform/dtm-cluster/pkg/dtm"
)

type dtmHandler struct {
	*checkHandler
}

func (h *dtmHandler) dtmDo(ctx context.Context, dispose *dtmcli.SagaDispose) error {
	start := time.Now()
	_ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	err := dtmcli.WithSaga(_ctx, dispose)
	dtmElapsed := time.Since(start)
	logger.Sugar().Infow(
		"CreatePowerRentalOrderWithFees",
		"OrderID", *h.OrderID,
		"Start", start,
		"DtmElapsed", dtmElapsed,
		"Error", err,
	)
	return wlog.WrapError(err)
}
