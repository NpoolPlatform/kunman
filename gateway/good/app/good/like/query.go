package like

import (
	"context"
	"fmt"

	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/like"
	likemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/like"
	likemw "github.com/NpoolPlatform/kunman/middleware/good/app/good/like"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*checkHandler
	likes []*likemwpb.Like
	infos []*npool.Like
	apps  map[string]*appmwpb.App
	users map[string]*usermwpb.User
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, like := range h.likes {
			appIDs = append(appIDs, like.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = goodgwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, like := range h.likes {
			userIDs = append(userIDs, like.UserID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, like := range h.likes {
		info := &npool.Like{
			ID:        like.ID,
			EntID:     like.EntID,
			AppID:     like.AppID,
			UserID:    like.UserID,
			GoodID:    like.GoodID,
			AppGoodID: like.AppGoodID,
			GoodName:  like.GoodName,
			Like:      like.Like,
			CreatedAt: like.CreatedAt,
		}

		app, ok := h.apps[like.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[like.UserID]
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

func (h *Handler) GetLike(ctx context.Context) (*npool.Like, error) {
	likeHandler, err := likemw.NewHandler(
		ctx,
		likemw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	like, err := likeHandler.GetLike(ctx)
	if err != nil {
		return nil, err
	}
	if like == nil {
		return nil, fmt.Errorf("invalid like")
	}

	handler := &queryHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
		likes: []*likemwpb.Like{like},
		apps:  map[string]*appmwpb.App{},
		users: map[string]*usermwpb.User{},
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

func (h *Handler) GetLikes(ctx context.Context) ([]*npool.Like, uint32, error) {
	handler := &queryHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
		apps:  map[string]*appmwpb.App{},
		users: map[string]*usermwpb.User{},
	}
	if h.UserID != nil {
		if err := handler.CheckUser(ctx); err != nil {
			return nil, 0, err
		}
	}

	conds := &likemwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}

	likeHandler, err := likemw.NewHandler(
		ctx,
		likemw.WithConds(conds),
		likemw.WithOffset(h.Offset),
		likemw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	likes, total, err := likeHandler.GetLikes(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(likes) == 0 {
		return nil, total, nil
	}

	handler.likes = likes
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
