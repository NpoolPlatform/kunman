package control

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entappuserctrl "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appusercontrol"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID              *uuid.UUID
	AppID              *uuid.UUID
	UserID             *uuid.UUID
	GoogleAuthVerified *bool
	SigninVerifyType   *basetypes.SignMethod
	Kol                *bool
	KolConfirmed       *bool
	SelectedLangID     *uuid.UUID
}

func CreateSet(c *ent.AppUserControlCreate, req *Req) *ent.AppUserControlCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.GoogleAuthVerified != nil {
		c.SetGoogleAuthenticationVerified(*req.GoogleAuthVerified)
	}
	if req.SigninVerifyType != nil {
		c.SetSigninVerifyType(req.SigninVerifyType.String())
	}
	if req.Kol != nil {
		c.SetKol(*req.Kol)
	}
	if req.KolConfirmed != nil {
		c.SetKolConfirmed(*req.KolConfirmed)
	}
	if req.SelectedLangID != nil {
		c.SetSelectedLangID(*req.SelectedLangID)
	}
	return c
}

func UpdateSet(u *ent.AppUserControlUpdateOne, req *Req) *ent.AppUserControlUpdateOne {
	if req.GoogleAuthVerified != nil {
		u.SetGoogleAuthenticationVerified(*req.GoogleAuthVerified)
	}
	if req.SigninVerifyType != nil {
		u.SetSigninVerifyType(req.SigninVerifyType.String())
	}
	if req.Kol != nil {
		u.SetKol(*req.Kol)
	}
	if req.KolConfirmed != nil {
		u.SetKolConfirmed(*req.KolConfirmed)
	}
	if req.SelectedLangID != nil {
		u.SetSelectedLangID(*req.SelectedLangID)
	}
	return u
}

type Conds struct {
	EntID        *cruder.Cond
	AppID        *cruder.Cond
	UserID       *cruder.Cond
	Kol          *cruder.Cond
	KolConfirmed *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.AppUserControlQuery, conds *Conds) (*ent.AppUserControlQuery, error) {
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
			q.Where(entappuserctrl.EntID(id))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappuserctrl.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entappuserctrl.UserID(id))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	if conds.Kol != nil {
		kol, ok := conds.Kol.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid kol")
		}
		switch conds.Kol.Op {
		case cruder.EQ:
			q.Where(entappuserctrl.Kol(kol))
		case cruder.NEQ:
			q.Where(entappuserctrl.KolNEQ(kol))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	if conds.KolConfirmed != nil {
		confirmed, ok := conds.KolConfirmed.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid kol confirmed")
		}
		switch conds.KolConfirmed.Op {
		case cruder.EQ:
			q.Where(entappuserctrl.KolConfirmed(confirmed))
		case cruder.NEQ:
			q.Where(entappuserctrl.KolConfirmedNEQ(confirmed))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	return q, nil
}
