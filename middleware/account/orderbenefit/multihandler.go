package orderbenefit

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/message/account/middleware/v1/orderbenefit"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
)

type MultiHandler struct {
	Handlers []*Handler
}

type MultiCreateHandler struct {
	*MultiHandler
}

type MultiDeleteHandler struct {
	*MultiHandler
}

func (h *MultiHandler) AppendHandler(handler *Handler) {
	h.Handlers = append(h.Handlers, handler)
}

func (h *MultiHandler) GetHandlers() []*Handler {
	return h.Handlers
}

func NewMultiCreateHandler(ctx context.Context, reqs []*orderbenefit.AccountReq, must bool) (*MultiCreateHandler, error) {
	mh := &MultiHandler{}
	if len(reqs) == 0 && must {
		return nil, fmt.Errorf("invalid reqs")
	}

	for _, req := range reqs {
		handler, err := NewHandler(
			ctx,
			WithEntID(req.EntID, false),
			WithAppID(req.AppID, true),
			WithUserID(req.UserID, true),
			WithCoinTypeID(req.CoinTypeID, false),
			WithAccountID(req.AccountID, false),
			WithOrderID(req.OrderID, true),
			WithAddress(req.Address, false),
			WithActive(req.Active, false),
			WithLocked(req.Locked, false),
			WithBlocked(req.Blocked, false),
		)
		if err != nil {
			return nil, err
		}
		mh.AppendHandler(handler)
	}

	return &MultiCreateHandler{mh}, nil
}

func NewMultiDeleteHandler(ctx context.Context, reqs []*orderbenefit.AccountReq, must bool) (*MultiDeleteHandler, error) {
	mh := &MultiHandler{}
	if len(reqs) == 0 && must {
		return nil, fmt.Errorf("invalid reqs")
	}

	for _, req := range reqs {
		handler, err := NewHandler(
			ctx,
			WithID(req.ID, true),
			WithAppID(req.AppID, true),
			WithUserID(req.UserID, true),
			WithOrderID(req.OrderID, true),
			WithAccountID(req.AccountID, false),
		)
		if err != nil {
			return nil, err
		}
		mh.AppendHandler(handler)
	}

	return &MultiDeleteHandler{mh}, nil
}

func (h *MultiCreateHandler) CreateAccountsWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.CreateAccountWithTx(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (h *MultiCreateHandler) CreateAccounts(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateAccountsWithTx(_ctx, tx)
	})
}

func (h *MultiDeleteHandler) DeleteAccountsWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.DeleteAccountWithTx(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (h *MultiDeleteHandler) DeleteAccounts(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteAccountsWithTx(_ctx, tx)
	})
}
