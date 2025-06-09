package coin

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/coin"
	"github.com/google/uuid"
)

type sqlHandler struct {
	*Handler
	BondPoolID     *uuid.UUID
	BondCoinTypeID *uuid.UUID
	BondCoinType   *basetypes.CoinType
	bondVals       map[string]string
	baseVals       map[string]string
	idVals         map[string]string
}

func (h *Handler) newSQLHandler() *sqlHandler {
	return &sqlHandler{
		Handler:  h,
		bondVals: make(map[string]string),
		baseVals: make(map[string]string),
		idVals:   make(map[string]string),
	}
}

//nolint:gocognit
func (h *sqlHandler) baseKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldEntID] = string(strBytes)
	}
	if h.PoolID != nil {
		strBytes, err := json.Marshal(h.PoolID.String())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldPoolID] = string(strBytes)
		h.BondPoolID = h.PoolID
	}
	if h.CoinTypeID != nil {
		strBytes, err := json.Marshal(h.CoinTypeID.String())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldCoinTypeID] = string(strBytes)
		h.BondCoinTypeID = h.CoinTypeID
	}
	if h.CoinType != nil {
		strBytes, err := json.Marshal(h.CoinType.String())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldCoinType] = string(strBytes)
		h.BondCoinType = h.CoinType
	}
	if h.FeeRatio != nil {
		strBytes, err := json.Marshal(h.FeeRatio.String())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldFeeRatio] = string(strBytes)
	}
	if h.FixedRevenueAble != nil {
		strBytes, err := json.Marshal(*h.FixedRevenueAble)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldFixedRevenueAble] = string(strBytes)
	}
	if h.LeastTransferAmount != nil {
		strBytes, err := json.Marshal(*h.LeastTransferAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldLeastTransferAmount] = string(strBytes)
	}
	if h.BenefitIntervalSeconds != nil {
		strBytes, err := json.Marshal(*h.BenefitIntervalSeconds)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldBenefitIntervalSeconds] = string(strBytes)
	}
	if h.Remark != nil {
		strBytes, err := json.Marshal(*h.Remark)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[coin.FieldRemark] = string(strBytes)
	}

	if h.BondPoolID == nil {
		return wlog.Errorf("please give poolid")
	}
	strBytes, err := json.Marshal(h.BondPoolID.String())
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[coin.FieldPoolID] = string(strBytes)

	if h.BondCoinTypeID == nil {
		return wlog.Errorf("please give cointypeID")
	}
	strBytes, err = json.Marshal(h.BondCoinTypeID.String())
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[coin.FieldCoinTypeID] = string(strBytes)

	if h.BondCoinType == nil {
		return wlog.Errorf("please give cointype")
	}
	strBytes, err = json.Marshal(h.BondCoinType.String())
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[coin.FieldCoinType] = string(strBytes)
	return nil
}

func (h *sqlHandler) idKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[coin.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[coin.FieldEntID] = string(strBytes)
	}
	return nil
}

//nolint:gocognit
func (h *sqlHandler) genCreateSQL() (string, error) {
	err := h.baseKeys()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	delete(h.baseVals, coin.FieldID)

	now := uint32(time.Now().Unix())
	h.baseVals[coin.FieldCreatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[coin.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[coin.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}
	bondVals := []string{}

	for k, v := range h.baseVals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	for k, v := range h.bondVals {
		bondVals = append(bondVals, fmt.Sprintf("%v=%v", k, v))
	}

	sql := fmt.Sprintf("insert into %v (%v) select * from (select %v) as tmp where not exists (select * from %v where %v and deleted_at=0);",
		coin.Table,
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		coin.Table,
		strings.Join(bondVals, " AND "),
	)
	return sql, nil
}

//nolint:gocognit
func (h *sqlHandler) genUpdateSQL() (string, error) {
	// get normal feilds
	err := h.baseKeys()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	delete(h.baseVals, coin.FieldID)
	delete(h.baseVals, coin.FieldEntID)

	if len(h.baseVals) == 0 {
		return "", wlog.Errorf("update nothing")
	}

	now := uint32(time.Now().Unix())
	h.baseVals[coin.FieldUpdatedAt] = fmt.Sprintf("%v", now)

	keys := []string{}
	for k, v := range h.baseVals {
		keys = append(keys, fmt.Sprintf("%v=%v", k, v))
	}

	err = h.idKeys()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	if len(h.idVals) == 0 {
		return "", wlog.Errorf("have neither id and ent_id")
	}

	// get id and ent_id feilds
	idKeys := []string{}
	// get sub query feilds
	bondVals := []string{}
	for k, v := range h.idVals {
		idKeys = append(idKeys, fmt.Sprintf("%v=%v", k, v))
		bondVals = append(bondVals, fmt.Sprintf("tmp_table.%v!=%v", k, v))
	}

	for k, v := range h.bondVals {
		bondVals = append(bondVals, fmt.Sprintf("tmp_table.%v=%v", k, v))
	}

	sql := fmt.Sprintf("update %v set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from %v as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		coin.Table,
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		coin.Table,
		strings.Join(bondVals, " AND "),
	)
	return sql, nil
}
