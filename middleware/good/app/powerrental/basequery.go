package powerrental

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entapplegacypowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/applegacypowerrental"
	entapppowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/apppowerrental"
	entappgoodstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appstock"
	entdevicetype "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/deviceinfo"
	entmanufacturer "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/devicemanufacturer"
	entextrainfo "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/extrainfo"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entgoodcoin "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoin"
	entgoodreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodreward"
	entpowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/powerrental"
	entstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/stock"
	entvendorbrand "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/vendorbrand"
	entvendorlocation "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/vendorlocation"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodBaseSelect
}

func (h *baseQueryHandler) selectAppGoodBase(stm *ent.AppGoodBaseQuery) *ent.AppGoodBaseSelect {
	return stm.Select(entappgoodbase.FieldID)
}

func (h *baseQueryHandler) queryAppGoodBase(cli *ent.Client) error {
	if h.AppGoodID == nil {
		return wlog.Errorf("invalid appgoodid")
	}
	h.stmSelect = h.selectAppGoodBase(
		cli.AppGoodBase.
			Query().
			Where(
				entappgoodbase.DeletedAt(0),
				entappgoodbase.EntID(*h.AppGoodID),
			),
	)
	return nil
}

func (h *baseQueryHandler) queryAppGoodBases(cli *ent.Client) (*ent.AppGoodBaseSelect, error) {
	stm, err := appgoodbasecrud.SetQueryConds(cli.AppGoodBase.Query(), h.AppGoodBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectAppGoodBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldID),
			t1.C(entappgoodbase.FieldID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldAppID),
			t1.C(entappgoodbase.FieldGoodID),
			sql.As(t1.C(entappgoodbase.FieldPurchasable), "app_good_purchasable"),
			sql.As(t1.C(entappgoodbase.FieldOnline), "app_good_online"),
			t1.C(entappgoodbase.FieldEnableProductPage),
			t1.C(entappgoodbase.FieldProductPage),
			t1.C(entappgoodbase.FieldVisible),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			t1.C(entappgoodbase.FieldDisplayIndex),
			t1.C(entappgoodbase.FieldBanner),
			t1.C(entappgoodbase.FieldCreatedAt),
			t1.C(entappgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinGoodBase(s *sql.Selector) {
	t1 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		).
		OnP(
			sql.Or(
				sql.EQ(t1.C(entgoodbase.FieldGoodType), types.GoodType_PowerRental.String()),
				sql.EQ(t1.C(entgoodbase.FieldGoodType), types.GoodType_LegacyPowerRental.String()),
			),
		).
		AppendSelect(
			t1.C(entgoodbase.FieldGoodType),
			t1.C(entgoodbase.FieldBenefitType),
			sql.As(t1.C(entgoodbase.FieldName), "good_name"),
			sql.As(t1.C(entgoodbase.FieldServiceStartAt), "good_service_start_at"),
			sql.As(t1.C(entgoodbase.FieldStartMode), "good_start_mode"),
			t1.C(entgoodbase.FieldTestOnly),
			t1.C(entgoodbase.FieldBenefitIntervalHours),
			sql.As(t1.C(entgoodbase.FieldPurchasable), "good_purchasable"),
			sql.As(t1.C(entgoodbase.FieldOnline), "good_online"),
			t1.C(entgoodbase.FieldState),
		)
}

func (h *baseQueryHandler) queryJoinGoodCoin(s *sql.Selector) error {
	t := sql.Table(entgoodcoin.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t.C(entgoodcoin.FieldGoodID),
		).
		Distinct()
	if h.GoodCoinConds.CoinTypeID != nil {
		id, ok := h.GoodCoinConds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid cointypeid")
		}
		s.OnP(
			sql.EQ(t.C(entgoodcoin.FieldCoinTypeID), id),
		)
		s.Where(
			sql.EQ(t.C(entgoodcoin.FieldCoinTypeID), id),
		)
	}
	if h.GoodCoinConds.CoinTypeIDs != nil {
		uids, ok := h.GoodCoinConds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid cointypeids")
		}
		_uids := func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return
		}()
		s.OnP(
			sql.In(t.C(entgoodcoin.FieldCoinTypeID), _uids...),
		)
		s.Where(
			sql.In(t.C(entgoodcoin.FieldCoinTypeID), _uids...),
		)
	}
	return nil
}

func (h *baseQueryHandler) queryJoinGoodReward(s *sql.Selector) {
	t := sql.Table(entgoodreward.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t.C(entgoodreward.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entextrainfo.FieldDeletedAt), 0),
		).
		AppendSelect(
			t.C(entgoodreward.FieldLastRewardAt),
		)
}

func (h *baseQueryHandler) queryJoinExtraInfo(s *sql.Selector) {
	t := sql.Table(entextrainfo.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodbase.FieldEntID),
			t.C(entextrainfo.FieldAppGoodID),
		).
		OnP(
			sql.EQ(t.C(entextrainfo.FieldDeletedAt), 0),
		).
		AppendSelect(
			t.C(entextrainfo.FieldLikes),
			t.C(entextrainfo.FieldDislikes),
			t.C(entextrainfo.FieldScoreCount),
			t.C(entextrainfo.FieldRecommendCount),
			t.C(entextrainfo.FieldCommentCount),
			t.C(entextrainfo.FieldScore),
		)
}

func (h *baseQueryHandler) queryJoinGoodStock(s *sql.Selector) {
	t1 := sql.Table(entstock.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entstock.FieldGoodID),
		).
		AppendSelect(
			sql.As(t1.C(entstock.FieldEntID), "good_stock_id"),
			sql.As(t1.C(entstock.FieldTotal), "good_total"),
			sql.As(t1.C(entstock.FieldSpotQuantity), "good_spot_quantity"),
		)
}

func (h *baseQueryHandler) queryJoinAppGoodStock(s *sql.Selector) {
	t1 := sql.Table(entappgoodstock.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entappgoodstock.FieldAppGoodID),
		).
		AppendSelect(
			sql.As(t1.C(entappgoodstock.FieldEntID), "app_good_stock_id"),
			sql.As(t1.C(entappgoodstock.FieldReserved), "app_good_reserved"),
			sql.As(t1.C(entappgoodstock.FieldSpotQuantity), "app_good_spot_quantity"),
			sql.As(t1.C(entappgoodstock.FieldLocked), "app_good_locked"),
			sql.As(t1.C(entappgoodstock.FieldWaitStart), "app_good_wait_start"),
			sql.As(t1.C(entappgoodstock.FieldInService), "app_good_in_service"),
			sql.As(t1.C(entappgoodstock.FieldSold), "app_good_sold"),
		)
}

func (h *baseQueryHandler) queryJoinAppPowerRental(s *sql.Selector) error {
	t1 := sql.Table(entapppowerrental.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entapppowerrental.FieldAppGoodID),
		)
	if h.ID != nil {
		s.OnP(
			sql.EQ(t1.C(entapppowerrental.FieldID), *h.ID),
		)
	}
	if h.AppPowerRentalConds.ID != nil {
		id, ok := h.AppPowerRentalConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(
			sql.EQ(t1.C(entapppowerrental.FieldID), id),
		)
	}
	if h.EntID != nil {
		s.OnP(
			sql.EQ(t1.C(entapppowerrental.FieldEntID), *h.EntID),
		)
	}
	if h.AppPowerRentalConds.EntID != nil {
		uid, ok := h.AppPowerRentalConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(
			sql.EQ(t1.C(entapppowerrental.FieldEntID), uid),
		)
	}
	s.AppendSelect(
		t1.C(entapppowerrental.FieldID),
		t1.C(entapppowerrental.FieldEntID),
		t1.C(entapppowerrental.FieldAppGoodID),
		t1.C(entapppowerrental.FieldCancelMode),
		t1.C(entapppowerrental.FieldCancelableBeforeStartSeconds),
		t1.C(entapppowerrental.FieldEnableSetCommission),
		t1.C(entapppowerrental.FieldMinOrderAmount),
		t1.C(entapppowerrental.FieldMaxOrderAmount),
		t1.C(entapppowerrental.FieldMaxUserAmount),
		t1.C(entapppowerrental.FieldMinOrderDurationSeconds),
		t1.C(entapppowerrental.FieldMaxOrderDurationSeconds),
		t1.C(entapppowerrental.FieldSaleStartAt),
		t1.C(entapppowerrental.FieldSaleEndAt),
		t1.C(entapppowerrental.FieldSaleMode),
		t1.C(entapppowerrental.FieldFixedDuration),
		t1.C(entapppowerrental.FieldPackageWithRequireds),
		t1.C(entapppowerrental.FieldUnitPrice),
		sql.As(t1.C(entapppowerrental.FieldServiceStartAt), "app_good_service_start_at"),
		sql.As(t1.C(entapppowerrental.FieldStartMode), "app_good_start_mode"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinAppLegacyPowerRental(s *sql.Selector) {
	t1 := sql.Table(entapplegacypowerrental.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entapplegacypowerrental.FieldAppGoodID),
		).
		OnP(
			sql.EQ(t1.C(entapplegacypowerrental.FieldDeletedAt), 0),
		).
		AppendSelect(
			t1.C(entapplegacypowerrental.FieldTechniqueFeeRatio),
		)
}

func (h *baseQueryHandler) queryJoinPowerRental(s *sql.Selector) error {
	t1 := sql.Table(entpowerrental.Table)
	t2 := sql.Table(entdevicetype.Table)
	t3 := sql.Table(entmanufacturer.Table)
	t4 := sql.Table(entvendorlocation.Table)
	t5 := sql.Table(entvendorbrand.Table)

	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entpowerrental.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entpowerrental.FieldDeletedAt), 0),
		)
	if h.PowerRentalConds.GoodID != nil {
		uid, ok := h.PowerRentalConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t1.C(entpowerrental.FieldGoodID), uid))
	}
	if h.PowerRentalConds.GoodIDs != nil {
		uids, ok := h.PowerRentalConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(sql.In(t1.C(entpowerrental.FieldGoodID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return
		}()...))
	}
	s.Join(t2).
		On(
			t1.C(entpowerrental.FieldDeviceTypeID),
			t2.C(entdevicetype.FieldEntID),
		).
		Join(t3).
		On(
			t2.C(entdevicetype.FieldManufacturerID),
			t3.C(entmanufacturer.FieldEntID),
		).
		Join(t4).
		On(
			t1.C(entpowerrental.FieldVendorLocationID),
			t4.C(entvendorlocation.FieldEntID),
		).
		Join(t5).
		On(
			t4.C(entvendorlocation.FieldBrandID),
			t5.C(entvendorbrand.FieldEntID),
		)
	s.AppendSelect(
		t1.C(entpowerrental.FieldGoodID),
		t1.C(entpowerrental.FieldDeviceTypeID),
		t1.C(entpowerrental.FieldVendorLocationID),
		t1.C(entpowerrental.FieldQuantityUnit),
		t1.C(entpowerrental.FieldQuantityUnitAmount),
		t1.C(entpowerrental.FieldDeliveryAt),
		t1.C(entpowerrental.FieldUnitLockDeposit),
		t1.C(entpowerrental.FieldDurationDisplayType),
		t1.C(entpowerrental.FieldUnitLockDeposit),
		t1.C(entpowerrental.FieldStockMode),

		sql.As(t2.C(entdevicetype.FieldType), "device_type"),
		sql.As(t2.C(entdevicetype.FieldPowerConsumption), "device_power_consumption"),
		sql.As(t2.C(entdevicetype.FieldShipmentAt), "device_shipment_at"),
		sql.As(t3.C(entmanufacturer.FieldName), "device_manufacturer_name"),
		sql.As(t3.C(entmanufacturer.FieldLogo), "device_manufacturer_logo"),

		sql.As(t4.C(entvendorlocation.FieldCountry), "vendor_country"),
		sql.As(t4.C(entvendorlocation.FieldProvince), "vendor_province"),
		sql.As(t5.C(entvendorbrand.FieldName), "vendor_brand"),
		sql.As(t5.C(entvendorbrand.FieldLogo), "vendor_logo"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinGoodBase(s)
		h.queryJoinGoodStock(s)
		h.queryJoinAppGoodStock(s)
		h.queryJoinGoodReward(s)
		h.queryJoinExtraInfo(s)
		if err := h.queryJoinAppPowerRental(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppPowerRental", "Error", err)
		}
		h.queryJoinAppLegacyPowerRental(s)
		if err := h.queryJoinPowerRental(s); err != nil {
			logger.Sugar().Errorw("queryJoinPowerRental", "Error", err)
		}
		if err := h.queryJoinGoodCoin(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodCoin", "Error", err)
		}
	})
}
