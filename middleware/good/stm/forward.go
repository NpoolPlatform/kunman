package goodstm

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
)

var forwards = map[types.GoodState][]types.GoodState{
	types.GoodState_GoodStatePreWait:        {types.GoodState_GoodStateWait},
	types.GoodState_GoodStateWait:           {types.GoodState_GoodStateCreateGoodUser},
	types.GoodState_GoodStateCreateGoodUser: {types.GoodState_GoodStateCheckHashRate, types.GoodState_GoodStateFail},
	types.GoodState_GoodStateCheckHashRate:  {types.GoodState_GoodStateReady, types.GoodState_GoodStateFail},
}

type forwardHandler struct {
	*Handler
}

func (h *forwardHandler) forward() (*types.GoodState, error) {
	states, ok := forwards[*h.CurrentGoodState]
	if !ok {
		return nil, wlog.Errorf("invalid goodstate")
	}
	for _, state := range states {
		if state == *h.NextGoodState {
			return &state, nil
		}
	}
	return nil, wlog.Errorf("invalid goodstate")
}
