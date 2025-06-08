package coin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"

	chainbasecrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/chain"
	basecrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin"
	extracrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/extra"
	settingcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/setting"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createChainBase(ctx context.Context, tx *ent.Tx) error {
	if h.ChainType == nil {
		return nil
	}
	conds := &chainbasecrud.Conds{
		Name: &cruder.Cond{Op: cruder.EQ, Val: *h.ChainType},
	}
	if h.ENV != nil {
		conds.ENV = &cruder.Cond{Op: cruder.EQ, Val: *h.ENV}
	}
	if h.ChainID != nil {
		conds.ChainID = &cruder.Cond{Op: cruder.EQ, Val: *h.ChainID}
	}

	stm, err := chainbasecrud.SetQueryConds(tx.ChainBase.Query(), conds)
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info != nil {
		return nil
	}

	if _, err := chainbasecrud.CreateSet(
		tx.ChainBase.Create(),
		&chainbasecrud.Req{
			Name:       h.ChainType,
			NativeUnit: h.ChainNativeUnit,
			AtomicUnit: h.ChainAtomicUnit,
			UnitExp:    h.ChainUnitExp,
			ENV:        h.ENV,
			ChainID:    h.ChainID,
			Nickname:   h.ChainNickname,
			GasType:    h.GasType,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createNativeCoinBase(ctx context.Context, tx *ent.Tx) error {
	if h.ChainNativeCoinName == nil {
		h.FeeCoinTypeID = h.EntID
		return nil
	}
	if *h.ChainNativeUnit == *h.Unit {
		h.FeeCoinTypeID = h.EntID
		return nil
	}

	stm, err := basecrud.SetQueryConds(
		tx.CoinBase.Query(),
		&basecrud.Conds{
			Name: &cruder.Cond{Op: cruder.EQ, Val: *h.ChainNativeCoinName},
			ENV:  &cruder.Cond{Op: cruder.EQ, Val: *h.ENV},
		},
	)
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if info == nil {
		info, err = basecrud.CreateSet(
			tx.CoinBase.Create(),
			&basecrud.Req{
				Name: h.ChainNativeCoinName,
				Unit: h.ChainNativeUnit,
				ENV:  h.ENV,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
	}

	h.FeeCoinTypeID = &info.EntID

	return nil
}
func (h *createHandler) createCoinBase(ctx context.Context, tx *ent.Tx) error {
	stm, err := basecrud.SetQueryConds(
		tx.CoinBase.Query(),
		&basecrud.Conds{
			Name: &cruder.Cond{Op: cruder.EQ, Val: *h.Name},
			ENV:  &cruder.Cond{Op: cruder.EQ, Val: *h.ENV},
		},
	)
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info != nil {
		h.EntID = &info.EntID
		return nil
	}

	if _, err := basecrud.CreateSet(
		tx.CoinBase.Create(),
		&basecrud.Req{
			EntID:          h.EntID,
			Name:           h.Name,
			Logo:           h.Logo,
			Presale:        h.Presale,
			Unit:           h.Unit,
			ENV:            h.ENV,
			ReservedAmount: h.ReservedAmount,
			ForPay:         h.ForPay,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createNativeCoinExtra(ctx context.Context, tx *ent.Tx) error {
	if h.FeeCoinTypeID == nil {
		return nil
	}

	stm, err := extracrud.SetQueryConds(
		tx.CoinExtra.Query(),
		&extracrud.Conds{
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.FeeCoinTypeID},
		},
	)
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info != nil {
		return nil
	}

	if _, err := extracrud.CreateSet(
		tx.CoinExtra.Create(),
		&extracrud.Req{
			CoinTypeID: h.FeeCoinTypeID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createCoinExtra(ctx context.Context, tx *ent.Tx) error {
	stm, err := extracrud.SetQueryConds(
		tx.CoinExtra.Query(),
		&extracrud.Conds{
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.EntID},
		},
	)
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info != nil {
		return nil
	}

	if _, err := extracrud.CreateSet(
		tx.CoinExtra.Create(),
		&extracrud.Req{
			CoinTypeID: h.EntID,
			HomePage:   h.HomePage,
			Specs:      h.Specs,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createNativeCoinSetting(ctx context.Context, tx *ent.Tx) error {
	if h.FeeCoinTypeID == nil {
		return nil
	}

	stm, err := settingcrud.SetQueryConds(
		tx.Setting.Query(),
		&settingcrud.Conds{
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.FeeCoinTypeID},
		},
	)
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info != nil {
		return nil
	}

	if _, err := settingcrud.CreateSet(
		tx.Setting.Create(),
		&settingcrud.Req{
			CoinTypeID:    h.FeeCoinTypeID,
			FeeCoinTypeID: h.FeeCoinTypeID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createCoinSetting(ctx context.Context, tx *ent.Tx) error {
	stm, err := settingcrud.SetQueryConds(
		tx.Setting.Query(),
		&settingcrud.Conds{
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.EntID},
		},
	)
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info != nil {
		return nil
	}

	if _, err := settingcrud.CreateSet(
		tx.Setting.Create(),
		&settingcrud.Req{
			CoinTypeID:                  h.EntID,
			FeeCoinTypeID:               h.FeeCoinTypeID,
			WithdrawFeeByStableUSD:      h.WithdrawFeeByStableUSD,
			WithdrawFeeAmount:           h.WithdrawFeeAmount,
			CollectFeeAmount:            h.CollectFeeAmount,
			HotWalletFeeAmount:          h.HotWalletFeeAmount,
			LowFeeAmount:                h.LowFeeAmount,
			HotLowFeeAmount:             h.HotLowFeeAmount,
			HotWalletAccountAmount:      h.HotWalletAccountAmount,
			PaymentAccountCollectAmount: h.PaymentAccountCollectAmount,
			LeastTransferAmount:         h.LeastTransferAmount,
			NeedMemo:                    h.NeedMemo,
			RefreshCurrency:             h.RefreshCurrency,
			CheckNewAddressBalance:      h.CheckNewAddressBalance,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCoin(ctx context.Context) (*npool.Coin, error) {
	if h.Name == nil {
		return nil, fmt.Errorf("invalid coinname")
	}
	if h.Unit == nil {
		return nil, fmt.Errorf("invalid coinunit")
	}
	if h.ENV == nil {
		return nil, fmt.Errorf("invalid coinenv")
	}

	// TODO: deduplicate

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createChainBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createNativeCoinBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createNativeCoinExtra(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createNativeCoinSetting(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createCoinBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createCoinExtra(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createCoinSetting(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoin(ctx)
}
