package appcoin

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	appcoincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin"

	entappcoin "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/appcoin"
	entcoinbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	entcoinextra "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinextra"
	entappexrate "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/exchangerate"
	entsetting "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/setting"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppCoinSelect
	infos []*npool.Coin
	total uint32
}

func (h *queryHandler) selectAppCoin(stm *ent.AppCoinQuery) {
	h.stm = stm.Select(
		entappcoin.FieldID,
		entappcoin.FieldEntID,
		entappcoin.FieldAppID,
		entappcoin.FieldCoinTypeID,
		entappcoin.FieldName,
		entappcoin.FieldDisplayNames,
		entappcoin.FieldLogo,
		entappcoin.FieldForPay,
		entappcoin.FieldWithdrawAutoReviewAmount,
		entappcoin.FieldProductPage,
		entappcoin.FieldDisabled,
		entappcoin.FieldCreatedAt,
		entappcoin.FieldUpdatedAt,
		entappcoin.FieldDisplay,
		entappcoin.FieldDailyRewardAmount,
		entappcoin.FieldDisplayIndex,
		entappcoin.FieldMaxAmountPerWithdraw,
	)
}

func (h *queryHandler) queryAppCoin(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppCoin.Query().Where(entappcoin.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappcoin.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappcoin.EntID(*h.EntID))
	}
	h.selectAppCoin(stm)
	return nil
}

func (h *queryHandler) queryAppCoins(ctx context.Context, cli *ent.Client) error {
	stm, err := appcoincrud.SetQueryConds(cli.AppCoin.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectAppCoin(stm)
	return nil
}

func (h *queryHandler) queryJoinCoinBase(s *sql.Selector) {
	t := sql.Table(entcoinbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entappcoin.FieldCoinTypeID),
			t.C(entcoinbase.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entcoinbase.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entcoinbase.FieldName), "coin_name"),
			sql.As(t.C(entcoinbase.FieldUnit), "unit"),
			sql.As(t.C(entcoinbase.FieldEnv), "env"),
			sql.As(t.C(entcoinbase.FieldPresale), "presale"),
			sql.As(t.C(entcoinbase.FieldReservedAmount), "reserved_amount"),
			sql.As(t.C(entcoinbase.FieldDisabled), "coin_disabled"),
			sql.As(t.C(entcoinbase.FieldForPay), "coin_for_pay"),
		)
}

func (h *queryHandler) queryJoinCoinExtra(s *sql.Selector) {
	t := sql.Table(entcoinextra.Table)
	s.LeftJoin(t).
		On(
			s.C(entappcoin.FieldCoinTypeID),
			t.C(entcoinextra.FieldCoinTypeID),
		).
		OnP(
			sql.EQ(t.C(entcoinextra.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entcoinextra.FieldHomePage), "home_page"),
			sql.As(t.C(entcoinextra.FieldSpecs), "specs"),
			sql.As(t.C(entcoinextra.FieldStableUsd), "stable_usd"),
		)
}

func (h *queryHandler) queryJoinCoinSetting(s *sql.Selector) {
	t1 := sql.Table(entsetting.Table)
	s.
		LeftJoin(t1).
		On(
			s.C(entappcoin.FieldCoinTypeID),
			t1.C(entsetting.FieldCoinTypeID),
		).
		OnP(
			sql.EQ(t1.C(entsetting.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t1.C(entsetting.FieldFeeCoinTypeID), "fee_coin_type_id"),
			sql.As(t1.C(entsetting.FieldWithdrawFeeByStableUsd), "withdraw_fee_by_stable_usd"),
			sql.As(t1.C(entsetting.FieldWithdrawFeeAmount), "withdraw_fee_amount"),
			sql.As(t1.C(entsetting.FieldCollectFeeAmount), "collect_fee_amount"),
			sql.As(t1.C(entsetting.FieldHotWalletFeeAmount), "hot_wallet_fee_amount"),
			sql.As(t1.C(entsetting.FieldLowFeeAmount), "low_fee_amount"),
			sql.As(t1.C(entsetting.FieldHotLowFeeAmount), "hot_low_fee_amount"),
			sql.As(t1.C(entsetting.FieldHotWalletAccountAmount), "hot_wallet_account_amount"),
			sql.As(t1.C(entsetting.FieldPaymentAccountCollectAmount), "payment_account_collect_amount"),
			sql.As(t1.C(entsetting.FieldLeastTransferAmount), "least_transfer_amount"),
			sql.As(t1.C(entsetting.FieldNeedMemo), "need_memo"),
			sql.As(t1.C(entsetting.FieldCheckNewAddressBalance), "check_new_address_balance"),
		)

	t2 := sql.Table(entcoinbase.Table)
	s.
		LeftJoin(t2).
		On(
			t1.C(entsetting.FieldFeeCoinTypeID),
			t2.C(entcoinbase.FieldEntID),
		).
		OnP(
			sql.EQ(t2.C(entcoinbase.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t2.C(entcoinbase.FieldName), "fee_coin_name"),
			sql.As(t2.C(entcoinbase.FieldLogo), "fee_coin_logo"),
			sql.As(t2.C(entcoinbase.FieldUnit), "fee_coin_unit"),
			sql.As(t2.C(entcoinbase.FieldEnv), "fee_coin_env"),
		)
}

func (h *queryHandler) queryJoinExrate(s *sql.Selector) {
	t := sql.Table(entappexrate.Table)
	s.LeftJoin(t).
		On(
			s.C(entappcoin.FieldCoinTypeID),
			t.C(entappexrate.FieldCoinTypeID),
		).
		On(
			s.C(entappcoin.FieldAppID),
			t.C(entappexrate.FieldAppID),
		).
		OnP(
			sql.EQ(t.C(entappexrate.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entappexrate.FieldMarketValue), "market_value"),
			sql.As(t.C(entappexrate.FieldSettleValue), "settle_value"),
			sql.As(t.C(entappexrate.FieldSettlePercent), "settle_percent"),
			sql.As(t.C(entappexrate.FieldSettleTips), "settle_tips"),
			sql.As(t.C(entappexrate.FieldSetter), "setter"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinCoinBase(s)
		h.queryJoinCoinExtra(s)
		h.queryJoinCoinSetting(s)
		h.queryJoinExrate(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		_ = json.Unmarshal([]byte(info.DisplayNamesStr), &info.DisplayNames)
		if !info.CoinForPay {
			info.ForPay = info.CoinForPay
		}
		if !info.Disabled {
			info.Disabled = info.CoinDisabled
		}
		if info.MarketValue == "" {
			info.MarketValue = decimal.NewFromInt(0).String()
		}
		if info.SettleValue == "" {
			info.SettleValue = decimal.NewFromInt(0).String()
		}
		if info.MaxAmountPerWithdraw == "" {
			info.MaxAmountPerWithdraw = decimal.NewFromInt(0).String()
		}
		_ = json.Unmarshal([]byte(info.SettleTipsStr), &info.SettleTips)
	}
}

func (h *Handler) GetCoin(ctx context.Context) (*npool.Coin, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppCoin(cli); err != nil {
			return err
		}
		handler.queryJoin()
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetCoins(ctx context.Context) ([]*npool.Coin, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppCoins(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Order(ent.Asc(entappcoin.FieldDisplayIndex)).
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}

func (h *Handler) GetCoinOnly(ctx context.Context) (*npool.Coin, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppCoins(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Order(ent.Asc(entappcoin.FieldDisplayIndex)).
			Offset(0).
			Limit(2)
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("invalid coin")
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()
	return handler.infos[0], nil
}
