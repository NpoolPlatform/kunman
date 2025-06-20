package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	orderbenefitmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/orderbenefit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	orderbenefitmw "github.com/NpoolPlatform/kunman/middleware/account/orderbenefit"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func GetOrderBenefits(ctx context.Context, orderIDs []string) (map[string][]*orderbenefitmwpb.Account, error) {
	orderBenefitMap := make(map[string][]*orderbenefitmwpb.Account)

	handler, err := orderbenefitmw.NewHandler(
		ctx,
		orderbenefitmw.WithConds(
			&orderbenefitmwpb.Conds{
				OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: orderIDs},
			},
		),
		orderbenefitmw.WithOffset(0),
		// TODO: we should get all benefit accounts here
		orderbenefitmw.WithLimit(int32(len(orderIDs)*10)),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	infos, _, err := handler.GetAccounts(ctx)
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		accounts, ok := orderBenefitMap[info.OrderID]
		if !ok {
			accounts = []*orderbenefitmwpb.Account{}
		}

		accounts = append(accounts, info)
		orderBenefitMap[info.OrderID] = accounts
	}

	return orderBenefitMap, nil
}
