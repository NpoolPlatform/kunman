package executor

import (
	"context"
	"fmt"
	"net/mail"
	"regexp"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/notif/notification/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	applangmwpb "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/applang"
	notifmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	emailtmplmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/email"
	smstmplmgrpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/sms"
	sendmwpb "github.com/NpoolPlatform/kunman/message/third/middleware/v1/send"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	applangmw "github.com/NpoolPlatform/kunman/middleware/g11n/applang"
	notifmw "github.com/NpoolPlatform/kunman/middleware/notif/notif"
	emailtmplmw "github.com/NpoolPlatform/kunman/middleware/notif/template/email"
	smstmplmw "github.com/NpoolPlatform/kunman/middleware/notif/template/sms"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type notifHandler struct {
	*notifmwpb.Notif
	persistent     chan interface{}
	done           chan interface{}
	notifiable     bool
	eventNotifs    []*notifmwpb.Notif
	user           *usermwpb.User
	lang           *applangmwpb.Lang
	messageRequest *sendmwpb.SendMessageInput
}

func (h *notifHandler) getUser(ctx context.Context) error {
	handler, err := usermw.NewHandler(
		ctx,
		usermw.WithEntID(&h.UserID, true),
		usermw.WithAppID(&h.AppID, true),
	)
	if err != nil {
		return err
	}

	user, err := handler.GetUser(ctx)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("invalid user")
	}
	h.user = user
	return nil
}

func (h *notifHandler) getLang(ctx context.Context) error {
	conds := &applangmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
	}
	if h.user.SelectedLangID != nil {
		conds.LangID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.user.SelectedLangID}
	} else {
		conds.Main = &basetypes.BoolVal{Op: cruder.EQ, Value: true}
	}

	handler, err := applangmw.NewHandler(
		ctx,
		applangmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	lang, err := handler.GetLangOnly(ctx)
	if err != nil {
		return err
	}
	if lang != nil {
		h.lang = lang
		return nil
	}

	conds.Main = &basetypes.BoolVal{Op: cruder.EQ, Value: true}

	handler, err = applangmw.NewHandler(
		ctx,
		applangmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	lang, err = handler.GetLangOnly(ctx)
	if err != nil {
		return err
	}
	if lang == nil {
		return fmt.Errorf("invalid main lang")
	}
	h.lang = lang
	return nil
}

func (h *notifHandler) generateEmailMessage(ctx context.Context) error {
	conds := &emailtmplmwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		LangID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.LangID},
		UsedFor: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.EventType)},
	}
	handler, err := emailtmplmw.NewHandler(
		ctx,
		emailtmplmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	tmpl, err := handler.GetEmailTemplateOnly(ctx)
	if err != nil {
		return err
	}
	if tmpl == nil {
		return fmt.Errorf("invalid template")
	}

	if _, err := mail.ParseAddress(h.user.EmailAddress); err != nil {
		h.notifiable = false
		return nil
	}

	h.messageRequest.From = tmpl.Sender
	h.messageRequest.To = h.user.EmailAddress
	h.messageRequest.ToCCs = tmpl.CCTos
	h.messageRequest.ReplyTos = tmpl.ReplyTos
	h.messageRequest.AccountType = basetypes.SignMethod_Email
	return nil
}

func (h *notifHandler) generateSMSMessage(ctx context.Context) error {
	conds := &smstmplmgrpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		LangID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.LangID},
		UsedFor: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.EventType)},
	}
	handler, err := smstmplmw.NewHandler(
		ctx,
		smstmplmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	tmpl, err := handler.GetSMSTemplateOnly(ctx)
	if err != nil {
		return err
	}
	if tmpl == nil {
		return fmt.Errorf("invalid template")
	}

	re := regexp.MustCompile(
		`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[` +
			`\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?)` +
			`{0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)` +
			`[\-\.\ \\\/]?(\d+))?$`,
	)
	if !re.MatchString(h.user.PhoneNO) {
		h.notifiable = false
		return nil
	}

	h.messageRequest.To = h.user.PhoneNO
	h.messageRequest.AccountType = basetypes.SignMethod_Mobile
	return nil
}

func (h *notifHandler) generateMessageRequest(ctx context.Context) error {
	h.messageRequest = &sendmwpb.SendMessageInput{
		Subject: h.Title,
		Content: h.Content,
	}
	switch h.Channel {
	case basetypes.NotifChannel_ChannelEmail:
		return h.generateEmailMessage(ctx)
	case basetypes.NotifChannel_ChannelSMS:
		return h.generateSMSMessage(ctx)
	}
	h.notifiable = false
	return nil
}

func (h *notifHandler) getEventNotifs(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &notifmwpb.Conds{
		Channel: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.Channel)},
		EventID: &basetypes.StringVal{Op: cruder.EQ, Value: h.EventID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.UserID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
	}

	for {
		handler, err := notifmw.NewHandler(
			ctx,
			notifmw.WithConds(conds),
			notifmw.WithOffset(offset),
			notifmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		notifs, _, err := handler.GetNotifs(ctx)
		if err != nil {
			return err
		}
		if len(notifs) == 0 {
			return nil
		}
		h.eventNotifs = notifs
		offset += limit
	}
}

//nolint:gocritic
func (h *notifHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Notif", h.Notif,
			"Notifiable", h.notifiable,
			"Error", *err,
		)
	}

	persistentNotif := &types.PersistentNotif{
		Notif:          h.Notif,
		MessageRequest: h.messageRequest,
		EventNotifs:    h.eventNotifs,
	}
	// TODO: We don't know how to process err here
	if h.notifiable {
		asyncfeed.AsyncFeed(ctx, persistentNotif, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentNotif, h.done)
}

//nolint:gocritic
func (h *notifHandler) exec(ctx context.Context) error {
	var err error
	defer h.final(ctx, &err)

	if err = h.getUser(ctx); err != nil {
		return err
	}
	if err = h.getLang(ctx); err != nil {
		return err
	}
	h.notifiable = h.lang.LangID == h.LangID
	if !h.notifiable {
		return nil
	}
	if err = h.generateMessageRequest(ctx); err != nil {
		return err
	}
	if err = h.getEventNotifs(ctx); err != nil {
		return err
	}

	return nil
}
