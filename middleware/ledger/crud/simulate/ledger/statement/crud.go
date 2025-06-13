package statement

import (
	"fmt"

	"github.com/shopspring/decimal"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entstatement "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/simulatestatement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uint32
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	UserID     *uuid.UUID
	CoinTypeID *uuid.UUID
	IOType     *basetypes.IOType
	IOSubType  *basetypes.IOSubType
	Amount     *decimal.Decimal
	IOExtra    *string
	SendCoupon *bool
	Cashable   *bool
	CreatedAt  *uint32
	DeletedAt  *uint32
}

func CreateSet(c *ent.SimulateStatementCreate, in *Req) *ent.SimulateStatementCreate {
	if in.EntID != nil {
		c.SetEntID(*in.EntID)
	}
	if in.AppID != nil {
		c.SetAppID(*in.AppID)
	}
	if in.UserID != nil {
		c.SetUserID(*in.UserID)
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(*in.CoinTypeID)
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
		c.SetIoExtra(*in.IOExtra)
	}
	if in.SendCoupon != nil {
		c.SetSendCoupon(*in.SendCoupon)
	}
	if in.Cashable != nil {
		c.SetCashable(*in.Cashable)
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(*in.CreatedAt)
	}
	return c
}

func UpdateSet(u *ent.SimulateStatementUpdateOne, req *Req) *ent.SimulateStatementUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	CoinTypeID  *cruder.Cond
	IOType      *cruder.Cond
	IOSubType   *cruder.Cond
	Amount      *cruder.Cond
	IOExtra     *cruder.Cond
	StartAt     *cruder.Cond
	EndAt       *cruder.Cond
	IDs         *cruder.Cond
	EntIDs      *cruder.Cond
	IOSubTypes  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	UserIDs     *cruder.Cond
	SendCoupon  *cruder.Cond
	Cashable    *cruder.Cond
}

func SetQueryConds(q *ent.SimulateStatementQuery, conds *Conds) (*ent.SimulateStatementQuery, error) { //nolint
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
	if conds.CoinTypeID != nil {
		coinTypeID, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type id")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entstatement.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.IOType != nil {
		ioType, ok := conds.IOType.Val.(basetypes.IOType)
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
		ioSubType, ok := conds.IOSubType.Val.(basetypes.IOSubType)
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
			q.Where(entstatement.IoExtraContains(extra))
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
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type ids %v", conds.CoinTypeIDs.Val)
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entstatement.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid coin type ids op field %v", conds.CoinTypeIDs.Op)
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
	if conds.SendCoupon != nil {
		sendcoupon, ok := conds.SendCoupon.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid sendcoupon")
		}
		switch conds.SendCoupon.Op {
		case cruder.EQ:
			q.Where(entstatement.SendCoupon(sendcoupon))
		default:
			return nil, fmt.Errorf("invalid sendcoupon op field %v", conds.SendCoupon.Op)
		}
	}
	if conds.Cashable != nil {
		cashable, ok := conds.Cashable.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid cashable")
		}
		switch conds.Cashable.Op {
		case cruder.EQ:
			q.Where(entstatement.Cashable(cashable))
		default:
			return nil, fmt.Errorf("invalid cashable op field %v", conds.Cashable.Op)
		}
	}
	return q, nil
}
