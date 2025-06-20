package notif

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	templatemwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template"
	usernotifcrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/user"
	usernotifmw "github.com/NpoolPlatform/kunman/middleware/notif/notif/user"
	templatemwcli "github.com/NpoolPlatform/kunman/middleware/notif/template"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type generateHandler struct {
	*Handler
}

func (h *generateHandler) getNotifUsers(ctx context.Context, userID *uuid.UUID, notifType basetypes.NotifType, eventType basetypes.UsedFor) ([]string, error) {
	userIDs := []string{}
	const maxLimit = int32(100)
	switch notifType {
	case basetypes.NotifType_NotifMulticast:
		usernotifHandler, err := usernotifmw.NewHandler(
			ctx,
			usernotifmw.WithOffset(0),
			usernotifmw.WithLimit(maxLimit),
		)
		if err != nil {
			return nil, err
		}
		usernotifHandler.Conds = &usernotifcrud.Conds{
			AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			EventType: &cruder.Cond{Op: cruder.EQ, Val: eventType},
		}
		userNotifs, _, err := usernotifHandler.GetNotifUsers(ctx)
		if err != nil {
			return nil, err
		}
		if len(userNotifs) == 0 {
			return nil, fmt.Errorf("invalid user notif")
		}
		for _, row := range userNotifs {
			userIDs = append(userIDs, row.UserID)
		}
	case basetypes.NotifType_NotifUnicast:
		if userID == nil {
			return nil, fmt.Errorf("invalid userid")
		}
		userID := userID.String()
		userIDs = append(userIDs, userID)
	default:
		return nil, fmt.Errorf("invalid notiftype")
	}
	return userIDs, nil
}

func (h *generateHandler) createUserNotifs(
	ctx context.Context,
	appID, eventID, userID string,
	extra *string,
	eventType basetypes.UsedFor,
	notifType basetypes.NotifType,
	vars *templatemwpb.TemplateVars,
) ([]*npool.NotifReq, error) {
	reqs := []*npool.NotifReq{}
	templateHandler, err := templatemwcli.NewHandler(
		ctx,
		templatemwcli.WithAppID(&appID, true),
		templatemwcli.WithUserID(&userID, true),
		templatemwcli.WithUsedFor(&eventType, true),
		templatemwcli.WithVars(vars, false),
	)
	if err != nil {
		return nil, err
	}
	_reqs, err := templateHandler.GenerateNotifs(ctx)
	if err != nil {
		return nil, err
	}
	for _, req := range _reqs {
		req.Extra = extra
		req.NotifType = &notifType
		req.EventID = &eventID
	}
	reqs = append(reqs, _reqs...)
	return reqs, nil
}

func (h *Handler) GenerateNotifs(ctx context.Context) ([]*npool.Notif, error) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.EventType == nil {
		return nil, fmt.Errorf("invalid eventtype")
	}
	if h.NotifType == nil {
		return nil, fmt.Errorf("invalid notiftype")
	}
	appID := h.AppID.String()
	eventID := uuid.NewString()

	handler := &generateHandler{
		Handler: h,
	}
	userIDs, err := handler.getNotifUsers(ctx, h.UserID, *h.NotifType, *h.EventType)
	if err != nil {
		return nil, err
	}

	reqs := []*npool.NotifReq{}
	for _, _userID := range userIDs {
		_reqs, err := handler.createUserNotifs(
			ctx,
			appID,
			eventID,
			_userID,
			h.Extra,
			*h.EventType,
			*h.NotifType,
			h.Vars,
		)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, _reqs...)
	}

	notifGenerateHandler, err := NewHandler(
		ctx,
		WithReqs(reqs, true),
	)
	if err != nil {
		return nil, err
	}
	notifs, err := notifGenerateHandler.CreateNotifs(ctx)
	if err != nil {
		return nil, err
	}

	return notifs, nil
}

type MultiNotifReq struct {
	UserID    *uuid.UUID
	EventType basetypes.UsedFor
	Vars      *templatemwpb.TemplateVars
	Extra     *string
	NotifType basetypes.NotifType
}

func (h *Handler) GenerateMultiNotifs(ctx context.Context) ([]*npool.Notif, error) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}

	appID := h.AppID.String()
	eventID := uuid.NewString()
	handler := &generateHandler{
		Handler: h,
	}
	reqs := []*npool.NotifReq{}

	for _, req := range h.MultiNotifReqs {
		userIDs, err := handler.getNotifUsers(ctx, req.UserID, req.NotifType, req.EventType)
		if err != nil {
			return nil, err
		}
		for _, _userID := range userIDs {
			_reqs, err := handler.createUserNotifs(
				ctx,
				appID,
				eventID,
				_userID,
				req.Extra,
				req.EventType,
				req.NotifType,
				req.Vars,
			)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs...)
		}
	}

	notifGenerateHandler, err := NewHandler(
		ctx,
		WithReqs(reqs, true),
	)
	if err != nil {
		return nil, err
	}
	notifs, err := notifGenerateHandler.CreateNotifs(ctx)
	if err != nil {
		return nil, err
	}

	return notifs, nil
}
