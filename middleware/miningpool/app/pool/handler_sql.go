package apppool

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/apppool"
	"github.com/google/uuid"
)

type sqlHandler struct {
	*Handler
	BondAppID  *uuid.UUID
	BondPoolID *uuid.UUID
	bondVals   map[string]string
	baseVals   map[string]string
	idVals     map[string]string
}

func (h *Handler) newSQLHandler() *sqlHandler {
	return &sqlHandler{
		Handler:  h,
		bondVals: make(map[string]string),
		baseVals: make(map[string]string),
		idVals:   make(map[string]string),
	}
}

func (h *sqlHandler) baseKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[apppool.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[apppool.FieldEntID] = string(strBytes)
	}
	if h.AppID != nil {
		strBytes, err := json.Marshal(*h.AppID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[apppool.FieldAppID] = string(strBytes)
		h.BondAppID = h.AppID
	}
	if h.PoolID != nil {
		strBytes, err := json.Marshal(*h.PoolID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[apppool.FieldPoolID] = string(strBytes)
		h.BondPoolID = h.PoolID
	}

	if h.BondAppID == nil {
		return wlog.Errorf("please give appid")
	}
	strBytes, err := json.Marshal(*h.BondAppID)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[apppool.FieldAppID] = string(strBytes)

	if h.BondPoolID == nil {
		return wlog.Errorf("please give poolid")
	}
	strBytes, err = json.Marshal(*h.BondPoolID)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[apppool.FieldPoolID] = string(strBytes)
	return nil
}

func (h *sqlHandler) idKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[apppool.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[apppool.FieldEntID] = string(strBytes)
	}
	return nil
}

//nolint:gocognit
func (h *sqlHandler) genCreateSQL() (string, error) {
	err := h.baseKeys()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	delete(h.baseVals, apppool.FieldID)

	now := uint32(time.Now().Unix())
	h.baseVals[apppool.FieldCreatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[apppool.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[apppool.FieldDeletedAt] = fmt.Sprintf("%v", 0)

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
		apppool.Table,
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		apppool.Table,
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
	delete(h.baseVals, apppool.FieldID)
	delete(h.baseVals, apppool.FieldEntID)

	if len(h.baseVals) == 0 {
		return "", wlog.Errorf("update nothing")
	}

	now := uint32(time.Now().Unix())
	h.baseVals[apppool.FieldUpdatedAt] = fmt.Sprintf("%v", now)

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
	idVals := []string{}
	// get sub query feilds
	bondVals := []string{}
	for k, v := range h.idVals {
		idVals = append(idVals, fmt.Sprintf("%v=%v", k, v))
		bondVals = append(bondVals, fmt.Sprintf("tmp_table.%v!=%v", k, v))
	}

	for k, v := range h.bondVals {
		bondVals = append(bondVals, fmt.Sprintf("tmp_table.%v=%v", k, v))
	}

	sql := fmt.Sprintf("update %v set %v where %v and deleted_at=0 and  not exists (select 1 from(select * from %v as tmp_table where %v and tmp_table.deleted_at=0 limit 1) as tmp);",
		apppool.Table,
		strings.Join(keys, ","),
		strings.Join(idVals, " AND "),
		apppool.Table,
		strings.Join(bondVals, " AND "),
	)
	return sql, nil
}
