package common

import (
	"context"

	compensatemw "github.com/NpoolPlatform/kunman/middleware/order/compensate"

	"github.com/google/uuid"
)

func GetCompensateOrderNumbers(ctx context.Context, compensateFromIDs []string) (map[string]uint32, error) {
	for _, compensateFromID := range compensateFromIDs {
		if _, err := uuid.Parse(compensateFromID); err != nil {
			return nil, err
		}
	}

	handler, err := compensatemw.NewHandler(
		ctx,
		compensatemw.WithCompensateFromIDs(compensateFromIDs, true),
	)
	if err != nil {
		return nil, err
	}

	orderNumbers, err := handler.CountCompensateOrders(ctx)
	if err != nil {
		return nil, err
	}
	orderNumberMap := map[string]uint32{}
	for _, orderNumber := range orderNumbers {
		orderNumberMap[orderNumber.CompensateFromID] = orderNumber.Orders
	}
	return orderNumberMap, nil
}
