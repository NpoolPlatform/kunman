package orderuser

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools"
	"github.com/google/uuid"
)

func (h *Handler) UpdateOrderUser(ctx context.Context) error {
	info, err := h.GetOrderUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if info == nil {
		return wlog.Errorf("invalid id or ent_id")
	}

	updateH := updateInPoolHandle{
		Handler: h,
	}

	err = updateH.handleUpdateReq(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	gooduserID, err := uuid.Parse(info.GoodUserID)
	if err != nil {
		return wlog.WrapError(err)
	}

	sqlH := h.newSQLHandler()
	sqlH.BondGoodUserID = &gooduserID
	sqlH.BondName = &info.Name

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		sql, err := sqlH.genUpdateSQL()
		if err != nil {
			return wlog.WrapError(err)
		}

		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}

		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return wlog.Errorf("failed to update orderuser: %v", err)
		}
		return nil
	})
}

type updateInPoolHandle struct {
	*Handler
}

//nolint:gocognit
func (h *updateInPoolHandle) handleUpdateReq(ctx context.Context) error {
	handle := &baseInfoHandle{Handler: h.Handler}
	err := handle.getBaseInfo(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	mgr, err := pools.NewPoolManager(handle.baseInfo.MiningPoolType, handle.baseInfo.CoinType.Enum(), handle.baseInfo.AuthToken)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.Proportion != nil {
		err = mgr.SetRevenueProportion(ctx, handle.baseInfo.Distributor, handle.baseInfo.Recipient, handle.Proportion.String())
		if err != nil {
			return wlog.WrapError(err)
		}
	}

	if h.RevenueAddress != nil {
		err = mgr.SetRevenueAddress(ctx, handle.baseInfo.Recipient, *handle.RevenueAddress)
		if err != nil {
			return wlog.WrapError(err)
		}
		if h.AutoPay == nil {
			autopay := true
			h.AutoPay = &autopay
		}
	}

	if h.AutoPay != nil &&
		*h.AutoPay &&
		h.RevenueAddress == nil {
		return wlog.Errorf("cannot set autopay to true without an revenue address")
	}

	if h.AutoPay != nil {
		autoPay := *h.AutoPay
		paused := true

		if autoPay {
			autoPay, err = mgr.ResumePayment(ctx, handle.baseInfo.Recipient)
		} else {
			paused, err = mgr.PausePayment(ctx, handle.baseInfo.Recipient)
		}
		if err != nil {
			return wlog.WrapError(err)
		}
		if !paused {
			autoPay = false
		}

		h.AutoPay = &autoPay
	}

	return nil
}
