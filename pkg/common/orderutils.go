package common

import (
	"context"

	compensatemwcli "github.com/NpoolPlatform/order-middleware/pkg/client/compensate"

	"github.com/google/uuid"
)

func GetCompensateOrderNumbers(ctx context.Context, compensateFromIDs []string) (map[string]uint32, error) {
	for _, compensateFromID := range compensateFromIDs {
		if _, err := uuid.Parse(compensateFromID); err != nil {
			return nil, err
		}
	}

	orderNumbers, err := compensatemwcli.CountCompensateOrders(ctx, compensateFromIDs)
	if err != nil {
		return nil, err
	}
	orderNumberMap := map[string]uint32{}
	for _, orderNumber := range orderNumbers {
		orderNumberMap[orderNumber.CompensateFromID] = orderNumber.Orders
	}
	return orderNumberMap, nil
}
