package orderuser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/orderuser"
	"github.com/google/uuid"
)

type sqlHandler struct {
	*Handler
	BondGoodUserID *uuid.UUID
	BondName       *string
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
func (h *sqlHandler) baseKeysDefault() error {
	if h.EntID == nil {
		entID, err := uuid.NewUUID()
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &entID
	}
	return nil
}

//nolint:gocognit,gocyclo
func (h *sqlHandler) baseKeysFiled() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[orderuser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[orderuser.FieldEntID] = string(strBytes)
	}
	if h.AppID != nil {
		strBytes, err := json.Marshal(*h.AppID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[orderuser.FieldAppID] = string(strBytes)
	}
	if h.UserID != nil {
		strBytes, err := json.Marshal(*h.UserID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[orderuser.FieldUserID] = string(strBytes)
	}
	if h.GoodUserID != nil {
		strBytes, err := json.Marshal(*h.GoodUserID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[orderuser.FieldGoodUserID] = string(strBytes)
		h.BondGoodUserID = h.GoodUserID
	}
	if h.Name != nil {
		strBytes, err := json.Marshal(*h.Name)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[orderuser.FieldName] = string(strBytes)
		h.BondName = h.Name
	}
	if h.ReadPageLink != nil {
		strBytes, err := json.Marshal(*h.ReadPageLink)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.baseVals[orderuser.FieldReadPageLink] = string(strBytes)
	}
	if h.BondGoodUserID == nil {
		return wlog.Errorf("please give gooduserid")
	}
	strBytes, err := json.Marshal(h.BondGoodUserID.String())
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[orderuser.FieldGoodUserID] = string(strBytes)

	if h.BondName == nil {
		return wlog.Errorf("please give name")
	}
	strBytes, err = json.Marshal(*h.BondName)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.bondVals[orderuser.FieldName] = string(strBytes)
	return nil
}

func (h *sqlHandler) idKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[orderuser.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.idVals[orderuser.FieldEntID] = string(strBytes)
	}
	return nil
}

//nolint:gocognit
func (h *sqlHandler) genCreateSQL() (string, error) {
	err := h.baseKeysDefault()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	err = h.baseKeysFiled()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	delete(h.baseVals, orderuser.FieldID)

	now := uint32(time.Now().Unix())
	h.baseVals[orderuser.FieldCreatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[orderuser.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[orderuser.FieldDeletedAt] = fmt.Sprintf("%v", 0)

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

	sql := fmt.Sprintf("insert into %v (%v) select * from (select %v) as tmp "+
		"where not exists (select * from %v where %v and deleted_at=0)"+
		"and exists "+
		"(select * from good_users A "+
		"left join root_users B on A.root_user_id=B.ent_id "+
		"left join app_pools C on B.pool_id=C.pool_id "+
		"where A.ent_id='%v' and C.app_id='%v' and A.deleted_at=0 and B.deleted_at=0 and C.deleted_at=0);",
		orderuser.Table,
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		orderuser.Table,
		strings.Join(bondVals, " AND "),
		h.GoodUserID.String(),
		h.AppID.String(),
	)

	return sql, nil
}

//nolint:gocognit
func (h *sqlHandler) genUpdateSQL() (string, error) {
	// get normal feilds
	err := h.baseKeysFiled()
	if err != nil {
		return "", wlog.WrapError(err)
	}
	delete(h.baseVals, orderuser.FieldID)
	delete(h.baseVals, orderuser.FieldEntID)

	now := uint32(time.Now().Unix())
	h.baseVals[orderuser.FieldUpdatedAt] = fmt.Sprintf("%v", now)

	if len(h.baseVals) == 0 {
		return "", wlog.Errorf("update nothing")
	}

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
		orderuser.Table,
		strings.Join(keys, ","),
		strings.Join(idKeys, " AND "),
		orderuser.Table,
		strings.Join(bondVals, " AND "),
	)

	return sql, nil
}
