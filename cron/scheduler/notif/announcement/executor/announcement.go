package executor

import (
	"context"
	"fmt"
	"net/mail"
	"regexp"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/notif/announcement/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	applangmwpb "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/applang"
	ancmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
	ancsendmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/sendstate"
	ancusermwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/user"
	emailtmplmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/email"
	smstmplmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/sms"
	sendmwpb "github.com/NpoolPlatform/kunman/message/third/middleware/v1/send"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	applangmw "github.com/NpoolPlatform/kunman/middleware/g11n/applang"
	anchandler "github.com/NpoolPlatform/kunman/middleware/notif/announcement/handler"
	ancsendmw "github.com/NpoolPlatform/kunman/middleware/notif/announcement/sendstate"
	ancusermw "github.com/NpoolPlatform/kunman/middleware/notif/announcement/user"
	emailtmplmw "github.com/NpoolPlatform/kunman/middleware/notif/template/email"
	smstmplmw "github.com/NpoolPlatform/kunman/middleware/notif/template/sms"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type announcementHandler struct {
	*ancmwpb.Announcement
	persistent chan interface{}
	done       chan interface{}
	sendStats  map[string]*ancsendmwpb.SendState
}

func (h *announcementHandler) getSendStats(ctx context.Context, users []*usermwpb.User) error {
	uids := []string{}
	for _, user := range users {
		uids = append(uids, user.EntID)
	}

	conds := &ancsendmwpb.Conds{
		AppID:          &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		AnnouncementID: &basetypes.StringVal{Op: cruder.EQ, Value: h.EntID},
		Channel:        &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.Channel)},
		UserIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: uids},
	}
	handler, err := ancsendmw.NewHandler(
		ctx,
		ancsendmw.WithConds(conds),
		anchandler.WithOffset(0),
		anchandler.WithLimit(int32(len(uids))),
	)
	if err != nil {
		return err
	}

	stats, _, err := handler.GetSendStates(ctx)
	if err != nil {
		return err
	}
	for _, stat := range stats {
		h.sendStats[stat.UserID] = stat
	}
	return nil
}

func (h *announcementHandler) getLang(ctx context.Context, user *usermwpb.User) (*applangmwpb.Lang, error) {
	conds := &applangmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
	}
	if user.SelectedLangID != nil {
		conds.LangID = &basetypes.StringVal{Op: cruder.EQ, Value: *user.SelectedLangID}
	} else {
		conds.Main = &basetypes.BoolVal{Op: cruder.EQ, Value: true}
	}

	handler, err := applangmw.NewHandler(
		ctx,
		applangmw.WithConds(conds),
	)
	if err != nil {
		return nil, err
	}

	lang, err := handler.GetLangOnly(ctx)
	if err != nil {
		return nil, err
	}
	if lang != nil {
		return lang, nil
	}
	if user.SelectedLangID == nil {
		return nil, fmt.Errorf("invalid mainlang")
	}
	conds.LangID = nil
	conds.Main = &basetypes.BoolVal{Op: cruder.EQ, Value: true}

	handler, err = applangmw.NewHandler(
		ctx,
		applangmw.WithConds(conds),
	)
	if err != nil {
		return nil, err
	}

	lang, err = handler.GetLangOnly(ctx)
	if err != nil {
		return nil, err
	}
	if lang == nil {
		return nil, fmt.Errorf("invalid mainlang")
	}
	return lang, nil
}

func (h *announcementHandler) emailRequest(ctx context.Context, user *usermwpb.User) (*sendmwpb.SendMessageInput, error) {
	req := &sendmwpb.SendMessageInput{
		Subject: h.Title,
		Content: h.Content,
	}

	conds := &emailtmplmwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		LangID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.LangID},
		UsedFor: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.UsedFor_Announcement)},
	}
	handler, err := emailtmplmw.NewHandler(
		ctx,
		emailtmplmw.WithConds(conds),
	)
	if err != nil {
		return nil, err
	}

	tmpl, err := handler.GetEmailTemplateOnly(ctx)
	if err != nil {
		return nil, err
	}
	if tmpl == nil {
		return nil, fmt.Errorf("invalid template")
	}

	req.From = tmpl.Sender
	req.To = user.EmailAddress
	req.ToCCs = tmpl.CCTos
	req.ReplyTos = tmpl.ReplyTos
	req.AccountType = basetypes.SignMethod_Email

	return req, nil
}

func (h *announcementHandler) smsRequest(ctx context.Context, user *usermwpb.User) (*sendmwpb.SendMessageInput, error) {
	req := &sendmwpb.SendMessageInput{
		Subject: h.Title,
		Content: h.Content,
	}

	conds := &smstmplmwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		LangID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.LangID},
		UsedFor: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.UsedFor_Announcement)},
	}
	handler, err := smstmplmw.NewHandler(
		ctx,
		smstmplmw.WithConds(conds),
	)
	if err != nil {
		return nil, err
	}

	tmpl, err := handler.GetSMSTemplateOnly(ctx)
	if err != nil {
		return nil, err
	}
	if tmpl == nil {
		return nil, fmt.Errorf("invalid template")
	}

	req.To = user.PhoneNO
	req.AccountType = basetypes.SignMethod_Mobile
	return req, nil
}

func (h *announcementHandler) unicast(ctx context.Context, user *usermwpb.User) error {
	if _, ok := h.sendStats[user.EntID]; ok {
		return nil
	}

	switch h.Channel {
	case basetypes.NotifChannel_ChannelEmail:
		if _, err := mail.ParseAddress(user.EmailAddress); err != nil {
			return nil
		}
	case basetypes.NotifChannel_ChannelSMS:
		re := regexp.MustCompile(
			`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[` +
				`\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?)` +
				`{0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)` +
				`[\-\.\ \\\/]?(\d+))?$`,
		)
		if !re.MatchString(user.PhoneNO) {
			return nil
		}
	default:
		return nil
	}

	lang, err := h.getLang(ctx, user)
	if err != nil {
		return err
	}
	if lang.LangID != h.LangID {
		return nil
	}

	var req *sendmwpb.SendMessageInput
	switch h.Channel {
	case basetypes.NotifChannel_ChannelEmail:
		if req, err = h.emailRequest(ctx, user); err != nil {
			return err
		}
	case basetypes.NotifChannel_ChannelSMS:
		if req, err = h.smsRequest(ctx, user); err != nil {
			return err
		}
	}

	asyncfeed.AsyncFeed(ctx, &types.PersistentAnnouncement{
		Announcement:   h.Announcement,
		SendAppID:      user.AppID,
		SendUserID:     user.EntID,
		MessageRequest: req,
	}, h.persistent)

	return nil
}

func (h *announcementHandler) multicastUsers(ctx context.Context, users []*usermwpb.User) error {
	if err := h.getSendStats(ctx, users); err != nil {
		return err
	}
	for _, user := range users {
		if err := h.unicast(ctx, user); err != nil {
			logger.Sugar().Errorw(
				"multicastUsers",
				"AnnouncementID", h.ID,
				"User", user,
				"Error", err,
			)
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func (h *announcementHandler) broadcast(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &usermwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
	}

	for {
		handler, err := usermw.NewHandler(
			ctx,
			usermw.WithConds(conds),
			usermw.WithOffset(offset),
			usermw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		users, _, err := handler.GetUsers(ctx)
		if err != nil {
			return err
		}

		if len(users) == 0 {
			return nil
		}

		if err := h.multicastUsers(ctx, users); err != nil {
			logger.Sugar().Errorw(
				"broadcast",
				"AnnouncementID", h.ID,
				"Error", err,
			)
			return err
		}

		offset += limit
	}
}

func (h *announcementHandler) multicast(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &ancusermwpb.Conds{
		AppID:          &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		AnnouncementID: &basetypes.StringVal{Op: cruder.EQ, Value: h.EntID},
	}

	for {
		ancHandler, err := ancusermw.NewHandler(
			ctx,
			ancusermw.WithConds(conds),
			anchandler.WithOffset(offset),
			anchandler.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		ancUsers, _, err := ancHandler.GetAnnouncementUsers(ctx)
		if err != nil {
			return err
		}
		if len(ancUsers) == 0 {
			return nil
		}

		offset += limit

		uids := []string{}
		for _, user := range ancUsers {
			uids = append(uids, user.UserID)
		}

		_conds := &usermwpb.Conds{
			EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: uids},
		}
		userHandler, err := usermw.NewHandler(
			ctx,
			usermw.WithConds(_conds),
			usermw.WithOffset(0),
			usermw.WithLimit(int32(len(uids))),
		)
		if err != nil {
			return err
		}

		users, _, err := userHandler.GetUsers(ctx)
		if err != nil {
			return err
		}
		if len(users) == 0 {
			continue
		}

		if err := h.multicastUsers(ctx, users); err != nil {
			return err
		}
	}
}

func (h *announcementHandler) exec(ctx context.Context) error {
	h.sendStats = map[string]*ancsendmwpb.SendState{}

	defer asyncfeed.AsyncFeed(ctx, h.Announcement, h.done)

	switch h.AnnouncementType {
	case basetypes.NotifType_NotifBroadcast:
		if err := h.broadcast(ctx); err != nil {
			return err
		}
	case basetypes.NotifType_NotifMulticast:
		if err := h.multicast(ctx); err != nil {
			return err
		}
	}

	return nil
}
