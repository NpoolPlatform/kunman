package app

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	appusertypes "github.com/NpoolPlatform/kunman/message/basetypes/appuser/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/app"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID                       *uint32
	EntID                    *uuid.UUID
	CreatedBy                *uuid.UUID
	Name                     *string
	Logo                     *string
	Description              *string
	Banned                   *bool
	BanMessage               *string
	SignupMethods            []basetypes.SignMethod
	ExtSigninMethods         []basetypes.SignMethod
	RecaptchaMethod          *basetypes.RecaptchaMethod
	KycEnable                *bool
	SigninVerifyEnable       *bool
	InvitationCodeMust       *bool
	CreateInvitationCodeWhen *basetypes.CreateInvitationCodeWhen
	MaxTypedCouponsPerOrder  *uint32
	Maintaining              *bool
	CouponWithdrawEnable     *bool
	CommitButtonTargets      []string
	UserID                   *uuid.UUID
	ResetUserMethod          *appusertypes.ResetUserMethod
	Reqs                     []*AppReq
	Conds                    *appcrud.Conds
	Offset                   int32
	Limit                    int32
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

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
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

func WithCreatedBy(createdBy *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if createdBy == nil {
			if must {
				return fmt.Errorf("invalid createdby")
			}
			return nil
		}
		_createdBy, err := uuid.Parse(*createdBy)
		if err != nil {
			return err
		}
		// TODO: confirm creator exist
		h.CreatedBy = &_createdBy
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		const leastNameLen = 8
		if len(*name) < leastNameLen {
			return fmt.Errorf("name %v too short", *name)
		}
		// TODO: confirm name not exist
		h.Name = name
		return nil
	}
}

func WithLogo(logo *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if logo == nil {
			if must {
				return fmt.Errorf("invalid logo")
			}
			return nil
		}
		const leastLogoLen = 8
		if len(*logo) < leastLogoLen {
			return fmt.Errorf("logo %v too short", *logo)
		}
		h.Logo = logo
		return nil
	}
}

func WithDescription(description *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Description = description
		return nil
	}
}

func WithBanned(banned *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Banned = banned
		return nil
	}
}

func WithBanMessage(message *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.BanMessage = message
		return nil
	}
}

func WithSignupMethods(methods []basetypes.SignMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, method := range methods {
			switch method {
			case basetypes.SignMethod_Username:
				return fmt.Errorf("username signup not implemented")
			case basetypes.SignMethod_Mobile:
			case basetypes.SignMethod_Email:
			default:
				return fmt.Errorf("signup method %v invalid", method)
			}
		}
		h.SignupMethods = methods
		return nil
	}
}

func WithExtSigninMethods(methods []basetypes.SignMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, method := range methods {
			switch method {
			case basetypes.SignMethod_Twitter:
				fallthrough //nolint
			case basetypes.SignMethod_Github:
				fallthrough //nolint
			case basetypes.SignMethod_Facebook:
				fallthrough //nolint
			case basetypes.SignMethod_Linkedin:
				fallthrough //nolint
			case basetypes.SignMethod_Wechat:
				fallthrough //nolint
			case basetypes.SignMethod_Google:
				return fmt.Errorf("%v signin not implemented", method)
			default:
				return fmt.Errorf("ext signin method %v invalid", method)
			}
		}
		h.ExtSigninMethods = methods
		return nil
	}
}

func WithRecaptchaMethod(method *basetypes.RecaptchaMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if method == nil {
			if must {
				return fmt.Errorf("invalid recaptchamethod")
			}
			return nil
		}
		switch *method {
		case basetypes.RecaptchaMethod_GoogleRecaptchaV3:
		case basetypes.RecaptchaMethod_NoRecaptcha:
		default:
			return fmt.Errorf("recaptcha method %v invalid", *method)
		}
		h.RecaptchaMethod = method
		return nil
	}
}

func WithKycEnable(enable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.KycEnable = enable
		return nil
	}
}

func WithSigninVerifyEnable(enable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SigninVerifyEnable = enable
		return nil
	}
}

func WithInvitationCodeMust(enable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.InvitationCodeMust = enable
		return nil
	}
}

func WithCreateInvitationCodeWhen(when *basetypes.CreateInvitationCodeWhen, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if when == nil {
			if must {
				return fmt.Errorf("invalid createinvitationcodewhen")
			}
			return nil
		}
		switch *when {
		case basetypes.CreateInvitationCodeWhen_Registration:
		case basetypes.CreateInvitationCodeWhen_SetToKol:
		case basetypes.CreateInvitationCodeWhen_HasPaidOrder:
		default:
			return fmt.Errorf("create invitation when %v invalid", *when)
		}
		h.CreateInvitationCodeWhen = when
		return nil
	}
}

func WithMaxTypedCouponsPerOrder(count *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MaxTypedCouponsPerOrder = count
		return nil
	}
}

func WithMaintaining(enable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Maintaining = enable
		return nil
	}
}

func WithCouponWithdrawEnable(enable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.CouponWithdrawEnable = enable
		return nil
	}
}

func WithCommitButtonTargets(targets []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.CommitButtonTargets = targets
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

func WithResetUserMethod(method *appusertypes.ResetUserMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if method == nil {
			if must {
				return fmt.Errorf("invalid reset user method")
			}
			return nil
		}
		switch *method {
		case appusertypes.ResetUserMethod_Normal:
		case appusertypes.ResetUserMethod_Link:
		default:
			return fmt.Errorf("invalid reset method %v", method)
		}
		h.ResetUserMethod = method
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &appcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.CreatedBy != nil {
			id, err := uuid.Parse(conds.GetCreatedBy().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CreatedBy = &cruder.Cond{
				Op:  conds.GetCreatedBy().GetOp(),
				Val: id,
			}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
			}
		}
		return nil
	}
}

func WithReqs(reqs []*npool.AppReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			_req, err := Grpc2CrudReq(req)
			if err != nil {
				return err
			}
			h.Reqs = append(h.Reqs, _req)
		}
		return nil
	}
}
