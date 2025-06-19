package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/payment/commission/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	orderpaymentstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order/payment"
	ledgerstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	orderpaymentstatementmw "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/statement/order/payment"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent        chan interface{}
	notif             chan interface{}
	done              chan interface{}
	paymentStatements []*orderpaymentstatementmwpb.Statement
	ledgerStatements  []*ledgerstatementmwpb.StatementReq
}

func (h *orderHandler) getOrderPaymentStatements(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &orderpaymentstatementmwpb.Conds{
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: h.OrderID},
	}

	for {
		handler, err := orderpaymentstatementmw.NewHandler(
			ctx,
			orderpaymentstatementmw.WithConds(conds),
			orderpaymentstatementmw.WithOffset(offset),
			orderpaymentstatementmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		statements, _, err := handler.GetStatements(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(statements) == 0 {
			return nil
		}
		h.paymentStatements = append(h.paymentStatements, statements...)
		offset += limit
	}
}

func (h *orderHandler) constructLedgerStatements() error {
	ioType := ledgertypes.IOType_Incoming
	ioSubType := ledgertypes.IOSubType_Commission

	for _, statement := range h.paymentStatements {
		amount, err := decimal.NewFromString(statement.CommissionAmount)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			continue
		}
		ioExtra := fmt.Sprintf(
			`{"PaymentID":"%v","OrderID":"%v","OrderUserID":"%v","InspireAppConfigID":"%v","CommissionConfigID":"%v","CommissionConfigType":"%v","PaymentStatementID":"%v"}`,
			h.PaymentID,
			h.OrderID,
			h.UserID,
			statement.AppConfigID,
			statement.CommissionConfigID,
			statement.CommissionConfigType,
			statement.EntID,
		)
		h.ledgerStatements = append(h.ledgerStatements, &ledgerstatementmwpb.StatementReq{
			AppID:      &h.AppID,
			UserID:     &statement.UserID,
			CoinTypeID: &statement.PaymentCoinTypeID,
			IOType:     &ioType,
			IOSubType:  &ioSubType,
			Amount:     &statement.CommissionAmount,
			IOExtra:    &ioExtra,
		})
	}
	return nil
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRentalOrder", h.PowerRentalOrder,
			"LedgerStatements", h.ledgerStatements,
			"OrderStatements", h.paymentStatements,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		PowerRentalOrder: h.PowerRentalOrder,
		LedgerStatements: h.ledgerStatements,
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.notif)
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.done)
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	if err := h.getOrderPaymentStatements(ctx); err != nil {
		return err
	}
	if err = h.constructLedgerStatements(); err != nil {
		return err
	}

	return nil
}
