package delegatedstaking

import (
	"time"

	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type DelegatedStaking interface {
	GoodServiceStartAt() uint32
	GoodStartMode() types.GoodStartMode
}

type delegatedstaking struct {
	delegatedstaking    *ent.DelegatedStaking
	goodBase            *ent.GoodBase
	appGoodBase         *ent.AppGoodBase
	appDelegatedStaking *ent.AppDelegatedStaking
}

func (pr *delegatedstaking) GoodServiceStartAt() uint32 {
	now := uint32(time.Now().Unix())
	if now < pr.goodBase.ServiceStartAt {
		return pr.goodBase.ServiceStartAt
	}
	return now
}

func (pr *delegatedstaking) GoodStartMode() types.GoodStartMode {
	return types.GoodStartMode(types.GoodStartMode_value[pr.goodBase.StartMode])
}
