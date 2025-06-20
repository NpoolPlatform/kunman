package subscription

import (
	"context"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	paymentgwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/payment"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	common "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/google/uuid"
)

type Handler struct {
	ID    *uint32
	EntID *string
	ordercommon.OrderCheckHandler
	common.CoinCheckHandler
	common.FiatCheckHandler
	common.AllocatedCouponCheckHandler
	common.AppGoodCheckHandler
	Balances                  []*paymentmwpb.PaymentBalanceReq
	PaymentTransferCoinTypeID *string
	PaymentFiatID             *string
	FiatPaymentChannel        *types.FiatPaymentChannel
	FiatChannelPaymentID      *string
	CouponIDs                 []string
	OrderIDs                  []string
	LifeSeconds               *uint32
	UserSetCanceled           *bool
	AdminSetCanceled          *bool
	CreateMethod              *types.OrderCreateMethod
	OrderType                 *types.OrderType
	Offset                    int32
	Limit                     int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
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

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		if err := h.OrderCheckHandler.CheckAppWithAppID(ctx, *id); err != nil {
			return wlog.WrapError(err)
		}
		h.OrderCheckHandler.AppID = id
		h.AppGoodCheckHandler.AppID = id
		h.AllocatedCouponCheckHandler.AppID = id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid userid")
			}
			return nil
		}
		if err := h.OrderCheckHandler.CheckUserWithUserID(ctx, *id); err != nil {
			return wlog.WrapError(err)
		}
		h.OrderCheckHandler.UserID = id
		h.AppGoodCheckHandler.UserID = id
		h.AllocatedCouponCheckHandler.UserID = id
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
		if err := h.OrderCheckHandler.CheckGoodWithGoodID(ctx, *id); err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = id
		h.OrderCheckHandler.GoodID = id
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appgoodid")
			}
			return nil
		}
		if err := h.OrderCheckHandler.CheckAppGoodWithAppGoodID(ctx, *id); err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodID = id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderid")
			}
			return nil
		}
		_, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = id
		return nil
	}
}

func WithPaymentBalances(bs []*paymentgwpb.PaymentBalance, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, balance := range bs {
			if err := h.CheckCoinWithCoinTypeID(ctx, balance.CoinTypeID); err != nil {
				return wlog.WrapError(err)
			}
			// Fill coin_usd_currency later
			h.Balances = append(h.Balances, &paymentmwpb.PaymentBalanceReq{
				CoinTypeID: &balance.CoinTypeID,
				Amount:     &balance.Amount,
			})
		}
		logger.Sugar().Infow("WithPaymentBalances", "Balances", h.Balances)
		return nil
	}
}

func WithPaymentTransferCoinTypeID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid paymenttransfercointypeid")
			}
			return nil
		}
		if err := h.CheckCoinWithCoinTypeID(ctx, *s); err != nil {
			return wlog.WrapError(err)
		}
		h.PaymentTransferCoinTypeID = s
		return nil
	}
}

func WithPaymentFiatID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid fiatid")
			}
			return nil
		}
		if err := h.CheckFiatWithFiatID(ctx, *s); err != nil {
			return wlog.WrapError(err)
		}
		h.PaymentFiatID = s
		return nil
	}
}

func WithFiatPaymentChannel(e *types.FiatPaymentChannel, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid paymentchannel")
			}
			return nil
		}
		h.FiatPaymentChannel = e
		return nil
	}
}

func WithFiatChannelPaymentID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid channelpaymentid")
			}
			return nil
		}
		h.FiatChannelPaymentID = s
		return nil
	}
}

func WithCouponIDs(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, couponID := range ss {
			if err := h.CheckAllocatedCouponWithAllocatedCouponID(ctx, couponID); err != nil {
				return wlog.WrapError(err)
			}
		}
		h.CouponIDs = ss
		return nil
	}
}

func WithUserSetCanceled(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.UserSetCanceled = b
		return nil
	}
}

func WithAdminSetCanceled(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AdminSetCanceled = b
		return nil
	}
}

func WithCreateMethod(e *types.OrderCreateMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid createmethod")
			}
			return nil
		}
		switch *e {
		case types.OrderCreateMethod_OrderCreatedByPurchase:
		case types.OrderCreateMethod_OrderCreatedByAdmin:
		default:
			return wlog.Errorf("invalid createmethod")
		}
		h.CreateMethod = e
		return nil
	}
}

func WithOrderType(orderType *types.OrderType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if orderType == nil {
			if must {
				return wlog.Errorf("invalid ordertype")
			}
			return nil
		}
		switch *orderType {
		case types.OrderType_Airdrop:
		case types.OrderType_Normal:
		default:
			return wlog.Errorf("invalid ordertype")
		}
		h.OrderType = orderType
		return nil
	}
}

func WithLifeSeconds(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.LifeSeconds = u
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
