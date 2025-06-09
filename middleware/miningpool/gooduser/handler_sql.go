package gooduser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/gooduser"
)

type sqlHandler struct {
	*Handler
	BondName *string
	bondVals map[string]string
	baseVals map[string]string
	idVals   map[string]string
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
		h.baseVals[gooduser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[gooduser.FieldEntID] = string(strBytes)
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[gooduser.FieldName] = string(strBytes)
		h.BondName = h.Name
	}
	if h.RootUserID != nil {
		strBytes, err := json.Marshal(*h.RootUserID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[gooduser.FieldRootUserID] = string(strBytes)
	}
	if h.ReadPageLink != nil {
		strBytes, err := json.Marshal(*h.ReadPageLink)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[gooduser.FieldReadPageLink] = string(strBytes)
	}

	if h.BondName == nil {
		return wlog.Errorf("please give name")
	}
	strBytes, err := json.Marshal(*h.BondName)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[gooduser.FieldName] = string(strBytes)
	return nil
}

func (h *sqlHandler) idKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[gooduser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[gooduser.FieldEntID] = string(strBytes)
	}
	return nil
}

//nolint:gocognit
func (h *sqlHandler) genCreateSQL() (string, error) {
	err := h.baseKeys()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	delete(h.baseVals, gooduser.FieldID)

	now := uint32(time.Now().Unix())
	h.baseVals[gooduser.FieldCreatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[gooduser.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[gooduser.FieldDeletedAt] = fmt.Sprintf("%v", 0)

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
		gooduser.Table,
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		gooduser.Table,
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
	delete(h.baseVals, gooduser.FieldID)
	delete(h.baseVals, gooduser.FieldEntID)

	if len(h.baseVals) == 0 {
		return "", wlog.Errorf("update nothing")
	}

	now := uint32(time.Now().Unix())
	h.baseVals[gooduser.FieldUpdatedAt] = fmt.Sprintf("%v", now)

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
		gooduser.Table,
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		gooduser.Table,
		strings.Join(bondVals, " AND "),
	)

	return sql, nil
}
