package rootuser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/rootuser"
	"github.com/google/uuid"
)

type sqlHandler struct {
	*Handler
	BondPoolID *uuid.UUID
	BondName   *string
	BondEmail  *string
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

//nolint:gocognit
func (h *sqlHandler) baseKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[rootuser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[rootuser.FieldEntID] = string(strBytes)
	}
	if h.PoolID != nil {
		strBytes, err := json.Marshal(h.PoolID.String())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[rootuser.FieldPoolID] = string(strBytes)
		h.BondPoolID = h.PoolID
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[rootuser.FieldName] = string(strBytes)
		h.BondName = h.Name
	}
	if h.Email != nil {
		strBytes, err := json.Marshal(*h.Email)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[rootuser.FieldEmail] = string(strBytes)
		h.BondEmail = h.Email
	}
	if h.AuthToken != nil {
		strBytes, err := json.Marshal(*h.AuthToken)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[rootuser.FieldAuthToken] = string(strBytes)
	}
	if h.AuthTokenSalt != nil {
		strBytes, err := json.Marshal(*h.AuthTokenSalt)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[rootuser.FieldAuthTokenSalt] = string(strBytes)
	}
	if h.Authed != nil {
		strBytes, err := json.Marshal(*h.Authed)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[rootuser.FieldAuthed] = string(strBytes)
	}
	if h.Remark != nil {
		strBytes, err := json.Marshal(*h.Remark)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[rootuser.FieldRemark] = string(strBytes)
	}

	if h.BondPoolID == nil {
		return wlog.Errorf("please give miningpooltype")
	}
	strBytes, err := json.Marshal(h.BondPoolID.String())
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[rootuser.FieldPoolID] = string(strBytes)

	if h.BondName == nil {
		return wlog.Errorf("please give name")
	}
	strBytes, err = json.Marshal(*h.BondName)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[rootuser.FieldName] = string(strBytes)

	if h.BondEmail == nil {
		return wlog.Errorf("please give email")
	}
	strBytes, err = json.Marshal(*h.BondEmail)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[rootuser.FieldEmail] = string(strBytes)
	return nil
}

func (h *sqlHandler) idKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[rootuser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[rootuser.FieldEntID] = string(strBytes)
	}
	return nil
}

func (h *sqlHandler) genCreateSQL() (string, error) {
	err := h.baseKeys()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	delete(h.baseVals, rootuser.FieldID)

	now := uint32(time.Now().Unix())
	h.baseVals[rootuser.FieldCreatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[rootuser.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[rootuser.FieldDeletedAt] = fmt.Sprintf("%v", 0)

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

	sql := fmt.Sprintf("insert into %v (%v) select * from (select %v) as tmp where not exists (select * from %v where %v and  deleted_at=0);",
		rootuser.Table,
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		rootuser.Table,
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
	delete(h.baseVals, rootuser.FieldID)
	delete(h.baseVals, rootuser.FieldEntID)

	if len(h.baseVals) == 0 {
		return "", wlog.Errorf("update nothing")
	}

	now := uint32(time.Now().Unix())
	h.baseVals[rootuser.FieldUpdatedAt] = fmt.Sprintf("%v", now)

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
		rootuser.Table,
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		rootuser.Table,
		strings.Join(bondVals, " AND "),
	)

	return sql, nil
}
