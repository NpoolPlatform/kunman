package auth

import (
	"context"
)

type auth struct {
	AppID   string `sql:"app_id"`
	RoleID  string `sql:"role_id"`
	UserID  string `sql:"user_id"`
	AppVID  string `sql:"app_vid"`
	AppBID  string `sql:"app_bid"`
	UserVID string `sql:"user_vid"`
	UserBID string `sql:"user_bid"`
}

type existHandler struct {
	*Handler
	infos []*auth
}

func (h *Handler) ExistAuth(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}
	if h.UserID != nil {
		exist, err := handler.existUserAuth(ctx)
		if err != nil {
			return false, err
		}
		if exist {
			return true, nil
		}
		exist, err = handler.existRoleAuth(ctx)
		if err != nil {
			return false, err
		}
		if exist {
			return true, nil
		}
	}
	return handler.existAppAuth(ctx)
}
