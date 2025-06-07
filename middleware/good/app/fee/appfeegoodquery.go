package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappfee "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appfee"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entfee "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/fee"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
)

type appFeeGoodQueryHandler struct {
	*Handler
	fee         *ent.Fee
	goodBase    *ent.GoodBase
	appFee      *ent.AppFee
	appGoodBase *ent.AppGoodBase
}

func (h *appFeeGoodQueryHandler) _getAppFeeGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppFee.Query()
		if h.ID != nil {
			stm.Where(entappfee.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entappfee.EntID(*h.EntID))
		}
		if h.AppGoodID != nil {
			stm.Where(entappfee.AppGoodID(*h.AppGoodID))
		}
		if h.appFee, err = stm.Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		if h.appGoodBase, err = cli.
			AppGoodBase.
			Query().
			Where(
				entappgoodbase.EntID(h.appFee.AppGoodID),
				entappgoodbase.DeletedAt(0),
			).Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *appFeeGoodQueryHandler) getAppFeeGood(ctx context.Context) error {
	if err := h._getAppFeeGood(ctx, false); err != nil {
		return wlog.WrapError(err)
	}
	if h.appGoodBase == nil {
		return nil
	}
	h.AppGoodBaseReq.GoodID = &h.appGoodBase.GoodID
	return h._getFeeGood(ctx, false)
}

func (h *appFeeGoodQueryHandler) requireAppFeeGood(ctx context.Context) error {
	if err := h._getAppFeeGood(ctx, true); err != nil {
		return wlog.WrapError(err)
	}
	h.AppGoodBaseReq.GoodID = &h.appGoodBase.GoodID
	return h._getFeeGood(ctx, true)
}

func (h *appFeeGoodQueryHandler) _getFeeGood(ctx context.Context, must bool) (err error) {
	if h.AppGoodBaseReq.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if h.fee, err = cli.Fee.Query().Where(
			entfee.GoodID(*h.AppGoodBaseReq.GoodID),
			entfee.DeletedAt(0),
		).Only(ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		if h.goodBase, err = cli.GoodBase.Query().Where(
			entgoodbase.EntID(h.fee.GoodID),
			entgoodbase.DeletedAt(0),
		).Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *appFeeGoodQueryHandler) getFeeGood(ctx context.Context) error { //nolint
	return h._getFeeGood(ctx, false)
}

func (h *appFeeGoodQueryHandler) requireFeeGood(ctx context.Context) error {
	return h._getFeeGood(ctx, true)
}

func (h *appFeeGoodQueryHandler) formalizeMinOrderDurationSeconds() error {
	unitSeconds := uint32(1) * 60 * 60

	switch h.fee.DurationDisplayType {
	case types.GoodDurationType_GoodDurationByHour.String():
	case types.GoodDurationType_GoodDurationByDay.String():
		unitSeconds = uint32(1) * 60 * 60 * 24
	case types.GoodDurationType_GoodDurationByMonth.String():
		unitSeconds = uint32(1) * 60 * 60 * 24 * 30
	case types.GoodDurationType_GoodDurationByYear.String():
		unitSeconds = uint32(1) * 60 * 60 * 24 * 365
	default:
		return wlog.Errorf("invalid gooddurationtype")
	}

	if h.MinOrderDurationSeconds == nil {
		h.MinOrderDurationSeconds = &unitSeconds
	}
	if *h.MinOrderDurationSeconds < unitSeconds {
		return wlog.Errorf("invalid minorderdurationseconds < %v", unitSeconds)
	}

	return nil
}
