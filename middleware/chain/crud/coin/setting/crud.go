package setting

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entsetting "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/setting"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID                       *uuid.UUID
	CoinTypeID                  *uuid.UUID
	FeeCoinTypeID               *uuid.UUID
	WithdrawFeeByStableUSD      *bool
	WithdrawFeeAmount           *decimal.Decimal
	CollectFeeAmount            *decimal.Decimal
	HotWalletFeeAmount          *decimal.Decimal
	LowFeeAmount                *decimal.Decimal
	HotLowFeeAmount             *decimal.Decimal
	HotWalletAccountAmount      *decimal.Decimal
	PaymentAccountCollectAmount *decimal.Decimal
	LeastTransferAmount         *decimal.Decimal
	NeedMemo                    *bool
	RefreshCurrency             *bool
	CheckNewAddressBalance      *bool
	DeletedAt                   *uint32
}

func CreateSet(c *ent.SettingCreate, req *Req) *ent.SettingCreate { //nolint
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FeeCoinTypeID != nil {
		c.SetFeeCoinTypeID(*req.FeeCoinTypeID)
	}
	if req.WithdrawFeeByStableUSD != nil {
		c.SetWithdrawFeeByStableUsd(*req.WithdrawFeeByStableUSD)
	}
	if req.WithdrawFeeAmount != nil {
		c.SetWithdrawFeeAmount(*req.WithdrawFeeAmount)
	}
	if req.CollectFeeAmount != nil {
		c.SetCollectFeeAmount(*req.CollectFeeAmount)
	}
	if req.HotWalletFeeAmount != nil {
		c.SetHotWalletFeeAmount(*req.HotWalletFeeAmount)
	}
	if req.LowFeeAmount != nil {
		c.SetLowFeeAmount(*req.LowFeeAmount)
	}
	if req.HotLowFeeAmount != nil {
		c.SetHotLowFeeAmount(*req.HotLowFeeAmount)
	}
	if req.HotWalletAccountAmount != nil {
		c.SetHotWalletAccountAmount(*req.HotWalletAccountAmount)
	}
	if req.PaymentAccountCollectAmount != nil {
		c.SetPaymentAccountCollectAmount(*req.PaymentAccountCollectAmount)
	}
	if req.LeastTransferAmount != nil {
		c.SetLeastTransferAmount(*req.LeastTransferAmount)
	}
	if req.NeedMemo != nil {
		c.SetNeedMemo(*req.NeedMemo)
	}
	if req.RefreshCurrency != nil {
		c.SetRefreshCurrency(*req.RefreshCurrency)
	}
	if req.CheckNewAddressBalance != nil {
		c.SetCheckNewAddressBalance(*req.CheckNewAddressBalance)
	}
	return c
}

func UpdateSet(u *ent.SettingUpdateOne, req *Req) *ent.SettingUpdateOne {
	if req.FeeCoinTypeID != nil {
		u.SetFeeCoinTypeID(*req.FeeCoinTypeID)
	}
	if req.WithdrawFeeByStableUSD != nil {
		u.SetWithdrawFeeByStableUsd(*req.WithdrawFeeByStableUSD)
	}
	if req.WithdrawFeeAmount != nil {
		u.SetWithdrawFeeAmount(*req.WithdrawFeeAmount)
	}
	if req.CollectFeeAmount != nil {
		u.SetCollectFeeAmount(*req.CollectFeeAmount)
	}
	if req.HotWalletFeeAmount != nil {
		u.SetHotWalletFeeAmount(*req.HotWalletFeeAmount)
	}
	if req.LowFeeAmount != nil {
		u.SetLowFeeAmount(*req.LowFeeAmount)
	}
	if req.HotLowFeeAmount != nil {
		u.SetHotLowFeeAmount(*req.HotLowFeeAmount)
	}
	if req.HotWalletAccountAmount != nil {
		u.SetHotWalletAccountAmount(*req.HotWalletAccountAmount)
	}
	if req.PaymentAccountCollectAmount != nil {
		u.SetPaymentAccountCollectAmount(*req.PaymentAccountCollectAmount)
	}
	if req.LeastTransferAmount != nil {
		u.SetLeastTransferAmount(*req.LeastTransferAmount)
	}
	if req.NeedMemo != nil {
		u.SetNeedMemo(*req.NeedMemo)
	}
	if req.RefreshCurrency != nil {
		u.SetRefreshCurrency(*req.RefreshCurrency)
	}
	if req.CheckNewAddressBalance != nil {
		u.SetCheckNewAddressBalance(*req.CheckNewAddressBalance)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}

	return u
}

type Conds struct {
	EntID         *cruder.Cond
	CoinTypeID    *cruder.Cond
	FeeCoinTypeID *cruder.Cond
}

func SetQueryConds(q *ent.SettingQuery, conds *Conds) (*ent.SettingQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entsetting.EntID(id))
		default:
			return nil, fmt.Errorf("invalid setting field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entsetting.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid setting field")
		}
	}
	if conds.FeeCoinTypeID != nil {
		id, ok := conds.FeeCoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid feecointypeid")
		}
		switch conds.FeeCoinTypeID.Op {
		case cruder.EQ:
			q.Where(entsetting.FeeCoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid setting field")
		}
	}
	q.Where(entsetting.DeletedAt(0))
	return q, nil
}
