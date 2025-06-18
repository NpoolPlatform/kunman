package generate

import (
	"context"
	"fmt"
	"strings"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/contact"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/contact"
)

func (h *Handler) GenerateContact(ctx context.Context) (*npool.TextInfo, error) {
	h.Conds = &crud.Conds{
		AppID: &cruder.Cond{
			Op: cruder.EQ, Val: *h.AppID,
		},
		UsedFor: &cruder.Cond{
			Op: cruder.EQ, Val: basetypes.UsedFor(basetypes.UsedFor_value[h.UsedFor.String()]),
		},
		AccountType: &cruder.Cond{
			Op: cruder.EQ, Val: basetypes.SignMethod(basetypes.SignMethod_value[basetypes.SignMethod_Email.String()]),
		},
	}
	info, err := h.GetContactOnly(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("contact not exist")
	}

	_body := fmt.Sprintf("From: %v<br>Name: %v<br>%v", *h.Sender, *h.SenderName, *h.Body)
	body := strings.ReplaceAll(_body, "\n", "<br>")

	return &npool.TextInfo{
		Subject:  *h.Subject,
		Content:  body,
		From:     info.Sender,
		To:       info.Account,
		ReplyTos: []string{*h.Sender},
	}, nil
}
