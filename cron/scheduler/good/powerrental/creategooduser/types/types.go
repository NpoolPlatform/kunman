package types

import (
	miningstockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/stock"
	goodpowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	goodusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/gooduser"
)

type PersistentGoodPowerRental struct {
	*goodpowerrentalmwpb.PowerRental
	MiningGoodStockReqs []*miningstockmwpb.MiningGoodStockReq
	GoodUserReqs        []*goodusermwpb.GoodUserReq
}
