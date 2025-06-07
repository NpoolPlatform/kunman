package fee

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	feecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/fee"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
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
		return wlog.Errorf("fail create fee: %v", err)
	}
	return nil
}

func (h *createHandler) createFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := feecrud.CreateSet(
		tx.Fee.Create(),
		&feecrud.Req{
			EntID:               h.EntID,
			GoodID:              h.GoodID,
			SettlementType:      h.SettlementType,
			UnitValue:           h.UnitValue,
			DurationDisplayType: h.DurationDisplayType,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateFee(ctx context.Context) error {
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
		return handler.createFee(ctx, tx)
	})
}
