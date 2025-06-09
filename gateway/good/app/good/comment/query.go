package comment

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	commentmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/comment"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/comment"
	commentmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/comment"

	"github.com/google/uuid"
)

type queryHandler struct {
	*checkHandler
	comments []*commentmwpb.Comment
	infos    []*npool.Comment
	apps     map[string]*appmwpb.App
	users    map[string]*usermwpb.User
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = goodgwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, comment := range h.comments {
			appIDs = append(appIDs, comment.AppID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = goodgwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, comment := range h.comments {
			userIDs = append(userIDs, comment.UserID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *queryHandler) formalize() {
	for _, comment := range h.comments {
		info := &npool.Comment{
			ID:            comment.ID,
			EntID:         comment.EntID,
			AppID:         comment.AppID,
			UserID:        comment.UserID,
			GoodID:        comment.GoodID,
			AppGoodID:     comment.AppGoodID,
			GoodName:      comment.GoodName,
			Content:       comment.Content,
			Anonymous:     comment.Anonymous,
			PurchasedUser: comment.PurchasedUser,
			TrialUser:     comment.TrialUser,
			Score:         comment.Score,
			Hide:          comment.Hide,
			HideReason:    comment.HideReason,
			CreatedAt:     comment.CreatedAt,
			UpdatedAt:     comment.UpdatedAt,
		}

		if _, err := uuid.Parse(comment.OrderID); err == nil {
			if comment.OrderID != uuid.Nil.String() {
				info.OrderID = &comment.OrderID
			}
		}
		if _, err := uuid.Parse(comment.ReplyToID); err == nil {
			if comment.ReplyToID != uuid.Nil.String() {
				info.ReplyToID = &comment.ReplyToID
			}
		}

		app, ok := h.apps[comment.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[comment.UserID]
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

func (h *Handler) GetComment(ctx context.Context) (*npool.Comment, error) {
	comment, err := commentmwcli.GetComment(ctx, *h.EntID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if comment == nil {
		return nil, wlog.Errorf("invalid comment")
	}

	handler := &queryHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
		comments: []*commentmwpb.Comment{comment},
		apps:     map[string]*appmwpb.App{},
		users:    map[string]*usermwpb.User{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetComments(ctx context.Context) ([]*npool.Comment, uint32, error) {
	handler := &queryHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
		apps:  map[string]*appmwpb.App{},
		users: map[string]*usermwpb.User{},
	}
	if h.UserID != nil {
		if err := handler.CheckUser(ctx); err != nil {
			return nil, 0, wlog.Errorf("invalid user")
		}
	}

	conds := &commentmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	comments, total, err := commentmwcli.GetComments(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(comments) == 0 {
		return nil, total, nil
	}

	handler.comments = comments
	if err := handler.getApps(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}
