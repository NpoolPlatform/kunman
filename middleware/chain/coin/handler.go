package coin

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	coincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                          *uint32
	EntID                       *uuid.UUID
	Name                        *string
	Logo                        *string
	Presale                     *bool
	Unit                        *string
	ENV                         *string
	ReservedAmount              *decimal.Decimal
	ForPay                      *bool
	HomePage                    *string
	Specs                       *string
	FeeCoinTypeID               *uuid.UUID
	WithdrawFeeByStableUSD      *bool
	WithdrawFeeAmount           *decimal.Decimal
	CollectFeeAmount            *decimal.Decimal
	HotWalletFeeAmount          *decimal.Decimal
	LowFeeAmount                *decimal.Decimal
	HotLowFeeAmount             *decimal.Decimal
	HotWalletAccountAmount      *decimal.Decimal
	PaymentAccountCollectAmount *decimal.Decimal
	Disabled                    *bool
	StableUSD                   *bool
	LeastTransferAmount         *decimal.Decimal
	NeedMemo                    *bool
	RefreshCurrency             *bool
	CheckNewAddressBalance      *bool
	ChainType                   *string
	ChainNativeUnit             *string
	ChainAtomicUnit             *string
	ChainUnitExp                *uint32
	GasType                     *basetypes.GasType
	ChainID                     *string
	ChainNickname               *string
	ChainNativeCoinName         *string
	Conds                       *coincrud.Conds
	Offset                      int32
	Limit                       int32
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

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid coinname")
		}
		h.Name = name
		return nil
	}
}

func WithLogo(logo *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Logo = logo
		return nil
	}
}

func WithPresale(presale *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Presale = presale
		return nil
	}
}

func WithUnit(unit *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if unit == nil {
			if must {
				return fmt.Errorf("invalid unit")
			}
			return nil
		}
		if *unit == "" {
			return fmt.Errorf("invalid coinunit")
		}
		h.Unit = unit
		return nil
	}
}

func WithENV(env *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if env == nil {
			if must {
				return fmt.Errorf("invalid ent")
			}
			return nil
		}
		switch *env {
		case "main":
		case "test":
		case "local":
		default:
			return fmt.Errorf("invalid coinenv")
		}
		h.ENV = env
		return nil
	}
}

func WithReservedAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.ReservedAmount = &_amount
		return nil
	}
}

func WithForPay(forPay *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ForPay = forPay
		return nil
	}
}

func WithHomePage(homePage *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.HomePage = homePage
		return nil
	}
}

func WithSpecs(specs *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Specs = specs
		return nil
	}
}

func WithFeeCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid feecointypeid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.FeeCoinTypeID = &_id
		return nil
	}
}

func WithWithdrawFeeByStableUSD(stable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.WithdrawFeeByStableUSD = stable
		return nil
	}
}

func WithWithdrawFeeAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid withdrawfeeamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.WithdrawFeeAmount = &_amount
		return nil
	}
}

func WithCollectFeeAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid collectfeeamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.CollectFeeAmount = &_amount
		return nil
	}
}

func WithHotWalletFeeAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid hotwalletfeeamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.HotWalletFeeAmount = &_amount
		return nil
	}
}

func WithLowFeeAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid lowfeeamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.LowFeeAmount = &_amount
		return nil
	}
}

func WithHotLowFeeAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid hotlowfeeamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.HotLowFeeAmount = &_amount
		return nil
	}
}

func WithHotWalletAccountAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid hotwalletaccountamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.HotWalletAccountAmount = &_amount
		return nil
	}
}

func WithPaymentAccountCollectAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid paymentaccountcollectamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.PaymentAccountCollectAmount = &_amount
		return nil
	}
}

func WithDisabled(disabled *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Disabled = disabled
		return nil
	}
}

func WithStableUSD(stable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.StableUSD = stable
		return nil
	}
}

func WithLeastTransferAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid leasttransferamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.LeastTransferAmount = &_amount
		return nil
	}
}

func WithNeedMemo(needMemo *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.NeedMemo = needMemo
		return nil
	}
}

func WithRefreshCurrency(refresh *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RefreshCurrency = refresh
		return nil
	}
}

func WithCheckNewAddressBalance(check *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.CheckNewAddressBalance = check
		return nil
	}
}

func WithChainType(chainType *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if chainType == nil {
			if must {
				return fmt.Errorf("invalid chaintype")
			}
			return nil
		}
		if *chainType == "" {
			return fmt.Errorf("invalid chaintype")
		}
		h.ChainType = chainType
		return nil
	}
}

func WithChainNativeUnit(unit *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if unit == nil {
			if must {
				return fmt.Errorf("invalid chainnativeunit")
			}
			return nil
		}
		if *unit == "" {
			return fmt.Errorf("invalid nativeunit")
		}
		h.ChainNativeUnit = unit
		return nil
	}
}

func WithChainAtomicUnit(unit *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if unit == nil {
			if must {
				return fmt.Errorf("invalid chainatomicunit")
			}
			return nil
		}
		if *unit == "" {
			return fmt.Errorf("invalid atomicunit")
		}
		h.ChainAtomicUnit = unit
		return nil
	}
}

func WithChainUnitExp(exp *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ChainUnitExp = exp
		return nil
	}
}

func WithGasType(gasType *basetypes.GasType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if gasType == nil {
			if must {
				return fmt.Errorf("invalid gastype")
			}
			return nil
		}
		switch *gasType {
		case basetypes.GasType_FixedGas:
		case basetypes.GasType_DynamicGas:
		case basetypes.GasType_GasUnsupported:
		default:
			return fmt.Errorf("invalid gastype")
		}
		h.GasType = gasType
		return nil
	}
}

func WithChainID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid chainid")
			}
			return nil
		}
		if *id == "" {
			return fmt.Errorf("invalid chainid")
		}
		h.ChainID = id
		return nil
	}
}

func WithChainNickname(nickname *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if nickname == nil {
			if must {
				return fmt.Errorf("invalid chainnickname")
			}
			return nil
		}
		if *nickname == "" {
			return fmt.Errorf("invalid nickname")
		}
		h.ChainNickname = nickname
		return nil
	}
}

func WithChainNativeCoinName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid chainnativecoinname")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid chainnativecoinname")
		}
		h.ChainNativeCoinName = name
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &coincrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.Presale != nil {
			h.Conds.Presale = &cruder.Cond{
				Op:  conds.GetPresale().GetOp(),
				Val: conds.GetPresale().GetValue(),
			}
		}
		if conds.ENV != nil {
			h.Conds.ENV = &cruder.Cond{
				Op:  conds.GetENV().GetOp(),
				Val: conds.GetENV().GetValue(),
			}
		}
		if conds.ForPay != nil {
			h.Conds.ForPay = &cruder.Cond{
				Op:  conds.GetForPay().GetOp(),
				Val: conds.GetForPay().GetValue(),
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.Disabled != nil {
			h.Conds.Disabled = &cruder.Cond{
				Op:  conds.GetDisabled().GetOp(),
				Val: conds.GetDisabled().GetValue(),
			}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
			}
		}
		if conds.Names != nil {
			h.Conds.Names = &cruder.Cond{
				Op:  conds.GetNames().GetOp(),
				Val: conds.GetNames().GetValue(),
			}
		}
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
