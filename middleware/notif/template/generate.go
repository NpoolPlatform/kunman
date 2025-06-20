package template

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	notifmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	notifchancrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/channel"

	notifchanmw "github.com/NpoolPlatform/kunman/middleware/notif/notif/channel"

	email "github.com/NpoolPlatform/kunman/middleware/notif/template/email"
	frontend "github.com/NpoolPlatform/kunman/middleware/notif/template/frontend"
	sms "github.com/NpoolPlatform/kunman/middleware/notif/template/sms"
)

func (h *Handler) GenerateNotifs(
	ctx context.Context,
) (
	[]*notifmwpb.NotifReq,
	error,
) {
	const maxChannels = int32(100)
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
	chanHandler, err := notifchanmw.NewHandler(
		ctx,
		notifchanmw.WithOffset(0),
		notifchanmw.WithLimit(maxChannels),
	)
	if err != nil {
		return nil, err
	}
	chanHandler.Conds = &notifchancrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: *h.UsedFor},
	}

	chans, _, err := chanHandler.GetChannels(ctx)
	if err != nil {
		return nil, err
	}
	if len(chans) == 0 {
		return nil, fmt.Errorf("invalid channels")
	}

	reqs := []*notifmwpb.NotifReq{}
	for _, ch := range chans {
		switch ch.Channel {
		case basetypes.NotifChannel_ChannelEmail:
			emailHandler, err := email.NewHandler(
				ctx,
				email.WithAppID(&appID, true),
				email.WithUserID(&userID, true),
				email.WithUsedFor(h.UsedFor, true),
				email.WithVars(h.Vars, false),
			)
			if err != nil {
				return nil, err
			}
			_reqs, err := emailHandler.GenerateNotifs(ctx)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs...)
		case basetypes.NotifChannel_ChannelSMS:
			smsHandler, err := sms.NewHandler(
				ctx,
				sms.WithAppID(&appID, true),
				sms.WithUserID(&userID, true),
				sms.WithUsedFor(h.UsedFor, true),
				sms.WithVars(h.Vars, false),
			)
			if err != nil {
				return nil, err
			}
			_reqs, err := smsHandler.GenerateNotifs(ctx)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs...)
		case basetypes.NotifChannel_ChannelFrontend:
			frontendHandler, err := frontend.NewHandler(
				ctx,
				frontend.WithAppID(&appID, true),
				frontend.WithUserID(&userID, true),
				frontend.WithUsedFor(h.UsedFor, true),
				frontend.WithVars(h.Vars, false),
			)
			if err != nil {
				return nil, err
			}
			_reqs, err := frontendHandler.GenerateNotifs(ctx)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs...)
		}
	}

	return reqs, nil
}

func (h *Handler) GenerateText(
	ctx context.Context,
) (*npool.TextInfo, error) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.LangID == nil {
		return nil, fmt.Errorf("invalid langid")
	}
	if h.UsedFor == nil {
		return nil, fmt.Errorf("invalid usedfor")
	}
	if h.Channel == nil {
		return nil, fmt.Errorf("invalid channel")
	}
	appID := h.AppID.String()
	langID := h.LangID.String()
	switch *h.Channel {
	case basetypes.NotifChannel_ChannelEmail:
		emailHandler, err := email.NewHandler(
			ctx,
			email.WithAppID(&appID, true),
			email.WithLangID(&langID, true),
			email.WithUsedFor(h.UsedFor, true),
			email.WithVars(h.Vars, false),
		)
		if err != nil {
			return nil, err
		}
		return emailHandler.GenerateText(ctx)
	case basetypes.NotifChannel_ChannelSMS:
		smsHandler, err := sms.NewHandler(
			ctx,
			sms.WithAppID(&appID, true),
			sms.WithLangID(&langID, true),
			sms.WithUsedFor(h.UsedFor, true),
			sms.WithVars(h.Vars, false),
		)
		if err != nil {
			return nil, err
		}
		return smsHandler.GenerateText(ctx)
	case basetypes.NotifChannel_ChannelFrontend:
		frontendHandler, err := frontend.NewHandler(
			ctx,
			frontend.WithAppID(&appID, true),
			frontend.WithLangID(&langID, true),
			frontend.WithUsedFor(h.UsedFor, true),
			frontend.WithVars(h.Vars, false),
		)
		if err != nil {
			return nil, err
		}
		return frontendHandler.GenerateText(ctx)
	}

	return nil, fmt.Errorf("unknown channel")
}
