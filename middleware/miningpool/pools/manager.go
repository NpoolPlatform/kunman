package pools

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	mpbasetype "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	basetype "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/types"
)

// TODO: support use default coinType
func NewPoolManager(poolType mpbasetype.MiningPoolType, coinType *basetype.CoinType, auth string) (types.PoolManager, error) {
	if poolType == mpbasetype.MiningPoolType_F2Pool {
		return f2pool.NewF2PoolManager(coinType, auth)
	}
	return nil, wlog.Errorf("has not implemented for %v-%v ", poolType, coinType)
}
