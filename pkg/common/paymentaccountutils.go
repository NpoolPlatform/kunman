//nolint:dupl
package common

import (
	"context"

	paymentaccountmwcli "github.com/NpoolPlatform/account-middleware/pkg/client/payment"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	paymentaccountmwpb "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func GetPaymentAccounts(ctx context.Context, paymentAccountIDs []string) (map[string]*paymentaccountmwpb.Account, error) {
	for _, paymentAccountID := range paymentAccountIDs {
		if _, err := uuid.Parse(paymentAccountID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	paymentAccounts, _, err := paymentaccountmwcli.GetAccounts(ctx, &paymentaccountmwpb.Conds{
		AccountIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: paymentAccountIDs},
	}, int32(0), int32(len(paymentAccountIDs)))
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	paymentAccountMap := map[string]*paymentaccountmwpb.Account{}
	for _, paymentAccount := range paymentAccounts {
		paymentAccountMap[paymentAccount.AccountID] = paymentAccount
	}
	return paymentAccountMap, nil
}
