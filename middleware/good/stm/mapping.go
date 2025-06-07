package goodstm

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
)

var stateMap = map[types.GoodState]types.MiningGoodStockState{
	types.GoodState_DefaultGoodState:        types.MiningGoodStockState_DefaultMiningGoodStockState,
	types.GoodState_GoodStatePreWait:        types.MiningGoodStockState_MiningGoodStockStatePreWait,
	types.GoodState_GoodStateWait:           types.MiningGoodStockState_MiningGoodStockStateWait,
	types.GoodState_GoodStateFail:           types.MiningGoodStockState_MiningGoodStockStateFail,
	types.GoodState_GoodStateReady:          types.MiningGoodStockState_MiningGoodStockStateReady,
	types.GoodState_GoodStateCheckHashRate:  types.MiningGoodStockState_MiningGoodStockStateCheckHashRate,
	types.GoodState_GoodStateCreateGoodUser: types.MiningGoodStockState_MiningGoodStockStateCreateGoodUser,
}

func GoodState2MiningGoodStockState(goodstate *types.GoodState) (*types.MiningGoodStockState, error) {
	_state, ok := stateMap[*goodstate]
	if !ok {
		return nil, wlog.Errorf("invalid goodstate")
	}
	return &_state, nil
}
