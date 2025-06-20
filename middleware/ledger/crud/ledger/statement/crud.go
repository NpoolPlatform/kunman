package statement

import (
	"fmt"

	"github.com/shopspring/decimal"

	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entstatement "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/statement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID           *uint32
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	UserID       *uuid.UUID
	CurrencyID   *uuid.UUID
	CurrencyType *types.CurrencyType
	IOType       *types.IOType
	IOSubType    *types.IOSubType
	Amount       *decimal.Decimal
	IOExtra      *string
	CreatedAt    *uint32
	DeletedAt    *uint32
}

func CreateSet(c *ent.StatementCreate, in *Req) *ent.StatementCreate {
	if in.EntID != nil {
		c.SetEntID(*in.EntID)
	}
	if in.AppID != nil {
		c.SetAppID(*in.AppID)
	}
	if in.UserID != nil {
		c.SetUserID(*in.UserID)
	}
	if in.CurrencyID != nil {
		c.SetCurrencyID(*in.CurrencyID)
	}
	if in.CurrencyType != nil {
		c.SetCurrencyType(in.CurrencyType.String())
	}
	if in.IOType != nil {
		c.SetIoType(in.IOType.String())
	}
	if in.IOSubType != nil {
		c.SetIoSubType(in.IOSubType.String())
	}
	if in.Amount != nil {
		c.SetAmount(*in.Amount)
	}
	if in.IOExtra != nil {
		c.SetIoExtraV1(*in.IOExtra)
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(*in.CreatedAt)
	}
	return c
}

func UpdateSet(u *ent.StatementUpdateOne, req *Req) *ent.StatementUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	CurrencyID  *cruder.Cond
	IOType      *cruder.Cond
	IOSubType   *cruder.Cond
	Amount      *cruder.Cond
	IOExtra     *cruder.Cond
	StartAt     *cruder.Cond
	EndAt       *cruder.Cond
	IDs         *cruder.Cond
	EntIDs      *cruder.Cond
	IOSubTypes  *cruder.Cond
	CurrencyIDs *cruder.Cond
	UserIDs     *cruder.Cond
}

func SetQueryConds(q *ent.StatementQuery, conds *Conds) (*ent.StatementQuery, error) { //nolint
	q.Where(entstatement.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entstatement.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid op field %v", conds.EntID.Op)
		}
	}
	if conds.AppID != nil {
		appID, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entstatement.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid app id op field %v", conds.AppID.Op)
		}
	}
	if conds.UserID != nil {
		userID, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entstatement.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid user id op field %v", conds.UserID.Op)
		}
	}
	if conds.CurrencyID != nil {
		coinTypeID, ok := conds.CurrencyID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid currency id")
		}
		switch conds.CurrencyID.Op {
		case cruder.EQ:
			q.Where(entstatement.CurrencyID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid currency id op field %v", conds.CurrencyID.Op)
		}
	}
	if conds.IOType != nil {
		ioType, ok := conds.IOType.Val.(types.IOType)
		if !ok {
			return nil, fmt.Errorf("invalid io type %v", conds.IOType.Val)
		}
		switch conds.IOType.Op {
		case cruder.EQ:
			q.Where(entstatement.IoType(ioType.String()))
		default:
			return nil, fmt.Errorf("invalid io type op field %v", conds.IOType.Op)
		}
	}
	if conds.IOSubType != nil {
		ioSubType, ok := conds.IOSubType.Val.(types.IOSubType)
		if !ok {
			return nil, fmt.Errorf("invalid io type %v", conds.IOSubType.Val)
		}
		switch conds.IOSubType.Op {
		case cruder.EQ:
			q.Where(entstatement.IoSubType(ioSubType.String()))
		default:
			return nil, fmt.Errorf("invalid io sub type op field %v", conds.IOSubType.Op)
		}
	}
	if conds.Amount != nil {
		amount, ok := conds.Amount.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid amount %v", conds.Amount.Val)
		}
		switch conds.Amount.Op {
		case cruder.LT:
			q.Where(entstatement.AmountLT(amount))
		case cruder.GT:
			q.Where(entstatement.AmountGT(amount))
		case cruder.EQ:
			q.Where(entstatement.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid amount op field %v", conds.Amount.Op)
		}
	}
	if conds.IOExtra != nil {
		extra, ok := conds.IOExtra.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid io extra %v", conds.IOExtra.Val)
		}
		switch conds.IOExtra.Op {
		case cruder.LIKE:
			q.Where(entstatement.IoExtraV1Contains(extra))
		default:
			return nil, fmt.Errorf("invalid io extra op field %v", conds.IOExtra.Op)
		}
	}
	if conds.StartAt != nil {
		startAt, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid start  %v", conds.StartAt)
		}
		switch conds.StartAt.Op {
		case cruder.EQ:
			q.Where(entstatement.CreatedAtGTE(startAt))
		default:
			return nil, fmt.Errorf("invalid start at op field %v", conds.StartAt.Op)
		}
	}
	if conds.EndAt != nil {
		endAT, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid end at  %v", conds.EndAt)
		}
		switch conds.EndAt.Op {
		case cruder.EQ:
			q.Where(entstatement.CreatedAtLTE(endAT))
		default:
			return nil, fmt.Errorf("invalid end at op field %v", conds.EndAt.Op)
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids %v", conds.EntIDs.Val)
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entstatement.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid statement op field %v", conds.EntIDs.Op)
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids %v", conds.IDs.Val)
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entstatement.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid statement op field %v", conds.IDs.Op)
		}
	}
	if conds.IOSubTypes != nil {
		subTypes, ok := conds.IOSubTypes.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid io sub types %v", conds.IOSubTypes.Val)
		}
		switch conds.IOSubTypes.Op {
		case cruder.IN:
			q.Where(entstatement.IoSubTypeIn(subTypes...))
		default:
			return nil, fmt.Errorf("invalid io sub types op field %v", conds.IOSubTypes.Op)
		}
	}
	if conds.CurrencyIDs != nil {
		ids, ok := conds.CurrencyIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid currency ids %v", conds.CurrencyIDs.Val)
		}
		switch conds.CurrencyIDs.Op {
		case cruder.IN:
			q.Where(entstatement.CurrencyIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid currency ids op field %v", conds.CurrencyIDs.Op)
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user ids %v", conds.UserIDs.Val)
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entstatement.UserIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid user ids op field %v", conds.UserIDs.Op)
		}
	}
	return q, nil
}
