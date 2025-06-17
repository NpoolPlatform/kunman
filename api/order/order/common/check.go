package common

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
)

func ValidateAdminCreateOrderType(orderType types.OrderType) error {
	switch orderType {
	case types.OrderType_Offline:
	case types.OrderType_Airdrop:
	default:
		return wlog.Errorf("invalid ordertype")
	}
	return nil
}
