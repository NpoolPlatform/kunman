package frontend

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	templatemwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/frontend"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	frontendtemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/frontend"

	"github.com/google/uuid"
)

type Handler struct {
	ID      *uint32
	EntID   *uuid.UUID
	AppID   *uuid.UUID
	LangID  *uuid.UUID
	UsedFor *basetypes.UsedFor
	Title   *string
	Content *string
	UserID  *uuid.UUID
	Vars    *templatemwpb.TemplateVars
	Reqs    []*frontendtemplatecrud.Req
	Conds   *frontendtemplatecrud.Conds
	Offset  int32
	Limit   int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithLangID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid langid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.LangID = &_id
		return nil
	}
}

//nolint:gocyclo
func WithUsedFor(_usedFor *basetypes.UsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _usedFor == nil {
			if must {
				return fmt.Errorf("invalid usedfor")
			}
			return nil
		}
		switch *_usedFor {
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Announcement:
		case basetypes.UsedFor_GoodBenefit1:
		case basetypes.UsedFor_UpdateEmail:
		case basetypes.UsedFor_UpdateMobile:
		case basetypes.UsedFor_UpdatePassword:
		case basetypes.UsedFor_UpdateGoogleAuth:
		case basetypes.UsedFor_NewLogin:
		case basetypes.UsedFor_OrderCompleted:
		case basetypes.UsedFor_ResetPassword:
		case basetypes.UsedFor_OrderChildsRenewNotify:
		case basetypes.UsedFor_OrderChildsRenew:
		case basetypes.UsedFor_WithdrawReviewNotify:
		default:
			return fmt.Errorf("invalid usedfor")
		}
		h.UsedFor = _usedFor
		return nil
	}
}

func WithTitle(title *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if title == nil {
			if must {
				return fmt.Errorf("invalid title")
			}
			return nil
		}
		if *title == "" {
			return fmt.Errorf("invalid title")
		}
		h.Title = title
		return nil
	}
}

func WithContent(content *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if content == nil {
			if must {
				return fmt.Errorf("invalid content")
			}
			return nil
		}
		if *content == "" {
			return fmt.Errorf("invalid content")
		}
		h.Content = content
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithVars(vars *templatemwpb.TemplateVars, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if vars == nil {
			if must {
				return fmt.Errorf("invalid vars")
			}
			return nil
		}
		h.Vars = vars
		return nil
	}
}

// nolint:gocyclo
func WithReqs(reqs []*npool.FrontendTemplateReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*frontendtemplatecrud.Req{}
		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.LangID == nil {
					return fmt.Errorf("invalid langid")
				}
				if req.UsedFor == nil {
					return fmt.Errorf("invalid usedfor")
				}
			}
			_req := &frontendtemplatecrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.AppID != nil {
				id, err := uuid.Parse(req.GetAppID())
				if err != nil {
					return err
				}
				_req.AppID = &id
			}
			if req.LangID != nil {
				id, err := uuid.Parse(req.GetLangID())
				if err != nil {
					return err
				}
				_req.LangID = &id
			}
			if req.UsedFor != nil {
				switch req.GetUsedFor() {
				case basetypes.UsedFor_WithdrawalRequest:
				case basetypes.UsedFor_WithdrawalCompleted:
				case basetypes.UsedFor_DepositReceived:
				case basetypes.UsedFor_KYCApproved:
				case basetypes.UsedFor_KYCRejected:
				case basetypes.UsedFor_Announcement:
				case basetypes.UsedFor_GoodBenefit1:
				case basetypes.UsedFor_UpdateEmail:
				case basetypes.UsedFor_UpdateMobile:
				case basetypes.UsedFor_UpdatePassword:
				case basetypes.UsedFor_UpdateGoogleAuth:
				case basetypes.UsedFor_NewLogin:
				case basetypes.UsedFor_OrderCompleted:
				case basetypes.UsedFor_OrderChildsRenewNotify:
				case basetypes.UsedFor_OrderChildsRenew:
				case basetypes.UsedFor_WithdrawReviewNotify:
				default:
					return fmt.Errorf("invalid usedfor")
				}
				_req.UsedFor = req.UsedFor
			}
			if req.Title != nil {
				_req.Title = req.Title
			}
			if req.Content != nil {
				_req.Content = req.Content
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

//nolint
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &frontendtemplatecrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op: conds.GetEntID().GetOp(), Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.LangID != nil {
			id, err := uuid.Parse(conds.GetLangID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.LangID = &cruder.Cond{
				Op:  conds.GetLangID().GetOp(),
				Val: id,
			}
		}
		if conds.UsedFor != nil {
			switch conds.GetUsedFor().GetValue() {
			case uint32(basetypes.UsedFor_WithdrawalRequest):
			case uint32(basetypes.UsedFor_WithdrawalCompleted):
			case uint32(basetypes.UsedFor_DepositReceived):
			case uint32(basetypes.UsedFor_KYCApproved):
			case uint32(basetypes.UsedFor_KYCRejected):
			case uint32(basetypes.UsedFor_Announcement):
			case uint32(basetypes.UsedFor_GoodBenefit1):
			case uint32(basetypes.UsedFor_UpdateEmail):
			case uint32(basetypes.UsedFor_UpdateMobile):
			case uint32(basetypes.UsedFor_UpdatePassword):
			case uint32(basetypes.UsedFor_UpdateGoogleAuth):
			case uint32(basetypes.UsedFor_NewLogin):
			case uint32(basetypes.UsedFor_OrderCompleted):
			case uint32(basetypes.UsedFor_OrderChildsRenewNotify):
			case uint32(basetypes.UsedFor_OrderChildsRenew):
			case uint32(basetypes.UsedFor_WithdrawReviewNotify):
			default:
				return fmt.Errorf("invalid usedfor")
			}
			usedFor := conds.GetUsedFor().GetValue()
			h.Conds.UsedFor = &cruder.Cond{
				Op:  conds.GetUsedFor().GetOp(),
				Val: basetypes.UsedFor(usedFor),
			}
		}
		if conds.AppIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetAppIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.AppIDs = &cruder.Cond{Op: conds.GetAppIDs().GetOp(), Val: ids}
		}
		if conds.LangIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetLangIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.LangIDs = &cruder.Cond{Op: conds.GetLangIDs().GetOp(), Val: ids}
		}
		if conds.UsedFors != nil {
			usedFors := []string{}
			for _, usedFor := range conds.GetUsedFors().GetValue() {
				switch usedFor {
				case uint32(basetypes.UsedFor_Signup):
				case uint32(basetypes.UsedFor_Signin):
				case uint32(basetypes.UsedFor_Update):
				case uint32(basetypes.UsedFor_Contact):
				case uint32(basetypes.UsedFor_SetWithdrawAddress):
				case uint32(basetypes.UsedFor_Withdraw):
				case uint32(basetypes.UsedFor_CreateInvitationCode):
				case uint32(basetypes.UsedFor_SetCommission):
				case uint32(basetypes.UsedFor_SetTransferTargetUser):
				case uint32(basetypes.UsedFor_Transfer):
				case uint32(basetypes.UsedFor_WithdrawalRequest):
				case uint32(basetypes.UsedFor_WithdrawalCompleted):
				case uint32(basetypes.UsedFor_DepositReceived):
				case uint32(basetypes.UsedFor_KYCApproved):
				case uint32(basetypes.UsedFor_KYCRejected):
				case uint32(basetypes.UsedFor_Announcement):
				case uint32(basetypes.UsedFor_GoodBenefit1):
				case uint32(basetypes.UsedFor_UpdateEmail):
				case uint32(basetypes.UsedFor_UpdateMobile):
				case uint32(basetypes.UsedFor_UpdatePassword):
				case uint32(basetypes.UsedFor_UpdateGoogleAuth):
				case uint32(basetypes.UsedFor_NewLogin):
				case uint32(basetypes.UsedFor_OrderCompleted):
				case uint32(basetypes.UsedFor_OrderChildsRenewNotify):
				case uint32(basetypes.UsedFor_OrderChildsRenew):
				case uint32(basetypes.UsedFor_WithdrawReviewNotify):
				default:
					return fmt.Errorf("invalid usedfor")
				}
				_usedFor := basetypes.UsedFor(usedFor).String()
				usedFors = append(usedFors, _usedFor)
			}
			h.Conds.UsedFors = &cruder.Cond{Op: conds.GetUsedFors().GetOp(), Val: usedFors}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
