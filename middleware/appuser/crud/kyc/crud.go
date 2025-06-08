package kyc

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entkyc "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/kyc"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	UserID       *uuid.UUID
	DocumentType *basetypes.KycDocumentType
	IDNumber     *string
	FrontImg     *string
	BackImg      *string
	SelfieImg    *string
	EntityType   *basetypes.KycEntityType
	ReviewID     *uuid.UUID
	State        *basetypes.KycState
	DeletedAt    *uint32
}

func CreateSet(c *ent.KycCreate, req *Req) *ent.KycCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.DocumentType != nil {
		c.SetDocumentType(req.DocumentType.String())
	}
	if req.IDNumber != nil {
		c.SetIDNumber(*req.IDNumber)
	}
	if req.FrontImg != nil {
		c.SetFrontImg(*req.FrontImg)
	}
	if req.BackImg != nil {
		c.SetBackImg(*req.BackImg)
	}
	if req.SelfieImg != nil {
		c.SetSelfieImg(*req.SelfieImg)
	}
	if req.EntityType != nil {
		c.SetEntityType(req.EntityType.String())
	}
	if req.ReviewID != nil {
		c.SetReviewID(*req.ReviewID)
	}
	c.SetState(basetypes.KycState_Reviewing.String())
	return c
}

func UpdateSet(u *ent.KycUpdateOne, req *Req) *ent.KycUpdateOne {
	if req.DocumentType != nil {
		u.SetDocumentType(req.DocumentType.String())
	}
	if req.IDNumber != nil {
		u.SetIDNumber(*req.IDNumber)
	}
	if req.FrontImg != nil {
		u.SetFrontImg(*req.FrontImg)
	}
	if req.BackImg != nil {
		u.SetBackImg(*req.BackImg)
	}
	if req.SelfieImg != nil {
		u.SetSelfieImg(*req.SelfieImg)
	}
	if req.EntityType != nil {
		u.SetEntityType(req.EntityType.String())
	}
	if req.ReviewID != nil {
		u.SetReviewID(*req.ReviewID)
	}
	if req.State != nil {
		u.SetState(req.State.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID           *cruder.Cond
	EntID        *cruder.Cond
	AppID        *cruder.Cond
	UserID       *cruder.Cond
	DocumentType *cruder.Cond
	IDNumber     *cruder.Cond
	EntityType   *cruder.Cond
	ReviewID     *cruder.Cond
	State        *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.KycQuery, conds *Conds) (*ent.KycQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entkyc.ID(id))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entkyc.EntID(id))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entkyc.AppID(id))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entkyc.UserID(id))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	if conds.IDNumber != nil {
		idNumber, ok := conds.IDNumber.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid id number")
		}
		switch conds.IDNumber.Op {
		case cruder.EQ:
			q.Where(entkyc.IDNumber(idNumber))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	if conds.DocumentType != nil {
		docType, ok := conds.DocumentType.Val.(basetypes.KycDocumentType)
		if !ok {
			return nil, fmt.Errorf("invalid document type")
		}
		switch conds.DocumentType.Op {
		case cruder.EQ:
			q.Where(entkyc.DocumentType(docType.String()))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	if conds.EntityType != nil {
		entType, ok := conds.EntityType.Val.(basetypes.KycEntityType)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}
		switch conds.EntityType.Op {
		case cruder.EQ:
			q.Where(entkyc.EntityType(entType.String()))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	if conds.ReviewID != nil {
		id, ok := conds.ReviewID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid reviewid")
		}
		switch conds.ReviewID.Op {
		case cruder.EQ:
			q.Where(entkyc.ReviewID(id))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(basetypes.KycState)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(entkyc.State(state.String()))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	return q, nil
}
