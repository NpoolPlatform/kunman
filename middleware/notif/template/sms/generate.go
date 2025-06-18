package sms

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template"

	notifmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"

	smstemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/sms"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	tmplreplace "github.com/NpoolPlatform/kunman/middleware/notif/template/replace"
)

func (h *Handler) GenerateNotifs(ctx context.Context) ([]*notifmwpb.NotifReq, error) {
	const maxTemplates = int32(100)
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid userid")
	}
	if h.UsedFor == nil {
		return nil, fmt.Errorf("invalid usedfor")
	}

	appID := h.AppID.String()
	userID := h.UserID.String()

	smstmplHandler, err := NewHandler(
		ctx,
		WithOffset(0),
		WithLimit(maxTemplates),
	)
	if err != nil {
		return nil, err
	}
	smstmplHandler.Conds = &smstemplatecrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UsedFor: &cruder.Cond{Op: cruder.EQ, Val: *h.UsedFor},
	}

	tmpls, _, err := smstmplHandler.GetSMSTemplates(ctx)
	if err != nil {
		return nil, err
	}
	if len(tmpls) == 0 {
		return nil, fmt.Errorf("invalid sms template")
	}

	reqs := []*notifmwpb.NotifReq{}
	for _, tmpl := range tmpls {
		title := tmplreplace.ReplaceAll(tmpl.Subject, h.Vars)
		content := tmplreplace.ReplaceAll(tmpl.Message, h.Vars)
		useTemplate := true
		channel1 := basetypes.NotifChannel_ChannelSMS

		reqs = append(reqs, &notifmwpb.NotifReq{
			AppID:       &appID,
			UserID:      &userID,
			LangID:      &tmpl.LangID,
			EventType:   h.UsedFor,
			UseTemplate: &useTemplate,
			Title:       &title,
			Content:     &content,
			Channel:     &channel1,
		})
	}

	return reqs, nil
}

func (h *Handler) GenerateText(ctx context.Context) (*npool.TextInfo, error) {
	smstmplHandler, err := NewHandler(
		ctx,
	)
	if err != nil {
		return nil, err
	}
	smstmplHandler.Conds = &smstemplatecrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		LangID:  &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
		UsedFor: &cruder.Cond{Op: cruder.EQ, Val: *h.UsedFor},
	}

	tmpl, err := smstmplHandler.GetSMSTemplateOnly(ctx)
	if err != nil {
		return nil, err
	}
	if tmpl == nil {
		return nil, nil
	}

	title := tmplreplace.ReplaceAll(tmpl.Subject, h.Vars)
	content := tmplreplace.ReplaceAll(tmpl.Message, h.Vars)

	return &npool.TextInfo{
		Subject: title,
		Content: content,
	}, nil
}
