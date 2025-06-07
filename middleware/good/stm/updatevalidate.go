package goodstm

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
)

type validateHandler struct {
	*Handler
	*rollbackHandler
	*forwardHandler
}

func (h *validateHandler) validateGoodState() (*types.GoodState, error) {
	if h.NextGoodState == nil {
		return nil, nil
	}
	if h.Rollback != nil && *h.Rollback && *h.CurrentGoodState == *h.NextGoodState {
		return h.rollback()
	}
	return h.forward()
}

func (h *Handler) ValidateUpdateForNewState() (*types.GoodState, error) {
	handler := &validateHandler{
		Handler:         h,
		rollbackHandler: &rollbackHandler{Handler: h},
		forwardHandler:  &forwardHandler{Handler: h},
	}
	state, err := handler.validateGoodState()
	return state, wlog.WrapError(err)
}
