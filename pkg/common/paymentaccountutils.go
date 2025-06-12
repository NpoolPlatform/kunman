//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	paymentaccountmw "github.com/NpoolPlatform/kunman/middleware/account/payment"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetPaymentAccounts(ctx context.Context, paymentAccountIDs []string) (map[string]*paymentaccountmwpb.Account, error) {
	for _, paymentAccountID := range paymentAccountIDs {
		if _, err := uuid.Parse(paymentAccountID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &paymentaccountmwpb.Conds{
		AccountIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: paymentAccountIDs},
	}
	handler, err := paymentaccountmw.NewHandler(
		ctx,
		paymentaccountmw.WithConds(conds),
		paymentaccountmw.WithOffset(0),
		paymentaccountmw.WithLimit(int32(len(paymentAccountIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	paymentAccounts, _, err := handler.GetAccounts(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	paymentAccountMap := map[string]*paymentaccountmwpb.Account{}
	for _, paymentAccount := range paymentAccounts {
		paymentAccountMap[paymentAccount.AccountID] = paymentAccount
	}
	return paymentAccountMap, nil
}
