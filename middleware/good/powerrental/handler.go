//nolint:dupl
package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	goodcoinrewardmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward"
	stockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/stock"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	goodcoincrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin"
	goodcoinrewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin/reward"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	goodrewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/reward"
	stockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock"
	mininggoodstockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock/mining"
	powerrentalcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/powerrental"
	device1 "github.com/NpoolPlatform/kunman/middleware/good/device"
	vendorlocation1 "github.com/NpoolPlatform/kunman/middleware/good/vender/location"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	powerrentalcrud.Req
	GoodBaseReq         *goodbasecrud.Req
	StockReq            *stockcrud.Req
	Rollback            *bool
	RewardReq           *goodrewardcrud.Req
	MiningGoodStockReqs []*mininggoodstockcrud.Req
	PowerRentalConds    *powerrentalcrud.Conds
	GoodBaseConds       *goodbasecrud.Conds
	GoodCoinConds       *goodcoincrud.Conds
	RewardConds         *goodrewardcrud.Conds
	CoinRewardReqs      []*goodcoinrewardcrud.Req
	Offset              int32
	Limit               int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		Req:              powerrentalcrud.Req{},
		GoodBaseReq:      &goodbasecrud.Req{},
		StockReq:         &stockcrud.Req{},
		RewardReq:        &goodrewardcrud.Req{},
		PowerRentalConds: &powerrentalcrud.Conds{},
		GoodBaseConds:    &goodbasecrud.Conds{},
		GoodCoinConds:    &goodcoincrud.Conds{},
		RewardConds:      &goodrewardcrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &id
		return nil
	}
}

func WithGoodID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = &id
		h.GoodBaseReq.EntID = &id
		h.StockReq.GoodID = &id
		h.RewardReq.GoodID = &id
		return nil
	}
}

func WithDeviceTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid devicetypeid")
			}
			return nil
		}
		handler, err := device1.NewHandler(
			ctx,
			device1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistDeviceType(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid devicetypeid")
		}
		h.DeviceTypeID = handler.EntID
		return nil
	}
}

func WithVendorLocationID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid vendorlocationid")
			}
			return nil
		}
		handler, err := vendorlocation1.NewHandler(
			ctx,
			vendorlocation1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistLocation(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid vendorlocation")
		}
		h.VendorLocationID = handler.EntID
		return nil
	}
}

func WithUnitPrice(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid unitprice")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return wlog.Errorf("invalid unitprice")
		}
		h.UnitPrice = &amount
		return nil
	}
}

func WithQuantityUnit(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid quantityunit")
			}
			return nil
		}
		const leastUnitLen = 2
		if len(*s) < leastUnitLen {
			return wlog.Errorf("invalid quantityunit")
		}
		h.QuantityUnit = s
		return nil
	}
}

func WithQuantityUnitAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid quantityunitamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return wlog.Errorf("invalid quantityunitamount")
		}
		h.QuantityUnitAmount = &amount
		return nil
	}
}

func WithDeliveryAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return wlog.Errorf("invalid deliveryat")
			}
			return nil
		}
		h.DeliveryAt = n
		return nil
	}
}

func WithUnitLockDeposit(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid unitlockdeposit")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid unitlockdeposit")
		}
		h.UnitLockDeposit = &amount
		return nil
	}
}

func WithDurationDisplayType(e *types.GoodDurationType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid durationdisplaytype")
			}
			return nil
		}
		switch *e {
		case types.GoodDurationType_GoodDurationByHour:
		case types.GoodDurationType_GoodDurationByDay:
		case types.GoodDurationType_GoodDurationByMonth:
		case types.GoodDurationType_GoodDurationByYear:
		default:
			return wlog.Errorf("invalid durationdisplaytype")
		}
		h.DurationDisplayType = e
		return nil
	}
}

func WithStockMode(e *types.GoodStockMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid stockmode")
			}
			return nil
		}
		switch *e {
		case types.GoodStockMode_GoodStockByMiningPool:
		case types.GoodStockMode_GoodStockByUnique:
		default:
			return wlog.Errorf("invalid stockmode")
		}
		h.StockMode = e
		return nil
	}
}

func WithGoodType(e *types.GoodType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid goodtype")
			}
			return nil
		}
		switch *e {
		case types.GoodType_PowerRental:
		case types.GoodType_LegacyPowerRental:
		default:
			return wlog.Errorf("invalid goodtype")
		}
		h.GoodBaseReq.GoodType = e
		return nil
	}
}

func WithBenefitType(e *types.BenefitType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid benefittype")
			}
			return nil
		}
		switch *e {
		case types.BenefitType_BenefitTypePlatform:
		case types.BenefitType_BenefitTypePool:
		case types.BenefitType_BenefitTypeOffline:
		default:
			return wlog.Errorf("invalid benefittype")
		}
		h.GoodBaseReq.BenefitType = e
		return nil
	}
}

func WithName(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid name")
			}
			return nil
		}
		const leastNameLen = 3
		if len(*s) < leastNameLen {
			return wlog.Errorf("invalid name")
		}
		h.GoodBaseReq.Name = s
		return nil
	}
}

func WithServiceStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.ServiceStartAt = n
		return nil
	}
}

func WithStartMode(e *types.GoodStartMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid startmode")
			}
			return nil
		}
		switch *e {
		case types.GoodStartMode_GoodStartModeTBD:
		case types.GoodStartMode_GoodStartModeConfirmed:
		case types.GoodStartMode_GoodStartModeNextDay:
		case types.GoodStartMode_GoodStartModeInstantly:
		case types.GoodStartMode_GoodStartModePreset:
		default:
			return wlog.Errorf("invalid startmode")
		}
		h.GoodBaseReq.StartMode = e
		return nil
	}
}

func WithTestOnly(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.TestOnly = b
		return nil
	}
}

func WithBenefitIntervalHours(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.BenefitIntervalHours = n
		return nil
	}
}

func WithPurchasable(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.Purchasable = b
		return nil
	}
}

func WithOnline(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.Online = b
		return nil
	}
}

func WithStockID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid stockid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.StockReq.EntID = &id
		return nil
	}
}

func WithTotal(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid total")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return wlog.Errorf("invalid total")
		}
		h.StockReq.Total = &amount
		return nil
	}
}

func WithStocks(stocks []*stockmwpb.MiningGoodStockReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, _stock := range stocks {
			entID := func() *uuid.UUID {
				uid, err := uuid.Parse(_stock.GetEntID())
				if err != nil {
					return nil
				}
				return &uid
			}()
			req := &mininggoodstockcrud.Req{
				EntID: entID,
				State: _stock.State,
			}

			if _stock.PoolRootUserID != nil {
				id, err := uuid.Parse(_stock.GetPoolRootUserID())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.PoolRootUserID = &id
			}

			if _stock.PoolGoodUserID != nil {
				id, err := uuid.Parse(_stock.GetPoolGoodUserID())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.PoolGoodUserID = &id
			}

			if _stock.Total != nil {
				amount, err := decimal.NewFromString(_stock.GetTotal())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.Total = &amount
			}

			h.MiningGoodStockReqs = append(h.MiningGoodStockReqs, req)
		}
		return nil
	}
}

func WithRewardState(e *types.BenefitState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid rewardstate")
			}
			return nil
		}
		switch *e {
		case types.BenefitState_BenefitWait:
		case types.BenefitState_BenefitTransferring:
		case types.BenefitState_BenefitBookKeeping:
		case types.BenefitState_BenefitUserBookKeeping:
		case types.BenefitState_BenefitSimulateBookKeeping:
		case types.BenefitState_BenefitDone:
		case types.BenefitState_BenefitFail:
		default:
			return wlog.Errorf("invalid rewardstate")
		}
		h.RewardReq.RewardState = e
		return nil
	}
}

func WithRewardAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return wlog.Errorf("invalid rewardat")
			}
			return nil
		}
		h.RewardReq.LastRewardAt = n
		return nil
	}
}

func WithRewards(rewards []*goodcoinrewardmwpb.RewardReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		coinTypeIDs := map[uuid.UUID]struct{}{}
		rewardTIDs := map[uuid.UUID]struct{}{}

		for _, reward := range rewards {
			_reward := &goodcoinrewardcrud.Req{}

			if reward.EntID != nil {
				_entID, err := uuid.Parse(*reward.EntID)
				if err != nil {
					return wlog.WrapError(err)
				}
				_reward.EntID = &_entID
			}

			_coinTypeID, err := uuid.Parse(reward.GetCoinTypeID())
			if err != nil {
				return wlog.WrapError(err)
			}
			if _, ok := coinTypeIDs[_coinTypeID]; ok {
				return wlog.Errorf("invalid cointypeid")
			}
			coinTypeIDs[_coinTypeID] = struct{}{}
			_reward.CoinTypeID = &_coinTypeID

			if reward.RewardTID != nil {
				_tid, err := uuid.Parse(*reward.RewardTID)
				if err != nil {
					return wlog.WrapError(err)
				}
				if _tid != uuid.Nil {
					if _, ok := rewardTIDs[_tid]; ok {
						return wlog.Errorf("invalid rewardid")
					}
					rewardTIDs[_tid] = struct{}{}
				}
				_reward.RewardTID = &_tid
			}

			if reward.RewardAmount != nil {
				_rewardAmount, err := decimal.NewFromString(*reward.RewardAmount)
				if err != nil {
					return wlog.WrapError(err)
				}
				_reward.LastRewardAmount = &_rewardAmount
			}

			if reward.NextRewardStartAmount != nil {
				_startAmount, err := decimal.NewFromString(*reward.NextRewardStartAmount)
				if err != nil {
					return wlog.WrapError(err)
				}
				_reward.NextRewardStartAmount = &_startAmount
			}

			h.CoinRewardReqs = append(h.CoinRewardReqs, _reward)
		}
		return nil
	}
}

func WithState(e *types.GoodState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid state")
			}
			return nil
		}
		switch *e {
		case types.GoodState_GoodStatePreWait:
		case types.GoodState_GoodStateWait:
		case types.GoodState_GoodStateCreateGoodUser:
		case types.GoodState_GoodStateCheckHashRate:
		case types.GoodState_GoodStateReady:
		case types.GoodState_GoodStateFail:
		default:
			return wlog.Errorf("invalid state")
		}
		h.GoodBaseReq.State = e
		return nil
	}
}

func WithRollback(e *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid rollback")
			}
			return nil
		}

		h.Rollback = e
		return nil
	}
}

func (h *Handler) withPowerRentalConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.PowerRentalConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PowerRentalConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PowerRentalConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.PowerRentalConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.StockMode != nil {
		h.PowerRentalConds.StockMode = &cruder.Cond{
			Op:  conds.GetStockMode().GetOp(),
			Val: types.GoodStockMode(conds.GetStockMode().GetValue()),
		}
	}
	return nil
}

func (h *Handler) withGoodBaseConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodBaseConds.EntID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.GoodBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.GoodType != nil {
		h.GoodBaseConds.GoodType = &cruder.Cond{
			Op:  conds.GetGoodType().GetOp(),
			Val: types.GoodType(conds.GetGoodType().GetValue()),
		}
	}
	if conds.GoodTypes != nil {
		es := []types.GoodType{}
		for _, e := range conds.GetGoodTypes().GetValue() {
			es = append(es, types.GoodType(e))
		}
		h.GoodBaseConds.GoodTypes = &cruder.Cond{
			Op:  conds.GetGoodTypes().GetOp(),
			Val: es,
		}
	}
	if conds.State != nil {
		h.GoodBaseConds.State = &cruder.Cond{
			Op:  conds.GetState().GetOp(),
			Val: types.GoodState(conds.GetState().GetValue()),
		}
	}
	return nil
}

func (h *Handler) withGoodCoinConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodCoinConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.GoodCoinConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.CoinTypeID != nil {
		id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodCoinConds.CoinTypeID = &cruder.Cond{
			Op:  conds.GetCoinTypeID().GetOp(),
			Val: id,
		}
	}
	if conds.CoinTypeIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetCoinTypeIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.GoodCoinConds.CoinTypeIDs = &cruder.Cond{
			Op:  conds.GetCoinTypeIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withRewardConds(conds *npool.Conds) error {
	if conds.RewardState != nil {
		h.RewardConds.RewardState = &cruder.Cond{
			Op:  conds.GetRewardState().GetOp(),
			Val: types.BenefitState(conds.GetRewardState().GetValue()),
		}
	}
	if conds.RewardAt != nil {
		h.RewardConds.RewardAt = &cruder.Cond{
			Op:  conds.GetRewardAt().GetOp(),
			Val: conds.GetRewardAt().GetValue(),
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.RewardConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withPowerRentalConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withGoodCoinConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withRewardConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withGoodBaseConds(conds)
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
