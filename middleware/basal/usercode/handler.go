package usercode

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
)

type Handler struct {
	Prefix      *string
	AppID       *string
	Account     *string
	AccountType *basetypes.SignMethod
	UsedFor     *basetypes.UsedFor
	VCode       *string
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

func WithPrefix(prefix *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if prefix == nil {
			return fmt.Errorf("prefix is empty")
		}
		h.Prefix = prefix
		return nil
	}
}

func WithAppID(appID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}
		h.AppID = appID
		return nil
	}
}

func WithAccount(account *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if account == nil {
			return fmt.Errorf("account is empty")
		}
		h.Account = account
		return nil
	}
}

func WithAccountType(_type *basetypes.SignMethod) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch *_type {
		case basetypes.SignMethod_Email:
		case basetypes.SignMethod_Mobile:
		case basetypes.SignMethod_Google:
		default:
			return fmt.Errorf("AccountType is invalid")
		}

		h.AccountType = _type
		return nil
	}
}

func WithUsedFor(usedFor *basetypes.UsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch *usedFor {
		case basetypes.UsedFor_Signup:
		case basetypes.UsedFor_Signin:
		case basetypes.UsedFor_Update:
		case basetypes.UsedFor_Contact:
		case basetypes.UsedFor_SetWithdrawAddress:
		case basetypes.UsedFor_Withdraw:
		case basetypes.UsedFor_CreateInvitationCode:
		case basetypes.UsedFor_SetCommission:
		case basetypes.UsedFor_SetTransferTargetUser:
		case basetypes.UsedFor_Transfer:
		default:
			return fmt.Errorf("UsedFor %v is invalid", *usedFor)
		}

		h.UsedFor = usedFor
		return nil
	}
}

func WithCode(code *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if code == nil {
			return fmt.Errorf("code is empty")
		}
		h.VCode = code
		return nil
	}
}
