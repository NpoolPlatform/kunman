package oneshot

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	oneshotcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/subscription/oneshot"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/good/goodbase"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sqlGoodBase string
}

func (h *createHandler) constructGoodBaseSQL(ctx context.Context) {
	handler, _ := goodbase1.NewHandler(ctx)

	h.GoodBaseReq.Purchasable = func() *bool { b := true; return &b }()
	h.GoodBaseReq.Online = func() *bool { b := true; return &b }()
	handler.Req = *h.GoodBaseReq

	h.sqlGoodBase = handler.ConstructCreateSQL()
}

func (h *createHandler) createGoodBase(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sqlGoodBase)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create oneshot: %v", err)
	}
	return nil
}

func (h *createHandler) createOneShot(ctx context.Context, tx *ent.Tx) error {
	if _, err := oneshotcrud.CreateSet(
		tx.SubscriptionOneShot.Create(),
		&oneshotcrud.Req{
			EntID:    h.EntID,
			GoodID:   h.GoodID,
			Quota:    h.Quota,
			USDPrice: h.USDPrice,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateOneShot(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.GoodID == nil {
		h.GoodID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		h.GoodBaseReq.EntID = h.GoodID
	}
	if h.GoodBaseReq.ServiceStartAt == nil {
		h.GoodBaseReq.ServiceStartAt = func() *uint32 { u := uint32(time.Now().Unix()); return &u }()
	}

	handler := &createHandler{
		Handler: h,
	}
	handler.constructGoodBaseSQL(ctx)

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createGoodBase(ctx, tx); err != nil {
			return err
		}
		return handler.createOneShot(ctx, tx)
	})
}
