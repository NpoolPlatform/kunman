package coin

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	coincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin"
	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	entbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	entextra "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinextra"
	entsetting "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/setting"
)

type queryHandler struct {
	*Handler
	stm   *ent.CoinBaseSelect
	infos []*npool.Coin
	total uint32
}

func (h *queryHandler) selectCoin(stm *ent.CoinBaseQuery) {
	h.stm = stm.Select(
		entbase.FieldID,
		entbase.FieldEntID,
		entbase.FieldName,
		entbase.FieldLogo,
		entbase.FieldPresale,
		entbase.FieldForPay,
		entbase.FieldUnit,
		entbase.FieldEnv,
		entbase.FieldReservedAmount,
		entbase.FieldDisabled,
		entbase.FieldCreatedAt,
		entbase.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCoin(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.CoinBase.Query().Where(entbase.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entbase.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entbase.EntID(*h.EntID))
	}
	h.selectCoin(stm)
	return nil
}

func (h *queryHandler) queryCoins(ctx context.Context, cli *ent.Client) error {
	stm, err := coincrud.SetQueryConds(cli.CoinBase.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectCoin(stm)
	return nil
}

func (h *queryHandler) queryJoinExtra(s *sql.Selector) {
	t := sql.Table(entextra.Table)
	s.
		LeftJoin(t).
		On(
			s.C(entbase.FieldEntID),
			t.C(entextra.FieldCoinTypeID),
		).
		AppendSelect(
			sql.As(t.C(entextra.FieldHomePage), "home_page"),
			sql.As(t.C(entextra.FieldSpecs), "specs"),
			sql.As(t.C(entextra.FieldStableUsd), "stable_usd"),
		)
}

func (h *queryHandler) queryJoinSetting(s *sql.Selector) {
	t1 := sql.Table(entsetting.Table)
	s.
		LeftJoin(t1).
		On(
			s.C(entbase.FieldEntID),
			t1.C(entsetting.FieldCoinTypeID),
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
			sql.As(t1.C(entsetting.FieldRefreshCurrency), "refresh_currency"),
			sql.As(t1.C(entsetting.FieldCheckNewAddressBalance), "check_new_address_balance"),
		)

	t2 := sql.Table(entbase.Table)
	s.
		LeftJoin(t2).
		On(
			t2.C(entbase.FieldEntID),
			t1.C(entsetting.FieldFeeCoinTypeID),
		).
		AppendSelect(
			sql.As(t2.C(entbase.FieldName), "fee_coin_name"),
			sql.As(t2.C(entbase.FieldLogo), "fee_coin_logo"),
			sql.As(t2.C(entbase.FieldUnit), "fee_coin_unit"),
			sql.As(t2.C(entbase.FieldEnv), "fee_coin_env"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinExtra(s)
		h.queryJoinSetting(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetCoin(ctx context.Context) (*npool.Coin, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoin(cli); err != nil {
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
		return nil, fmt.Errorf("too many record")
	}
	return handler.infos[0], nil
}

func (h *Handler) GetCoins(ctx context.Context) ([]*npool.Coin, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoins(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetCoinOnly(ctx context.Context) (*npool.Coin, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoins(_ctx, cli); err != nil {
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

	return handler.infos[0], nil
}
