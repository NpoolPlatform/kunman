package goodstm

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
)

var rollbacks = map[types.GoodState]types.GoodState{
	types.GoodState_GoodStateFail:           types.GoodState_GoodStateFail,
	types.GoodState_GoodStateReady:          types.GoodState_GoodStateCheckHashRate,
	types.GoodState_GoodStateCheckHashRate:  types.GoodState_GoodStateCreateGoodUser,
	types.GoodState_GoodStateCreateGoodUser: types.GoodState_GoodStateWait,
}

type rollbackHandler struct {
	*Handler
}

func (h *rollbackHandler) rollback() (*types.GoodState, error) {
	if h.CurrentGoodState == nil || *h.CurrentGoodState != *h.NextGoodState {
		return nil, wlog.Errorf("invalid goodstate")
	}
	state, ok := rollbacks[*h.CurrentGoodState]
	if !ok {
		return nil, wlog.Errorf("invalid goodstate")
	}
	return &state, nil
}
