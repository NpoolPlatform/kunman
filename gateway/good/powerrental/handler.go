//nolint:dupl
package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	devicecommon "github.com/NpoolPlatform/kunman/gateway/good/device/common"
	goodcommon "github.com/NpoolPlatform/kunman/gateway/good/good/common"
	locationcommon "github.com/NpoolPlatform/kunman/gateway/good/vender/location/common"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	goodstockgwpb "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/stock"
	goodstockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/stock"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID    *uint32
	EntID *string
	goodcommon.GoodCheckHandler
	devicecommon.DeviceTypeCheckHandler
	locationcommon.LocationCheckHandler
	DeviceTypeID         *string
	VendorLocationID     *string
	UnitPrice            *string
	QuantityUnit         *string
	QuantityUnitAmount   *string
	DeliveryAt           *uint32
	UnitLockDeposit      *string
	DurationDisplayType  *types.GoodDurationType
	GoodType             *types.GoodType
	Name                 *string
	ServiceStartAt       *uint32
	StartMode            *types.GoodStartMode
	TestOnly             *bool
	BenefitIntervalHours *uint32
	Purchasable          *bool
	Online               *bool
	StockMode            *types.GoodStockMode
	Total                *string
	MiningGoodStocks     []*goodstockmwpb.MiningGoodStockReq
	State                *types.GoodState
	Offset               int32
	Limit                int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
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

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.EntID = id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.GoodID = id
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
		if err := h.CheckDeviceTypeWithDeviceTypeID(ctx, *id); err != nil {
			return wlog.WrapError(err)
		}
		h.DeviceTypeID = id
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
		if err := h.CheckLocationWithLocationID(ctx, *id); err != nil {
			return wlog.WrapError(err)
		}
		h.VendorLocationID = id
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
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return wlog.Errorf("invalid unitprice")
		}
		h.UnitPrice = s
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
		if len(*s) < 2 {
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
		h.QuantityUnitAmount = s
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
		h.UnitLockDeposit = s
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
		h.GoodType = e
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
		h.Name = s
		return nil
	}
}

func WithServiceStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ServiceStartAt = n
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
		h.StartMode = e
		return nil
	}
}

func WithTestOnly(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.TestOnly = b
		return nil
	}
}

func WithBenefitIntervalHours(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.BenefitIntervalHours = n
		return nil
	}
}

func WithPurchasable(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Purchasable = b
		return nil
	}
}

func WithOnline(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Online = b
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
		h.State = e
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
		h.Total = s
		return nil
	}
}

func WithMiningGoodStocks(stocks []*goodstockgwpb.MiningGoodStockReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		miningGoodStocks := []*goodstockmwpb.MiningGoodStockReq{}

		ids := []string{}
		for _, req := range stocks {
			if req.PoolRootUserID != nil {
				ids = append(ids, *req.PoolRootUserID)
			}
			miningGoodStocks = append(miningGoodStocks, &goodstockmwpb.MiningGoodStockReq{
				EntID:          req.EntID,
				PoolRootUserID: req.PoolRootUserID,
				State:          req.State,
				Total:          req.Total,
			})
		}

		handle := &checkHandler{h}
		if err := handle.checkPoolRootUserIDs(ctx, ids); err != nil {
			return wlog.WrapError(err)
		}

		h.MiningGoodStocks = miningGoodStocks
		return nil
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
