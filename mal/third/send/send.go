package send

import (
	"context"
	"fmt"

	aws "github.com/NpoolPlatform/kunman/mal/third/aws"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

type sendHandler struct {
	*Handler
}

func (h *sendHandler) validate() error {
	if h.To == "" {
		return fmt.Errorf("invalid to")
	}
	if h.Content == "" {
		return fmt.Errorf("invalid content")
	}

	switch h.AccountType {
	case basetypes.SignMethod_Email:
		if h.From == "" {
			return fmt.Errorf("invalid from")
		}
		if h.Subject == "" {
			return fmt.Errorf("invalid subject")
		}
	case basetypes.SignMethod_Mobile:
	default:
		return fmt.Errorf("invalid accountType")
	}

	return nil
}

func (h *Handler) SendMessage(ctx context.Context) error {
	handler := &sendHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return err
	}

	switch h.AccountType {
	case basetypes.SignMethod_Email:
		h.ReplyTos = append(h.ReplyTos, h.From)
		return aws.SendEmail(h.Subject, h.Content, h.From, h.To, h.ReplyTos...)
	case basetypes.SignMethod_Mobile:
		return aws.SendSMS(h.Content, h.To)
	default:
		return fmt.Errorf("invalid accountType")
	}
}
