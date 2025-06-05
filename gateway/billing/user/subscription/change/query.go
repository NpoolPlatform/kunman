package change

import (
	"context"
	"fmt"

	gwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/billing/gateway/v1/user/subscription/change"
	subchangemwpb "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/user/subscription/change"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	subs  []*subchangemwpb.SubscriptionChange
	infos []*npool.UserSubscriptionChange
	apps  map[string]*appmwpb.App
	users map[string]*usermwpb.User
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = gwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, sub := range h.subs {
			appIDs = append(appIDs, sub.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = gwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, sub := range h.subs {
			userIDs = append(userIDs, sub.UserID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, sub := range h.subs {
		info := &npool.UserSubscriptionChange{
			ID:                 sub.ID,
			EntID:              sub.EntID,
			AppID:              sub.AppID,
			UserSubscriptionID: sub.UserSubscriptionID,
			OldPackageID:       sub.OldPackageID,
			NewPackageID:       sub.NewPackageID,
			CreatedAt:          sub.CreatedAt,
			UpdatedAt:          sub.UpdatedAt,
		}

		app, ok := h.apps[sub.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[sub.UserID]
		if ok {
			if user.Username != "" {
				info.Username = &user.Username
			}
			if user.EmailAddress != "" {
				info.EmailAddress = &user.EmailAddress
			}
			if user.PhoneNO != "" {
				info.PhoneNO = &user.PhoneNO
			}
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetSubscriptionChange(ctx context.Context) (*npool.UserSubscriptionChange, error) {
	sub, err := subchangemwcli.GetSubscriptionChange(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if sub == nil {
		return nil, fmt.Errorf("invalid subscription")
	}

	handler := &queryHandler{
		Handler: h,
		subs:    []*subchangemwpb.SubscriptionChange{sub},
		apps:    map[string]*appmwpb.App{},
		users:   map[string]*usermwpb.User{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetSubscriptionChanges(ctx context.Context) ([]*npool.UserSubscriptionChange, error) {
	conds := &subchangemwpb.Conds{}

	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}

	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}

	subs, err := subchangemwcli.GetSubscriptionChanges(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, err
	}
	if len(subs) == 0 {
		return nil, nil
	}

	handler := &queryHandler{
		Handler: h,
		subs:    subs,
		apps:    map[string]*appmwpb.App{},
		users:   map[string]*usermwpb.User{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	return handler.infos, nil
}

func (h *Handler) GetSubscriptionChangesCount(ctx context.Context) (uint32, error) {
	conds := &subchangemwpb.Conds{}

	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}

	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	return subchangemwcli.GetSubscriptionChangesCount(ctx, conds)
}
