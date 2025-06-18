package types

import (
	miningstockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/stock"
	goodpowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
)

type PersistentGoodPowerRental struct {
	*goodpowerrentalmwpb.PowerRental
	MiningGoodStockReqs []*miningstockmwpb.MiningGoodStockReq
}
