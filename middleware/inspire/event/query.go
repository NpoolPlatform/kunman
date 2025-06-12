package event

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"
	eventcoinmw "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event/coin"
	eventcouponmw "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event/coupon"
	eventcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event"
	eventcoincrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event/coin"
	eventcouponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event/coupon"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entevent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/event"
	eventcoin1 "github.com/NpoolPlatform/kunman/middleware/inspire/event/coin"
	eventcoupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/event/coupon"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.EventSelect
	stmSelect *ent.EventSelect
	infos     []*npool.Event
	total     uint32
}

func (h *queryHandler) selectEvent(stm *ent.EventQuery) *ent.EventSelect {
	return stm.Select(
		entevent.FieldID,
	)
}

func (h *queryHandler) queryEvent(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}

	stm := cli.Event.Query().Where(entevent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entevent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entevent.EntID(*h.EntID))
	}
	h.stmSelect = h.selectEvent(stm)
	return nil
}

func (h *queryHandler) queryEvents(cli *ent.Client) (*ent.EventSelect, error) {
	stm, err := eventcrud.SetQueryConds(cli.Event.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectEvent(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entevent.Table)
	s.LeftJoin(t).
		On(
			s.C(entevent.FieldID),
			t.C(entevent.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entevent.FieldEntID), "ent_id"),
			sql.As(t.C(entevent.FieldAppID), "app_id"),
			sql.As(t.C(entevent.FieldEventType), "event_type"),
			sql.As(t.C(entevent.FieldCredits), "credits"),
			sql.As(t.C(entevent.FieldCreditsPerUsd), "credits_per_usd"),
			sql.As(t.C(entevent.FieldMaxConsecutive), "max_consecutive"),
			sql.As(t.C(entevent.FieldGoodID), "good_id"),
			sql.As(t.C(entevent.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entevent.FieldInviterLayers), "inviter_layers"),
			sql.As(t.C(entevent.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entevent.FieldCreatedAt), "created_at"),
			sql.As(t.C(entevent.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return wlog.WrapError(err)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) queryEventCoins(ctx context.Context) error {
	eventIDs := []uuid.UUID{}
	for _, info := range h.infos {
		id := uuid.MustParse(info.EntID)
		eventIDs = append(eventIDs, id)
	}
	handler, err := eventcoin1.NewHandler(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	handler.Limit = constant.DefaultRowLimit
	handler.Offset = 0
	handler.Conds = &eventcoincrud.Conds{
		EventIDs: &cruder.Cond{Op: cruder.IN, Val: eventIDs},
	}

	coins := []*eventcoinmw.EventCoin{}
	for {
		_coins, _, err := handler.GetEventCoins(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(_coins) == 0 {
			break
		}
		coins = append(coins, _coins...)
		handler.Offset += handler.Limit
	}

	coinMap := map[string][]*eventcoinmw.EventCoin{}
	for _, coin := range coins {
		itemKey := fmt.Sprintf("%v_%v", coin.EventID, coin.AppID)
		_coins, ok := coinMap[itemKey]
		if ok {
			coinMap[itemKey] = append(_coins, coin)
			continue
		}
		coinMap[itemKey] = []*eventcoinmw.EventCoin{coin}
	}
	for _, info := range h.infos {
		itemKey := fmt.Sprintf("%v_%v", info.EntID, info.AppID)
		_coins, ok := coinMap[itemKey]
		if ok {
			info.Coins = _coins
		}
	}
	return nil
}

func (h *queryHandler) queryEventCoupons(ctx context.Context) error {
	eventIDs := []uuid.UUID{}
	for _, info := range h.infos {
		id := uuid.MustParse(info.EntID)
		eventIDs = append(eventIDs, id)
	}
	handler, err := eventcoupon1.NewHandler(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	handler.Limit = constant.DefaultRowLimit
	handler.Offset = 0
	handler.Conds = &eventcouponcrud.Conds{
		EventIDs: &cruder.Cond{Op: cruder.IN, Val: eventIDs},
	}

	coupons := []*eventcouponmw.EventCoupon{}
	for {
		_coupons, _, err := handler.GetEventCoupons(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(_coupons) == 0 {
			break
		}
		coupons = append(coupons, _coupons...)
		handler.Offset += handler.Limit
	}

	couponMap := map[string][]string{}
	for _, coupon := range coupons {
		itemKey := fmt.Sprintf("%v_%v", coupon.EventID, coupon.AppID)
		_coupons, ok := couponMap[itemKey]
		if ok {
			couponMap[itemKey] = append(_coupons, coupon.CouponID)
			continue
		}
		couponMap[itemKey] = []string{coupon.CouponID}
	}
	for _, info := range h.infos {
		itemKey := fmt.Sprintf("%v_%v", info.EntID, info.AppID)
		_coupons, ok := couponMap[itemKey]
		if ok {
			info.CouponIDs = _coupons
		}
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.EventType = basetypes.UsedFor(basetypes.UsedFor_value[info.EventTypeStr])
		if info.GoodID != nil && *info.GoodID == uuid.Nil.String() {
			info.GoodID = nil
		}
		if info.AppGoodID != nil && *info.AppGoodID == uuid.Nil.String() {
			info.AppGoodID = nil
		}
		amount, err := decimal.NewFromString(info.Credits)
		if err != nil {
			info.Credits = decimal.NewFromInt(0).String()
		} else {
			info.Credits = amount.String()
		}
		amount, err = decimal.NewFromString(info.CreditsPerUSD)
		if err != nil {
			info.CreditsPerUSD = decimal.NewFromInt(0).String()
		} else {
			info.CreditsPerUSD = amount.String()
		}
	}
}

func (h *Handler) GetEvent(ctx context.Context) (*npool.Event, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Event{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEvent(cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryJoin(); err != nil {
			return wlog.WrapError(err)
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	if err := handler.queryEventCoins(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.queryEventCoupons(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetEvents(ctx context.Context) ([]*npool.Event, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Event{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryEvents(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryEvents(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryJoin(); err != nil {
			return wlog.WrapError(err)
		}
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)
		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	if err := handler.queryEventCoins(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.queryEventCoupons(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}

func (h *Handler) GetEventOnly(ctx context.Context) (*npool.Event, error) {
	const rowLimit = 2
	h.Limit = rowLimit
	infos, _, err := h.GetEvents(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos[0], nil
}
